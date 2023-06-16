package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	b := &bot{}
	ctx := context.Background()
	b.Start(ctx)
	const port = 23232
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), Server(b)))
}
