package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// square the path param x in /square/:x
	http.HandleFunc("/square", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws := strings.Split(r.URL.Path, "/")
		if len(ws) < 3 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"error":"missing expected x value in /square/:x"}`)
		}

		x, err := strconv.Atoi(ws[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"error":"x must be numeric"}`)
		}
		x2 := square(x)
		fmt.Fprintf(w, `{"x":%d,"squared":%d}`, x, x2)
	}))
	http.ListenAndServe(":80", nil)
}

func square(x int) int {
	return x * x
}