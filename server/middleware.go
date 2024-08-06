/*
middleware.go
v0.1.0
1/2/24

This file is used to define middleware
*/
package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Logs the runtime of each http request
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start the timer
		startTime := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Calculate elapsed time
		elapsed := time.Since(startTime)

		// Get handler name
		handlerName := r.URL.String()
		log := fmt.Sprintf("Handler Name: %s, Elapsed Time (Seconds): %f", handlerName, elapsed.Seconds())
		logrus.Info(log)
	})
}
