package webhooks

import (
	"log"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func DefaultHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DumpRequest error: %+v", err)
			return
	}

	os.Stdout.Write(b)
	w.Write(b)
}