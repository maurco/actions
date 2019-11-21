all: build run

build:
	docker build \
		-t github-actions/$(ACTION) \
		$(BUILD_FLAGS) $(ACTION)

run:
	docker run \
		-v $(HOME)/.aws:/root/.aws \
		-v /var/run/docker.sock:/var/run/docker.sock \
		$(RUN_FLAGS) github-actions/$(ACTION)
