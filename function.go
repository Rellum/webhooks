package webhooks

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func DefaultHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DumpRequest error: %+v", err)
			return
	}

	w.Write(b)
}