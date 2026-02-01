package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"kkn.fi/vanity"
)

func main() {
	domain := os.Getenv("VANITY_DOMAIN")
	vcsUrl := os.Getenv("VANITY_VCSURL")

	opts := []vanity.Option{}
	if domain != "" {
		opts = append(opts, vanity.Domain(domain))
	}
	if vcsUrl != "" {
		opts = append(opts, vanity.VCSURL(vcsUrl))
	}

	h, err := vanity.NewHandlerWithOptions(opts...)
	if err != nil {
		log.Fatalf("failed to init vanity handler: %v", err)
	}
	handler := http.NewServeMux()
	// Avoid vanity handler panic on "/" when no static/index is configured.
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "" {
			if vcsUrl != "" {
				base := strings.TrimRight(vcsUrl, "/")
				http.Redirect(w, r, base+r.URL.Path, http.StatusFound)
				return
			}
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
	adapter := httpadapter.NewV2(handler)
	lambda.Start(adapter.ProxyWithContext)
}
