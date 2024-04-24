package middleware

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/mqnoy/go-todolist-rest-api/handler"
	"github.com/mqnoy/go-todolist-rest-api/pkg/clogger"
)

func PanicRecoverer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				buf = buf[:n]

				clogger.Logger().Errorf("recovering from err %v\n %s", err, buf)

				handler.ParseToErrorMsg(w, r, http.StatusInternalServerError, fmt.Errorf("%s", "Internal server error"))
			}
		}()

		h.ServeHTTP(w, r)
	})
}
