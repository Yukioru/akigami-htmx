install:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
.PHONY: install

vendors:
	mkdir -p public/vendors && wget https://unpkg.com/htmx.org/dist/htmx.min.js -nv -O public/vendors/htmx.min.js
.PHONY: vendors

dev:
	make -j 2 watch watch_server
.PHONY: dev

watch_server:
	./bin/air
.PHONY: watch_server

watch:
	make -j 2 watch_styles watch_scripts
.PHONY: watch

watch_scripts:
	npx esbuild scripts/app.ts --bundle --target=es2020 --outfile=public/assets/app.js --watch=forever
.PHONY: watch_scripts

watch_styles:
	npx tailwindcss -i ./styles/app.css -o ./public/assets/app.css --watch
.PHONY: watch_styles

build: build_scripts build_styles build_server
.PHONY: build

build_server:
	mkdir -p build
	cp .env build/
	cp -R public build/
	cp -R views build/
	mkdir -p build/locales/resources
	cp -R locales/resources build/locales/
	go build -o build/server main.go
.PHONY: build_server

build_scripts:
	npx esbuild scripts/app.ts --bundle --target=es2020 --outfile=public/assets/app.js --minify --sourcemap
.PHONY: build_scripts

build_styles:
	npx tailwindcss -i ./styles/app.css -o ./public/assets/app.css
.PHONY: build_styles

clean:
	rm -rf bin build public/vendors public/assets
.PHONY: clean

bootstrap: clean install vendors
.PHONY: bootstrap

all: clean install vendors build
.PHONY: all
