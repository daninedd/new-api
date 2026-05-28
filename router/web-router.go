package router

import (
	"bytes"
	"embed"
	"html"
	"net/http"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/controller"
	"github.com/QuantumNous/new-api/middleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// ThemeAssets holds the embedded frontend assets for both themes.
type ThemeAssets struct {
	DefaultBuildFS   embed.FS
	DefaultIndexPage []byte
	ClassicBuildFS   embed.FS
	ClassicIndexPage []byte
}

func SetWebRouter(router *gin.Engine, assets ThemeAssets) {
	defaultFS := common.EmbedFolder(assets.DefaultBuildFS, "web/default/dist")
	classicFS := common.EmbedFolder(assets.ClassicBuildFS, "web/classic/dist")
	themeFS := common.NewThemeAwareFS(defaultFS, classicFS)

	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(middleware.GlobalWebRateLimit())
	router.Use(middleware.Cache())
	router.Use(static.Serve("/", themeFS))
	router.NoRoute(func(c *gin.Context) {
		c.Set(middleware.RouteTagKey, "web")
		if strings.HasPrefix(c.Request.RequestURI, "/v1") || strings.HasPrefix(c.Request.RequestURI, "/api") || strings.HasPrefix(c.Request.RequestURI, "/assets") {
			controller.RelayNotFound(c)
			return
		}
		if !shouldServeWebIndex(c.Request.URL.Path) {
			c.String(http.StatusNotFound, "404 page not found")
			return
		}
		c.Header("Cache-Control", "no-cache")
		indexPage := assets.DefaultIndexPage
		if common.GetTheme() == "classic" {
			indexPage = assets.ClassicIndexPage
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", renderWebIndex(indexPage, c.Request.URL.Path))
	})
}

var webIndexExactPaths = map[string]struct{}{
	"/":                {},
	"/401":             {},
	"/403":             {},
	"/404":             {},
	"/500":             {},
	"/503":             {},
	"/about":           {},
	"/console/log":     {},
	"/console/topup":   {},
	"/forgot-password": {},
	"/oauth":           {},
	"/otp":             {},
	"/pricing":         {},
	"/privacy-policy":  {},
	"/rankings":        {},
	"/register":        {},
	"/reset":           {},
	"/setup":           {},
	"/sign-in":         {},
	"/sign-up":         {},
	"/user/reset":      {},
	"/user-agreement":  {},
}

var webIndexPathPrefixes = []string{
	"/channels",
	"/chat",
	"/chat2link",
	"/dashboard",
	"/errors",
	"/keys",
	"/models",
	"/oauth",
	"/playground",
	"/pricing",
	"/profile",
	"/redemption-codes",
	"/subscriptions",
	"/system-settings",
	"/usage-logs",
	"/users",
	"/wallet",
}

func shouldServeWebIndex(requestPath string) bool {
	if requestPath == "" {
		return true
	}
	requestPath = strings.TrimRight(requestPath, "/")
	if requestPath == "" {
		requestPath = "/"
	}
	if _, ok := webIndexExactPaths[requestPath]; ok {
		return true
	}
	for _, prefix := range webIndexPathPrefixes {
		if requestPath == prefix || strings.HasPrefix(requestPath, prefix+"/") {
			return true
		}
	}
	return false
}

type webSEOMeta struct {
	Title       string
	Description string
	Canonical   string
}

var defaultWebSEOMeta = webSEOMeta{
	Title:       "All-LLMs - Unified AI API Gateway",
	Description: "All-LLMs is a unified AI API gateway and model management dashboard for OpenAI, Claude, Gemini, DeepSeek, Qwen, Llama, and other providers.",
	Canonical:   "https://all-llms.com/",
}

func renderWebIndex(indexPage []byte, requestPath string) []byte {
	meta := getWebSEOMeta(requestPath)
	htmlPage := bytes.ReplaceAll(indexPage, []byte("<html lang=\"zh\">"), []byte("<html lang=\"en\">"))
	htmlPage = bytes.ReplaceAll(htmlPage, []byte("<html lang=\"zh-CN\">"), []byte("<html lang=\"en\">"))
	htmlPage = bytes.ReplaceAll(htmlPage, []byte("<html lang=\"en\">"), []byte("<html lang=\"en\">"))

	htmlPage = replaceOrInsertHeadTag(htmlPage, "<title>", "</title>", "<title>"+html.EscapeString(meta.Title)+"</title>")
	htmlPage = replaceOrInsertMetaContent(htmlPage, `name="title"`, `<meta name="title" content="`+html.EscapeString(meta.Title)+`" />`)
	htmlPage = replaceOrInsertMetaContent(htmlPage, `name="description"`, `<meta name="description" content="`+html.EscapeString(meta.Description)+`" />`)
	htmlPage = replaceOrInsertLinkRel(htmlPage, `rel="canonical"`, `<link rel="canonical" href="`+html.EscapeString(meta.Canonical)+`" />`)
	return htmlPage
}

func getWebSEOMeta(requestPath string) webSEOMeta {
	path := normalizeWebPath(requestPath)
	switch {
	case path == "/pricing" || strings.HasPrefix(path, "/pricing/"):
		return webSEOMeta{
			Title:       "Model Marketplace - All-LLMs",
			Description: "Compare and access AI models through All-LLMs with unified routing, transparent pricing, usage tracking, and provider management.",
			Canonical:   canonicalURL(path),
		}
	case path == "/about":
		return webSEOMeta{
			Title:       "About All-LLMs",
			Description: "Learn about All-LLMs, a unified gateway for connecting and managing AI model providers through compatible API routes.",
			Canonical:   canonicalURL(path),
		}
	case path == "/user-agreement":
		return webSEOMeta{
			Title:       "User Agreement - All-LLMs",
			Description: "Read the All-LLMs user agreement, including service terms, account responsibilities, and platform usage rules.",
			Canonical:   canonicalURL(path),
		}
	case path == "/privacy-policy":
		return webSEOMeta{
			Title:       "Privacy Policy - All-LLMs",
			Description: "Read the All-LLMs privacy policy to understand how account, usage, and service data are handled.",
			Canonical:   canonicalURL(path),
		}
	default:
		return defaultWebSEOMeta
	}
}

func normalizeWebPath(requestPath string) string {
	if requestPath == "" {
		return "/"
	}
	path := strings.TrimRight(requestPath, "/")
	if path == "" {
		return "/"
	}
	return path
}

func canonicalURL(path string) string {
	if path == "/" {
		return "https://all-llms.com/"
	}
	return "https://all-llms.com" + path
}

func replaceOrInsertHeadTag(page []byte, openTag string, closeTag string, replacement string) []byte {
	start := bytes.Index(page, []byte(openTag))
	if start >= 0 {
		end := bytes.Index(page[start:], []byte(closeTag))
		if end >= 0 {
			end += start + len(closeTag)
			return append(append(append([]byte{}, page[:start]...), []byte(replacement)...), page[end:]...)
		}
	}
	return insertBeforeHeadEnd(page, replacement)
}

func replaceOrInsertMetaContent(page []byte, marker string, replacement string) []byte {
	return replaceOrInsertSingleTag(page, "<meta", marker, replacement)
}

func replaceOrInsertLinkRel(page []byte, marker string, replacement string) []byte {
	return replaceOrInsertSingleTag(page, "<link", marker, replacement)
}

func replaceOrInsertSingleTag(page []byte, prefix string, marker string, replacement string) []byte {
	var result []byte
	searchFrom := 0
	lastCopied := 0
	replaced := false
	for {
		start := bytes.Index(page[searchFrom:], []byte(prefix))
		if start < 0 {
			if !replaced {
				return insertBeforeHeadEnd(page, replacement)
			}
			result = append(result, page[lastCopied:]...)
			return result
		}
		start += searchFrom
		end := bytes.Index(page[start:], []byte(">"))
		if end < 0 {
			if !replaced {
				return insertBeforeHeadEnd(page, replacement)
			}
			result = append(result, page[lastCopied:]...)
			return result
		}
		end += start + 1
		tag := page[start:end]
		if bytes.Contains(tag, []byte(marker)) {
			result = append(result, page[lastCopied:start]...)
			if !replaced {
				result = append(result, []byte(replacement)...)
				replaced = true
			}
			lastCopied = end
		}
		searchFrom = end
	}
}

func insertBeforeHeadEnd(page []byte, tag string) []byte {
	headEnd := bytes.Index(page, []byte("</head>"))
	if headEnd < 0 {
		return page
	}
	insert := []byte("    " + tag + "\n")
	return append(append(append([]byte{}, page[:headEnd]...), insert...), page[headEnd:]...)
}
