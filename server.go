package pubsub

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"math/rand"
	"time"
)

func Server() {
	mux := &http.ServeMux{}
	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		time.Sleep(time.Second * time.Duration(rand.Intn(2)))
		slog.Info("receive request", "path", "/", "body", string(body), "headers", fmt.Sprintf("%+v", r.Header))
		w.WriteHeader(http.StatusOK)
	})

	slog.Info("Listen Server 8080 ports.")
	http.ListenAndServe(":8080", mux)
}
