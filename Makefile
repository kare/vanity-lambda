
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bootstrap .

.PHONY: package
package:
	zip function.zip bootstrap

.PHONY: clean
clean:
	$(RM) function.zip bootstrap

