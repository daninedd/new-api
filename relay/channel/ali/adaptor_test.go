package ali

import (
	"net/http"
	"net/http/httptest"
	"testing"

	relaycommon "github.com/QuantumNous/new-api/relay/common"
	relayconstant "github.com/QuantumNous/new-api/relay/constant"

	"github.com/gin-gonic/gin"
)

func TestImageGenerationUsesUpstreamModelForSyncEndpoint(t *testing.T) {
	info := &relaycommon.RelayInfo{
		RelayMode:       relayconstant.RelayModeImagesGenerations,
		OriginModelName: "dall-e-3",
		ChannelMeta: &relaycommon.ChannelMeta{
			ChannelBaseUrl:    "https://dashscope.aliyuncs.com",
			UpstreamModelName: "qwen-image",
		},
	}

	adaptor := &Adaptor{}
	got, err := adaptor.GetRequestURL(info)
	if err != nil {
		t.Fatalf("GetRequestURL returned error: %v", err)
	}

	want := "https://dashscope.aliyuncs.com/api/v1/services/aigc/multimodal-generation/generation"
	if got != want {
		t.Fatalf("GetRequestURL = %q, want %q", got, want)
	}
}

func TestImageGenerationSyncHeaderUsesUpstreamModel(t *testing.T) {
	info := &relaycommon.RelayInfo{
		RelayMode:       relayconstant.RelayModeImagesGenerations,
		OriginModelName: "dall-e-3",
		ChannelMeta: &relaycommon.ChannelMeta{
			UpstreamModelName: "qwen-image",
		},
	}

	headers := http.Header{}
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest(http.MethodPost, "/v1/images/generations", nil)

	if err := (&Adaptor{}).SetupRequestHeader(c, &headers, info); err != nil {
		t.Fatalf("SetupRequestHeader returned error: %v", err)
	}

	if got := headers.Get("X-DashScope-Async"); got != "" {
		t.Fatalf("X-DashScope-Async = %q, want empty for sync image model", got)
	}
}
