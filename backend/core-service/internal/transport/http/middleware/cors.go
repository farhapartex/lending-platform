package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	HeaderOrigin           = "Origin"
	HeaderAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderMaxAge           = "Access-Control-Max-Age"
	HeaderRequestMethod    = "Access-Control-Request-Method"
	HeaderVary             = "Vary"
	wildcardOrigin         = "*"
	preflightMaxAgeSeconds = 600
)

var (
	allowedMethods  = []string{http.MethodGet, http.MethodOptions}
	allowedHeaders  = []string{"Content-Type", HeaderRequestID}
	exposedHeaders  = []string{HeaderRequestID}
	varyOnPreflight = []string{HeaderOrigin, HeaderRequestMethod}
)

func CORS(allowedOrigins []string) gin.HandlerFunc {
	permitted := newOriginSet(allowedOrigins)

	return func(c *gin.Context) {
		origin := c.GetHeader(HeaderOrigin)
		isPreflight := c.Request.Method == http.MethodOptions && c.GetHeader(HeaderRequestMethod) != ""

		if origin == "" || !permitted.allows(origin) {
			if isPreflight {
				c.Header(HeaderVary, strings.Join(varyOnPreflight, ", "))
				c.AbortWithStatus(http.StatusForbidden)

				return
			}

			c.Next()

			return
		}

		writeAllowHeaders(c, origin, permitted.isWildcard)

		if isPreflight {
			writePreflightHeaders(c)
			c.AbortWithStatus(http.StatusNoContent)

			return
		}

		c.Next()
	}
}

type originSet struct {
	origins    map[string]struct{}
	isWildcard bool
}

func newOriginSet(allowedOrigins []string) originSet {
	set := originSet{origins: make(map[string]struct{}, len(allowedOrigins))}

	for _, origin := range allowedOrigins {
		trimmed := strings.TrimSpace(origin)
		if trimmed == "" {
			continue
		}

		if trimmed == wildcardOrigin {
			set.isWildcard = true

			continue
		}

		set.origins[strings.ToLower(trimmed)] = struct{}{}
	}

	return set
}

func (s originSet) allows(origin string) bool {
	if s.isWildcard {
		return true
	}

	_, found := s.origins[strings.ToLower(origin)]

	return found
}

func writeAllowHeaders(c *gin.Context, origin string, isWildcard bool) {
	if isWildcard {
		c.Header(HeaderAllowOrigin, wildcardOrigin)
	} else {
		c.Header(HeaderAllowOrigin, origin)
		c.Header(HeaderVary, HeaderOrigin)
	}

	c.Header(HeaderExposeHeaders, strings.Join(exposedHeaders, ", "))
}

func writePreflightHeaders(c *gin.Context) {
	c.Header(HeaderAllowMethods, strings.Join(allowedMethods, ", "))
	c.Header(HeaderAllowHeaders, strings.Join(allowedHeaders, ", "))
	c.Header(HeaderMaxAge, strconv.Itoa(preflightMaxAgeSeconds))
	c.Header(HeaderVary, strings.Join(varyOnPreflight, ", "))
}
