package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

func StackMiddlewares(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			cur := middlewares[i]
			next = cur(next)
		}

		return next
	}
}
