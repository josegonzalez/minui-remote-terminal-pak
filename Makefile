TAG ?= latest
PAK_NAME := $(shell jq -r .label config.json)

ARCHITECTURES := arm64
PLATFORMS := rg35xxplus tg5040

JQ_VERSION ?= 1.7.1
MINUI_LIST_VERSION := 0.11.3
MINUI_PRESENTER_VERSION := 0.7.0

clean:
	rm -f bin/*/remote-term || true
	rm -f bin/*/jq || true
	rm -f bin/*/minui-list || true
	rm -f bin/*/minui-presenter || true

build: $(foreach platform,$(PLATFORMS),bin/$(platform)/minui-list bin/$(platform)/minui-presenter) $(foreach arch,$(ARCHITECTURES),bin/$(arch)/remote-term bin/$(arch)/jq)

bin/arm64/remote-term:
	mkdir -p bin/arm64
	docker buildx build --platform linux/arm64 --load -f docker/remote-term/Dockerfile --progress plain -t app/remote-term:latest-arm64 docker/remote-term
	docker container create --name extract app/remote-term:latest-arm64
	docker container cp extract:/go/src/github.com/josegonzalez/go-remote-term/remote-term bin/arm64/remote-term
	docker container rm extract
	chmod +x bin/arm64/remote-term

bin/arm64/jq:
	mkdir -p bin/arm64
	curl -f -o bin/arm64/jq -sSL https://github.com/jqlang/jq/releases/download/jq-$(JQ_VERSION)/jq-linux-arm64
	curl -sSL -o bin/arm64/jq.LICENSE "https://raw.githubusercontent.com/jqlang/jq/refs/heads/$(JQ_VERSION)/COPYING"

bin/%/minui-list:
	mkdir -p bin/$*
	curl -f -o bin/$*/minui-list -sSL https://github.com/josegonzalez/minui-list/releases/download/$(MINUI_LIST_VERSION)/minui-list-$*
	chmod +x bin/$*/minui-list

bin/%/minui-presenter:
	mkdir -p bin/$*
	curl -f -o bin/$*/minui-presenter -sSL https://github.com/josegonzalez/minui-presenter/releases/download/$(MINUI_PRESENTER_VERSION)/minui-presenter-$*
	chmod +x bin/$*/minui-presenter

release: build
	mkdir -p dist
	git archive --format=zip --output "dist/$(PAK_NAME).pak.zip" HEAD
	while IFS= read -r file; do zip -r "dist/$(PAK_NAME).pak.zip" "$$file"; done < .gitarchiveinclude
	ls -lah dist
