package route

import (
	"net/http"
	"time"
	"fmt"
)


func LoggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var startTimer time.Time = time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("[%s] Method %s Serving %s To %s Completed request in %v\n", startTimer.Format("02/01/2006 15:04:05"), r.Method, r.URL.Path, r.RemoteAddr, time.Since(startTimer))
	}
}
