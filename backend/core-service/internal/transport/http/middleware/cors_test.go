package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/middleware"
)

const webOrigin = "http://localhost:5173"

func newCORSEngine(allowedOrigins []string) *gin.Engine {
	gin.SetMode(gin.TestMode)

	engine := gin.New()
	engine.Use(middleware.CORS(allowedOrigins))
	engine.GET("/thing", func(c *gin.Context) { c.String(http.StatusOK, "served") })

	return engine
}

func send(engine *gin.Engine, method, path string, headers map[string]string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, path, nil)

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	engine.ServeHTTP(recorder, request)

	return recorder
}

func preflight(engine *gin.Engine, origin string) *httptest.ResponseRecorder {
	return send(engine, http.MethodOptions, "/thing", map[string]string{
		middleware.HeaderOrigin:        origin,
		middleware.HeaderRequestMethod: http.MethodGet,
	})
}

func TestCORSAllowsAPermittedOrigin(t *testing.T) {
	engine := newCORSEngine([]string{webOrigin})

	recorder := send(engine, http.MethodGet, "/thing", map[string]string{middleware.HeaderOrigin: webOrigin})

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected the request to be served, got %d", recorder.Code)
	}

	if got := recorder.Header().Get(middleware.HeaderAllowOrigin); got != webOrigin {
		t.Fatalf("expected the origin to be echoed, got %q", got)
	}

	if recorder.Body.String() != "served" {
		t.Fatalf("expected the handler to run, got %q", recorder.Body.String())
	}
}

func TestCORSEchoesTheOriginRatherThanAWildcard(t *testing.T) {
	engine := newCORSEngine([]string{webOrigin, "https://app.example.com"})

	recorder := send(engine, http.MethodGet, "/thing", map[string]string{
		middleware.HeaderOrigin: "https://app.example.com",
	})

	if got := recorder.Header().Get(middleware.HeaderAllowOrigin); got != "https://app.example.com" {
		t.Fatalf("expected the second origin to be echoed, got %q", got)
	}

	if got := recorder.Header().Get(middleware.HeaderVary); got != middleware.HeaderOrigin {
		t.Fatalf("expected Vary: Origin so caches do not serve one origin's headers to another, got %q", got)
	}
}

func TestCORSExposesTheRequestIDHeader(t *testing.T) {
	engine := newCORSEngine([]string{webOrigin})

	recorder := send(engine, http.MethodGet, "/thing", map[string]string{middleware.HeaderOrigin: webOrigin})

	if got := recorder.Header().Get(middleware.HeaderExposeHeaders); got != middleware.HeaderRequestID {
		t.Fatalf("expected the request id to be readable by the browser, got %q", got)
	}
}

func TestCORSIgnoresAnUnknownOriginButStillServesTheRequest(t *testing.T) {
	engine := newCORSEngine([]string{webOrigin})

	recorder := send(engine, http.MethodGet, "/thing", map[string]string{
		middleware.HeaderOrigin: "http://evil.example.com",
	})

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected the request to still be served, got %d", recorder.Code)
	}

	if got := recorder.Header().Get(middleware.HeaderAllowOrigin); got != "" {
		t.Fatalf("expected no allow header for an unknown origin, got %q", got)
	}
}

func TestCORSLeavesANonBrowserRequestAlone(t *testing.T) {
	engine := newCORSEngine([]string{webOrigin})

	recorder := send(engine, http.MethodGet, "/thing", nil)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected a request without an origin to be served, got %d", recorder.Code)
	}

	if got := recorder.Header().Get(middleware.HeaderAllowOrigin); got != "" {
		t.Fatalf("expected no cors headers without an origin, got %q", got)
	}
}

func TestCORSAnswersAPreflightWithoutRunningTheHandler(t *testing.T) {
	engine := newCORSEngine([]string{webOrigin})

	recorder := preflight(engine, webOrigin)

	if recorder.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", recorder.Code)
	}

	if recorder.Body.Len() != 0 {
		t.Fatalf("expected an empty preflight body, got %q", recorder.Body.String())
	}

	if got := recorder.Header().Get(middleware.HeaderAllowMethods); got != "GET, OPTIONS" {
		t.Fatalf("unexpected allowed methods %q", got)
	}

	if got := recorder.Header().Get(middleware.HeaderAllowHeaders); got != "Content-Type, X-Request-Id" {
		t.Fatalf("unexpected allowed headers %q", got)
	}

	if got := recorder.Header().Get(middleware.HeaderMaxAge); got != "600" {
		t.Fatalf("expected preflights to be cacheable for 600s, got %q", got)
	}
}

func TestCORSPreflightVariesOnTheRequestMethod(t *testing.T) {
	engine := newCORSEngine([]string{webOrigin})

	recorder := preflight(engine, webOrigin)

	if got := recorder.Header().Get(middleware.HeaderVary); got != "Origin, Access-Control-Request-Method" {
		t.Fatalf("unexpected vary header %q", got)
	}
}

func TestCORSRefusesAPreflightFromAnUnknownOrigin(t *testing.T) {
	engine := newCORSEngine([]string{webOrigin})

	recorder := preflight(engine, "http://evil.example.com")

	if recorder.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", recorder.Code)
	}

	if got := recorder.Header().Get(middleware.HeaderAllowOrigin); got != "" {
		t.Fatalf("expected no allow header, got %q", got)
	}

	if got := recorder.Header().Get(middleware.HeaderVary); got != "Origin, Access-Control-Request-Method" {
		t.Fatalf("expected a vary header even on refusal, got %q", got)
	}
}

func TestCORSRefusesAPreflightWithNoOriginAtAll(t *testing.T) {
	engine := newCORSEngine([]string{webOrigin})

	recorder := send(engine, http.MethodOptions, "/thing", map[string]string{
		middleware.HeaderRequestMethod: http.MethodGet,
	})

	if recorder.Code != http.StatusForbidden {
		t.Fatalf("expected 403 for an origin-less preflight, got %d", recorder.Code)
	}
}

func TestCORSTreatsABareOptionsAsANormalRequest(t *testing.T) {
	engine := newCORSEngine([]string{webOrigin})

	recorder := send(engine, http.MethodOptions, "/thing", map[string]string{middleware.HeaderOrigin: webOrigin})

	if recorder.Code == http.StatusNoContent {
		t.Fatal("expected an OPTIONS without a request-method header not to be treated as a preflight")
	}

	if got := recorder.Header().Get(middleware.HeaderAllowOrigin); got != webOrigin {
		t.Fatalf("expected the origin still to be allowed, got %q", got)
	}
}

func TestCORSMatchesTheOriginCaseInsensitively(t *testing.T) {
	engine := newCORSEngine([]string{"http://Localhost:5173"})

	recorder := send(engine, http.MethodGet, "/thing", map[string]string{middleware.HeaderOrigin: webOrigin})

	if got := recorder.Header().Get(middleware.HeaderAllowOrigin); got != webOrigin {
		t.Fatalf("expected a case insensitive host match, got %q", got)
	}
}

func TestCORSSupportsAWildcard(t *testing.T) {
	engine := newCORSEngine([]string{wildcard})

	recorder := send(engine, http.MethodGet, "/thing", map[string]string{
		middleware.HeaderOrigin: "http://anything.example.com",
	})

	if got := recorder.Header().Get(middleware.HeaderAllowOrigin); got != wildcard {
		t.Fatalf("expected a wildcard allow header, got %q", got)
	}

	if got := recorder.Header().Get(middleware.HeaderVary); got != "" {
		t.Fatalf("expected no vary header when every origin is allowed, got %q", got)
	}
}

func TestCORSWildcardAnswersPreflights(t *testing.T) {
	engine := newCORSEngine([]string{wildcard})

	recorder := preflight(engine, "http://anything.example.com")

	if recorder.Code != http.StatusNoContent {
		t.Fatalf("expected a wildcard preflight to succeed, got %d", recorder.Code)
	}
}

func TestCORSBlocksEveryOriginWhenNoneAreConfigured(t *testing.T) {
	for _, origins := range [][]string{nil, {}, {""}, {"   "}} {
		engine := newCORSEngine(origins)

		recorder := send(engine, http.MethodGet, "/thing", map[string]string{middleware.HeaderOrigin: webOrigin})

		if recorder.Code != http.StatusOK {
			t.Fatalf("expected the request to be served, got %d", recorder.Code)
		}

		if got := recorder.Header().Get(middleware.HeaderAllowOrigin); got != "" {
			t.Fatalf("expected no allow header with %v configured, got %q", origins, got)
		}

		if preflight(engine, webOrigin).Code != http.StatusForbidden {
			t.Fatalf("expected preflights to be refused with %v configured", origins)
		}
	}
}

func TestCORSTrimsConfiguredOrigins(t *testing.T) {
	engine := newCORSEngine([]string{"  " + webOrigin + "  "})

	recorder := send(engine, http.MethodGet, "/thing", map[string]string{middleware.HeaderOrigin: webOrigin})

	if got := recorder.Header().Get(middleware.HeaderAllowOrigin); got != webOrigin {
		t.Fatalf("expected surrounding spaces to be ignored, got %q", got)
	}
}

func TestCORSNeverAllowsCredentials(t *testing.T) {
	engine := newCORSEngine([]string{webOrigin})

	recorder := send(engine, http.MethodGet, "/thing", map[string]string{middleware.HeaderOrigin: webOrigin})

	if got := recorder.Header().Get("Access-Control-Allow-Credentials"); got != "" {
		t.Fatalf("phase 1 has no cookies or auth headers, so credentials must stay off, got %q", got)
	}
}

const wildcard = "*"
