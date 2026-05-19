package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/setting"
)

const (
	creemModerationPath    = "/v1/moderation/prompt"
	creemModerationTimeout = 5 * time.Second
)

type CreemModerationError struct {
	Code       string
	Message    string
	StatusCode int
}

func (e *CreemModerationError) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

type creemModerationRequest struct {
	Prompt     string `json:"prompt"`
	ExternalID string `json:"external_id,omitempty"`
}

type creemModerationResponse struct {
	ID         string `json:"id"`
	Object     string `json:"object"`
	Prompt     string `json:"prompt"`
	ExternalID string `json:"external_id,omitempty"`
	Decision   string `json:"decision"`
}

var (
	creemModerationHTTPClient      = &http.Client{Timeout: creemModerationTimeout}
	creemModerationBaseURLOverride = ""
)

func ShouldUseCreemModeration() bool {
	return strings.TrimSpace(setting.CreemApiKey) != ""
}

func ModeratePromptWithCreem(ctx context.Context, prompt string, externalID string) *CreemModerationError {
	prompt = strings.TrimSpace(prompt)
	if prompt == "" {
		logger.LogInfo(ctx, fmt.Sprintf("Creem prompt moderation skipped external_id=%q reason=empty_prompt", externalID))
		return nil
	}
	if !ShouldUseCreemModeration() {
		logger.LogInfo(ctx, fmt.Sprintf("Creem prompt moderation skipped external_id=%q reason=api_key_not_configured", externalID))
		return nil
	}

	startTime := time.Now()
	logger.LogInfo(ctx, fmt.Sprintf("Creem prompt moderation started external_id=%q prompt_length=%d test_mode=%t", externalID, len([]rune(prompt)), setting.CreemTestMode))

	reqBody, err := common.Marshal(creemModerationRequest{
		Prompt:     prompt,
		ExternalID: externalID,
	})
	if err != nil {
		moderationErr := newCreemModerationUnavailableError(fmt.Errorf("failed to serialize moderation request: %w", err))
		logCreemModerationFailure(ctx, externalID, "unavailable", moderationErr, startTime)
		return moderationErr
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, creemModerationURL(), bytes.NewReader(reqBody))
	if err != nil {
		moderationErr := newCreemModerationUnavailableError(fmt.Errorf("failed to create moderation request: %w", err))
		logCreemModerationFailure(ctx, externalID, "unavailable", moderationErr, startTime)
		return moderationErr
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", setting.CreemApiKey)

	resp, err := creemModerationHTTPClient.Do(req)
	if err != nil {
		moderationErr := newCreemModerationUnavailableError(fmt.Errorf("failed to send moderation request: %w", err))
		logCreemModerationFailure(ctx, externalID, "unavailable", moderationErr, startTime)
		return moderationErr
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		moderationErr := newCreemModerationUnavailableError(fmt.Errorf("failed to read moderation response: %w", readErr))
		logCreemModerationFailure(ctx, externalID, "unavailable", moderationErr, startTime)
		return moderationErr
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		moderationErr := newCreemModerationUnavailableError(fmt.Errorf("moderation_http_%d: %s", resp.StatusCode, strings.TrimSpace(string(body))))
		logCreemModerationFailure(ctx, externalID, "unavailable", moderationErr, startTime)
		return moderationErr
	}

	var result creemModerationResponse
	if err := common.Unmarshal(body, &result); err != nil {
		moderationErr := newCreemModerationUnavailableError(fmt.Errorf("failed to parse moderation response: %w", err))
		logCreemModerationFailure(ctx, externalID, "unavailable", moderationErr, startTime)
		return moderationErr
	}

	switch result.Decision {
	case "allow":
		logger.LogInfo(ctx, fmt.Sprintf("Creem prompt moderation passed external_id=%q moderation_id=%q decision=%s duration_ms=%d", externalID, result.ID, result.Decision, time.Since(startTime).Milliseconds()))
		return nil
	case "deny", "flag":
		moderationErr := &CreemModerationError{
			Code:       "prompt_rejected",
			Message:    "Your prompt could not be processed. Please revise and try again.",
			StatusCode: http.StatusBadRequest,
		}
		logger.LogWarn(ctx, fmt.Sprintf("Creem prompt moderation blocked external_id=%q moderation_id=%q decision=%s duration_ms=%d", externalID, result.ID, result.Decision, time.Since(startTime).Milliseconds()))
		return moderationErr
	default:
		moderationErr := newCreemModerationUnavailableError(fmt.Errorf("unexpected moderation decision: %q", result.Decision))
		logCreemModerationFailure(ctx, externalID, "unavailable", moderationErr, startTime)
		return moderationErr
	}
}

func logCreemModerationFailure(ctx context.Context, externalID string, status string, moderationErr *CreemModerationError, startTime time.Time) {
	logger.LogWarn(ctx, fmt.Sprintf("Creem prompt moderation %s external_id=%q code=%s status_code=%d duration_ms=%d error=%q", status, externalID, moderationErr.Code, moderationErr.StatusCode, time.Since(startTime).Milliseconds(), moderationErr.Message))
}

func creemModerationURL() string {
	if creemModerationBaseURLOverride != "" {
		return strings.TrimRight(creemModerationBaseURLOverride, "/") + creemModerationPath
	}
	if setting.CreemTestMode {
		return "https://test-api.creem.io" + creemModerationPath
	}
	return "https://api.creem.io" + creemModerationPath
}

func newCreemModerationUnavailableError(err error) *CreemModerationError {
	return &CreemModerationError{
		Code:       "moderation_unavailable",
		Message:    fmt.Sprintf("Prompt moderation is temporarily unavailable: %s", common.MaskSensitiveInfo(err.Error())),
		StatusCode: http.StatusServiceUnavailable,
	}
}
