package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

/*
// Logger middleware logs the request details and duration of the request processing.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record the start time
		start := time.Now()

		// Log the request details
		fmt.Printf("Request started - Method: %s, Path: %s\n", r.Method, r.URL.Path)

		// Pass the request to the next handler
		next.ServeHTTP(w, r)

		// Log the completion time
		fmt.Printf("Request completed - Method: %s, Path: %s, Duration: %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}


	func Security(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Add security headers
			w.Header().Set("Content-Security-Policy", "default-src 'self'")
			w.Header().Set("X-Frame-Options", "DENY")
			w.Header().Set("X-XSS-Protection", "1; mode=block")

			// Pass the request to the next handler
			next.ServeHTTP(w, r)
		})
	}
*/

func SecurityHeaders(ctx *fiber.Ctx) error {
	ctx.Response().Header.Add("X-Content-Type-Options", "nosniff")
	ctx.Response().Header.Add("X-Frame-Options", "DENY")
	return ctx.Next() // Proceed to the next middleware or handler
}
