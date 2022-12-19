package chi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bayusamudra5502/go-backend-template/lib/output"
	cm "github.com/go-chi/chi/v5/middleware"
)

var colorMap = map[string]output.Color{
	"GET": output.ForeGreen,
	"POST": output.ForeBlue,
	"PUT": output.ForeCyan,
	"PATCH": output.ForeMagenta,
	"DELETE": output.ForeRed,
}

func colorizeMethod(method string) string {
	val, ok := colorMap[method]

	if ok {
		return fmt.Sprintf("%s%s%s", val, method, output.Reset)
	} else {
		return fmt.Sprintf("%s%s%s", output.ForeWhite, method, output.Reset)
	}
}

func colorizeCode(code int) string {
	if code < 200 {
		return fmt.Sprintf("%s%d%s", output.ForeCyan, code, output.Reset)
	} else if code < 300 {
			return fmt.Sprintf("%s%d%s", output.ForeGreen, code, output.Reset)
	} else if code < 400 {
		return fmt.Sprintf("%s%d%s", output.ForeBlue, code, output.Reset)
	} else if code < 500 {
		return fmt.Sprintf("%s%d%s", output.ForeYellow, code, output.Reset)
	} else {
		return fmt.Sprintf("%s%d%s", output.ForeRed, code, output.Reset)
	}
}

func (m *ChiMiddleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		ww := cm.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			delta := time.Since(startTime)
			status := ww.Status()
			path := r.URL.Path
			method := r.Method

			m.log.Info(context.Background(), 
				fmt.Sprintf("Request %s %s %s (%dms)", 
					colorizeCode(status),
					colorizeMethod(method), 
					path, 
					delta.Milliseconds(),
				),
			)
		}()
		
		next.ServeHTTP(ww, r)
	})
}
