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

zip:
	rm -rf bin
	mkdir bin
	go build -o bin/gedcom2html$(EXT) ./gedcom2html
	go build -o bin/gedcomdiff$(EXT) ./gedcomdiff
	go build -o bin/gedcomq$(EXT) ./gedcomq
	go build -o bin/gedcomtune$(EXT) ./gedcomtune
	zip gedcom-$(NAME).zip -r bin

.PHONY: test zip
