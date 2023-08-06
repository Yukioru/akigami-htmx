install:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

vendors:
	mkdir -p public/vendors && cp node_modules/htmx.org/dist/htmx.min.js public/vendors/htmx.min.js

all: install vendors

.PHONY: install vendors