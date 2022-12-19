package chi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/bayusamudra5502/go-backend-template/model/web"
)

func (l ChiMiddleware) Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				parser := json.NewEncoder(w)

				w.WriteHeader(http.StatusInternalServerError)
				payload := web.ErrorData("internal server error", nil)
	
				err := parser.Encode(payload)

				if err != nil {
					l.log.Error(context.Background(), "Failed to parse error:" + err.Error())
					l.log.Error(context.Background(), "")
					return
				}
				
				
				stacks := strings.Split(string(debug.Stack()), "\n")
				l.log.Error(context.Background(), "Some panic occured when processing request:")
				l.log.Error(context.Background(), fmt.Sprint(rec))
				l.log.Error(context.Background(), "")
				
				l.log.Error(context.Background(), "Stack Trace:")
				for _, val := range stacks {
					l.log.Error(context.Background(), val)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}
