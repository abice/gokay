COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)
endif

all: generate build gen-test test cover
install-deps:
	glide install
build: generate
	if [ ! -d bin ]; then mkdir bin; fi
	go build -v -o bin/gokay
fmt:
	@find . -not -path "./vendor/*" -name '*.go' -type f | sed 's#\(.*\)/.*#\1#' | sort -u | xargs -n1 -I {} bash -c "cd {} && goimports -w *.go && gofmt -w -s -l *.go"
gen-test:
	gokay ./internal/gkexample/example.go
test: generate
	if [ ! -d coverage ]; then mkdir coverage; fi
	go test -v ./gkgen -race -cover -coverprofile=$(COVERAGEDIR)/gkgen.coverprofile
	go test -v ./gokay -race -cover -coverprofile=$(COVERAGEDIR)/gokay.coverprofile
	go test -v ./internal/gkexample -race -cover -coverprofile=$(COVERAGEDIR)/gkexample.coverprofile
cover:
	go tool cover -html=$(COVERAGEDIR)/gkgen.coverprofile -o $(COVERAGEDIR)/gkgen.html
	go tool cover -html=$(COVERAGEDIR)/gokay.coverprofile -o $(COVERAGEDIR)/gokay.html
	go tool cover -html=$(COVERAGEDIR)/gokay.coverprofile -o $(COVERAGEDIR)/gkexample.html
tc: test cover
coveralls:
	gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile
	goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN)
clean:
	go clean
	rm -f bin/gokay
	rm -rf coverage/

# go-bindata will take all of the template files and create readable assets from them in the executable.
# This way the templates are readable in atom (or another text editor with golang template language support)
generate:
	go-bindata -o gkgen/assets.go -pkg=gkgen gkgen/*.tmpl

phony: clean tc build
