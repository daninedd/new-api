package router

import (
	"embed"
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
		if common.GetTheme() == "classic" {
			c.Data(http.StatusOK, "text/html; charset=utf-8", assets.ClassicIndexPage)
		} else {
			c.Data(http.StatusOK, "text/html; charset=utf-8", assets.DefaultIndexPage)
		}
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
