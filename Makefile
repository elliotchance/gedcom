test:
 	# Run all the tests with the race detector enabled.
	go test -race ./...

test-coverage:
	echo "" > coverage.txt

	for d in $$(go list ./... | grep -v vendor); do \
		go test -race -coverprofile=profile.out -covermode=atomic $$d; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
		fi \
	done

check-go-fmt:
	exit $$(go fmt ./... | wc -l)

zip:
	rm -rf bin
	mkdir bin
	go build -o bin/gedcom2html ./gedcom2html
	go build -o bin/gedcom2json ./gedcom2json
	go build -o bin/gedcom2text ./gedcom2text
	go build -o bin/gedcomdiff ./gedcomdiff
	go build -o bin/gedcomtune ./gedcomtune
	zip gedcom-$(GOOS)-$(GOARCH).zip -r bin

.PHONY: test zip
