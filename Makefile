all: build-bin build-image run

build-bin:
	docker run --rm \
		-v $$(PWD)/$(ACTION):/app \
		-w /app \
		golang:1-alpine \
		go build -v -o ./bin/$(ACTION) ./...

build-image:
	docker build \
		-t github-actions/$(ACTION) \
		$(FLAGS) $(ACTION)

run:
	docker run --rm \
		-v $(HOME)/.aws:/root/.aws \
		-v /var/run/docker.sock:/var/run/docker.sock \
		$(FLAGS) github-actions/$(ACTION)
