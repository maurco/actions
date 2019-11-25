all: build-bin build-image run

build-bin:
	GOOS=linux GOARCH=amd64 \
	go build -v -o ./$(ACTION)/bin/$(ACTION) ./$(ACTION)/...

build-image:
	docker build \
		-t github-actions/$(ACTION) \
		$(FLAGS) $(ACTION)

run:
	docker run \
		-v $(HOME)/.aws:/root/.aws \
		-v /var/run/docker.sock:/var/run/docker.sock \
		$(FLAGS) github-actions/$(ACTION)
