package main

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	json.NewEncoder(ctx).Encode(map[string]interface{}{
		"message": "hello world",
	})
}

func main() {
	port := ":3999"
	log.Fatal(fasthttp.ListenAndServe(port, fastHTTPHandler))
}
