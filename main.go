package main

import (
	"log"
	"net/http"
	"os"

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
	handler.Handle("/", h)
	adapter := httpadapter.New(handler)
	lambda.Start(adapter.ProxyWithContext)
}
