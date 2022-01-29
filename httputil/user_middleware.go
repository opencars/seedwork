package httputil

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/opencars/seedwork"
)

func SessionCheckerMiddleware(checker seedwork.SessionChecker) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return Handler(func(w http.ResponseWriter, r *http.Request) error {
			sessionToken := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)

			user, err := checker.CheckSession(r.Context(), sessionToken, r.Header.Get("Cookie"))
			if err != nil {
				return ErrUnauthorized
			}

			ctx := WithUserID(r.Context(), user.ID)
			next.ServeHTTP(w, r.WithContext(ctx))

			return nil
		})
	}
}
