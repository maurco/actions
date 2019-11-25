all: build run

build:
	docker build \
		-t github-actions/$(ACTION) \
		$(FLAGS) $(ACTION)

run:
	docker run --rm \
		-v $(HOME)/.aws:/root/.aws \
		-v /var/run/docker.sock:/var/run/docker.sock \
		$(FLAGS) github-actions/$(ACTION)
