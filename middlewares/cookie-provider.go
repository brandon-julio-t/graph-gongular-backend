package middlewares

import (
	"context"
	"net/http"
)

type cookieProviderKeyStruct struct{ name string }

var cookieProviderKey = cookieProviderKeyStruct{name: "cookie-provider"}

func CookieProviderMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), cookieProviderKey, &w)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UseCookieProvider(ctx context.Context) *http.ResponseWriter {
	if w, ok := ctx.Value(cookieProviderKey).(*http.ResponseWriter); ok {
		return w
	}
	return nil
}
