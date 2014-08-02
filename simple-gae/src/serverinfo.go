package simplegae

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/serverinfo", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
