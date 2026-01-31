
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bootstrap .

.PHONY: package
package:
	zip function.zip bootstrap

.PHONY: deploy
deploy:
	aws s3 cp function.zip s3://cloud-infra-artifacts-kkn-eu-north-1/lambda/vanity/function.zip --profile kkn

.PHONY: clean
clean:
	$(RM) function.zip bootstrap

