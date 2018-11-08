test:
 	# Run all the tests with the race detector enabled.
	go test -race ./...

test-coverage:
	echo "" > coverage.txt

	for d in $$(go list ./... | grep -v vendor); do \
		go test -race -coverprofile=profile.out -covermode=atomic $$d || exit 1; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
		fi \
	done

check-go-fmt:
	exit $$(go fmt ./... | wc -l)

check-ghost:
	exit $$(ghost -max-line-complexity 2 -ignore-tests $$(find . -name "*.go" | grep -v vendor))

checks: check-go-fmt check-ghost

zip:
	rm -rf bin
	mkdir bin
	go build -o bin/gedcom2html ./gedcom2html
	go build -o bin/gedcom2json ./gedcom2json
	go build -o bin/gedcom2text ./gedcom2text
	go build -o bin/gedcomdiff ./gedcomdiff
	go build -o bin/gedcomq ./gedcomq
	go build -o bin/gedcomtune ./gedcomtune
	zip gedcom-$(GOOS)-$(GOARCH).zip -r bin

.PHONY: test zip
