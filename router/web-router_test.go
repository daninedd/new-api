package router

import (
	"strings"
	"testing"
)

func TestShouldServeWebIndex(t *testing.T) {
	tests := []struct {
		name        string
		requestPath string
		want        bool
	}{
		{
			name:        "home",
			requestPath: "/",
			want:        true,
		},
		{
			name:        "public pricing route",
			requestPath: "/pricing",
			want:        true,
		},
		{
			name:        "dynamic pricing route",
			requestPath: "/pricing/gpt-4.1",
			want:        true,
		},
		{
			name:        "authenticated app route",
			requestPath: "/dashboard/overview",
			want:        true,
		},
		{
			name:        "unknown route",
			requestPath: "/nonexistent-seo-test-20260528",
			want:        false,
		},
		{
			name:        "missing static asset",
			requestPath: "/static/js/missing.js",
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shouldServeWebIndex(tt.requestPath); got != tt.want {
				t.Fatalf("shouldServeWebIndex(%q) = %v, want %v", tt.requestPath, got, tt.want)
			}
		})
	}
}

func TestRenderWebIndexInjectsRouteSEO(t *testing.T) {
	indexPage := []byte(`<!doctype html>
<html lang="en">
  <head>
    <title>All-LLMs</title>
    <meta name="title" content="All-LLMs" />
    <meta name="description" content="Unified AI API gateway and admin dashboard." />
  </head>
  <body><div id="root"></div></body>
</html>`)

	rendered := string(renderWebIndex(indexPage, "/pricing"))
	assertContains(t, rendered, `<html lang="en">`)
	assertContains(t, rendered, `<title>Model Marketplace - All-LLMs</title>`)
	assertContains(t, rendered, `<meta name="title" content="Model Marketplace - All-LLMs" />`)
	assertContains(t, rendered, `<link rel="canonical" href="https://all-llms.com/pricing" />`)
	assertContains(t, rendered, "Compare and access AI models through All-LLMs")
}

func TestRenderWebIndexNormalizesClassicLanguage(t *testing.T) {
	indexPage := []byte(`<!doctype html>
<html lang="zh">
  <head>
    <meta name="description" lang="zh" content="旧中文描述" />
    <meta name="description" lang="en" content="Old English description" />
    <title>All-LLMs</title>
  </head>
  <body><div id="root"></div></body>
</html>`)

	rendered := string(renderWebIndex(indexPage, "/"))
	assertContains(t, rendered, `<html lang="en">`)
	assertContains(t, rendered, `<link rel="canonical" href="https://all-llms.com/" />`)
	if count := strings.Count(rendered, `name="description"`); count != 1 {
		t.Fatalf("description meta count = %d, want 1:\n%s", count, rendered)
	}
}

func TestGetWebSEOMeta(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		title     string
		canonical string
	}{
		{
			name:      "home",
			path:      "/",
			title:     "All-LLMs - Unified AI API Gateway",
			canonical: "https://all-llms.com/",
		},
		{
			name:      "pricing detail keeps path",
			path:      "/pricing/gpt-4.1/",
			title:     "Model Marketplace - All-LLMs",
			canonical: "https://all-llms.com/pricing/gpt-4.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getWebSEOMeta(tt.path)
			if got.Title != tt.title {
				t.Fatalf("Title = %q, want %q", got.Title, tt.title)
			}
			if got.Canonical != tt.canonical {
				t.Fatalf("Canonical = %q, want %q", got.Canonical, tt.canonical)
			}
		})
	}
}

func assertContains(t *testing.T, s string, substr string) {
	t.Helper()
	if !strings.Contains(s, substr) {
		t.Fatalf("expected rendered HTML to contain %q, got:\n%s", substr, s)
	}
}
