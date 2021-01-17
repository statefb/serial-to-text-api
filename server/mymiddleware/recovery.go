package mymiddleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("panic catch: %v", err)

				resp, _ := json.Marshal(map[string]string{
					"message": fmt.Sprintf("Internal Server Error: %v", err),
					"code":    "500", // 予期せぬエラー
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write(resp)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
