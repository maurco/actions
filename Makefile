all: build-bin build-image run-image

format:
	gofmt -w .

build-bin:
	docker run --rm \
		-w /app \
		-v $(PWD)/$(ACTION):/app \
		golang:1-alpine \
		go build -v -o bin/$(ACTION)-alpine ./...

build-image:
	docker build \
		-t github-actions/$(ACTION) \
		$(FLAGS) $(ACTION)

run-image:
	docker run --rm \
		-v $(HOME)/.aws:/root/.aws \
		-v /var/run/docker.sock:/var/run/docker.sock \
		$(FLAGS) github-actions/$(ACTION)
