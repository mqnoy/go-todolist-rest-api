GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get


# Build target
APP_NAME=core

generate_envrc: .env
	@echo "Generating .envrc from .env..."
	@sed 's/^\(.*\)=\(.*\)/export \1="\2"/' .env >> .envrc
	@echo "Generated .envrc successfully."

build:
	@latest_tag=$$(git describe --tags --abbrev=0); \
	echo "Building binary with version $$latest_tag"; \
	$(GOBUILD) -ldflags "-X main.version=$$latest_tag" -o build/$(APP_NAME) -v .

build-image:
	@latest_tag=$$(git describe --tags --abbrev=0); \
	echo "Building docker images with version $$latest_tag"; \
	docker build -t $(APP_NAME):$$latest_tag .

run-image:
	@docker compose up --build
