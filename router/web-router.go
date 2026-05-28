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
	"/":                              {},
	"/401":                           {},
	"/403":                           {},
	"/404":                           {},
	"/500":                           {},
	"/503":                           {},
	"/about":                         {},
	"/ai-model-router":               {},
	"/claude-api-gateway":            {},
	"/console/log":                   {},
	"/console/topup":                 {},
	"/forgot-password":               {},
	"/gemini-api-gateway":            {},
	"/oauth":                         {},
	"/openai-compatible-api-gateway": {},
	"/otp":                           {},
	"/pricing":                       {},
	"/privacy-policy":                {},
	"/rankings":                      {},
	"/register":                      {},
	"/reset":                         {},
	"/self-hosted-ai-gateway":        {},
	"/setup":                         {},
	"/sign-in":                       {},
	"/sign-up":                       {},
	"/user/reset":                    {},
	"/user-agreement":                {},
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
	PageType    string
}

var defaultWebSEOMeta = webSEOMeta{
	Title:       "All-LLMs - Unified AI API Gateway",
	Description: "All-LLMs is a unified AI API gateway and model management dashboard for OpenAI, Claude, Gemini, DeepSeek, Qwen, Llama, and other providers.",
	Canonical:   "https://all-llms.com/",
	PageType:    "WebSite",
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
	htmlPage = injectSocialMeta(htmlPage, meta)
	htmlPage = replaceOrInsertJSONLD(htmlPage, buildWebJSONLD(meta))
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
			PageType:    "CollectionPage",
		}
	case path == "/about":
		return webSEOMeta{
			Title:       "About All-LLMs",
			Description: "Learn about All-LLMs, a unified gateway for connecting and managing AI model providers through compatible API routes.",
			Canonical:   canonicalURL(path),
			PageType:    "AboutPage",
		}
	case path == "/openai-compatible-api-gateway":
		return webSEOMeta{
			Title:       "OpenAI-compatible API Gateway - All-LLMs",
			Description: "Use All-LLMs as an OpenAI-compatible API gateway to route requests across OpenAI, Claude, Gemini, DeepSeek, Qwen, Llama, and other providers.",
			Canonical:   canonicalURL(path),
			PageType:    "WebPage",
		}
	case path == "/claude-api-gateway":
		return webSEOMeta{
			Title:       "Claude API Gateway - All-LLMs",
			Description: "Connect Claude workflows through All-LLMs with unified access control, usage tracking, quotas, billing, and multi-provider routing.",
			Canonical:   canonicalURL(path),
			PageType:    "WebPage",
		}
	case path == "/gemini-api-gateway":
		return webSEOMeta{
			Title:       "Gemini API Gateway - All-LLMs",
			Description: "Route Gemini traffic through All-LLMs with shared API keys, monitoring, limits, fallback options, and centralized provider management.",
			Canonical:   canonicalURL(path),
			PageType:    "WebPage",
		}
	case path == "/ai-model-router":
		return webSEOMeta{
			Title:       "AI Model Router - All-LLMs",
			Description: "Route AI requests across providers and models with centralized policies for fallback, quotas, cost tracking, and observability.",
			Canonical:   canonicalURL(path),
			PageType:    "WebPage",
		}
	case path == "/self-hosted-ai-gateway":
		return webSEOMeta{
			Title:       "Self-hosted AI Gateway - All-LLMs",
			Description: "Deploy a self-hosted AI API gateway to manage providers, users, keys, billing, and usage logs in your own environment.",
			Canonical:   canonicalURL(path),
			PageType:    "WebPage",
		}
	case path == "/user-agreement":
		return webSEOMeta{
			Title:       "User Agreement - All-LLMs",
			Description: "Read the All-LLMs user agreement, including service terms, account responsibilities, and platform usage rules.",
			Canonical:   canonicalURL(path),
			PageType:    "WebPage",
		}
	case path == "/privacy-policy":
		return webSEOMeta{
			Title:       "Privacy Policy - All-LLMs",
			Description: "Read the All-LLMs privacy policy to understand how account, usage, and service data are handled.",
			Canonical:   canonicalURL(path),
			PageType:    "WebPage",
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

const ogImageURL = "https://all-llms.com/og-image.png"

func injectSocialMeta(page []byte, meta webSEOMeta) []byte {
	tags := []struct {
		marker string
		tag    string
	}{
		{`property="og:type"`, `<meta property="og:type" content="website" />`},
		{`property="og:site_name"`, `<meta property="og:site_name" content="All-LLMs" />`},
		{`property="og:title"`, `<meta property="og:title" content="` + html.EscapeString(meta.Title) + `" />`},
		{`property="og:description"`, `<meta property="og:description" content="` + html.EscapeString(meta.Description) + `" />`},
		{`property="og:url"`, `<meta property="og:url" content="` + html.EscapeString(meta.Canonical) + `" />`},
		{`property="og:image"`, `<meta property="og:image" content="` + ogImageURL + `" />`},
		{`property="og:image:width"`, `<meta property="og:image:width" content="1200" />`},
		{`property="og:image:height"`, `<meta property="og:image:height" content="630" />`},
		{`property="og:image:alt"`, `<meta property="og:image:alt" content="All-LLMs unified AI API gateway social preview" />`},
		{`name="twitter:card"`, `<meta name="twitter:card" content="summary_large_image" />`},
		{`name="twitter:title"`, `<meta name="twitter:title" content="` + html.EscapeString(meta.Title) + `" />`},
		{`name="twitter:description"`, `<meta name="twitter:description" content="` + html.EscapeString(meta.Description) + `" />`},
		{`name="twitter:image"`, `<meta name="twitter:image" content="` + ogImageURL + `" />`},
		{`name="twitter:image:alt"`, `<meta name="twitter:image:alt" content="All-LLMs unified AI API gateway social preview" />`},
	}
	for _, item := range tags {
		page = replaceOrInsertSingleTag(page, "<meta", item.marker, item.tag)
	}
	return page
}

func buildWebJSONLD(meta webSEOMeta) string {
	graph := []map[string]any{
		{
			"@type":  "Organization",
			"@id":    "https://all-llms.com/#organization",
			"name":   "All-LLMs",
			"url":    "https://all-llms.com/",
			"logo":   "https://all-llms.com/logo.png",
			"sameAs": []string{"https://github.com/QuantumNous/new-api"},
		},
		{
			"@type": "WebSite",
			"@id":   "https://all-llms.com/#website",
			"url":   "https://all-llms.com/",
			"name":  "All-LLMs",
			"publisher": map[string]string{
				"@id": "https://all-llms.com/#organization",
			},
		},
		{
			"@type":               "SoftwareApplication",
			"@id":                 "https://all-llms.com/#software",
			"name":                "All-LLMs",
			"applicationCategory": "DeveloperApplication",
			"operatingSystem":     "Web",
			"url":                 "https://all-llms.com/",
			"description":         defaultWebSEOMeta.Description,
			"image":               ogImageURL,
			"publisher": map[string]string{
				"@id": "https://all-llms.com/#organization",
			},
			"offers": map[string]string{
				"@type":         "Offer",
				"price":         "0",
				"priceCurrency": "USD",
			},
		},
		{
			"@type":       meta.PageType,
			"@id":         meta.Canonical + "#webpage",
			"url":         meta.Canonical,
			"name":        meta.Title,
			"description": meta.Description,
			"image":       ogImageURL,
			"isPartOf": map[string]string{
				"@id": "https://all-llms.com/#website",
			},
			"publisher": map[string]string{
				"@id": "https://all-llms.com/#organization",
			},
		},
	}
	data := map[string]any{
		"@context": "https://schema.org",
		"@graph":   graph,
	}
	jsonLD, err := common.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonLD)
}

func replaceOrInsertJSONLD(page []byte, jsonLD string) []byte {
	if jsonLD == "" {
		return page
	}
	return replaceOrInsertHeadTag(
		page,
		`<script type="application/ld+json">`,
		`</script>`,
		`<script type="application/ld+json">`+jsonLD+`</script>`,
	)
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
