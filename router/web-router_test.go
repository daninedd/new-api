package router

import "testing"

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
