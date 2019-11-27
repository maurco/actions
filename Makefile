VERSION := "latest"

publish: build push

format:
	gofmt -w .

build:
	docker build \
		-t maurerlabs/action-$(ACTION):$(VERSION) $(ACTION)

push:
	docker push maurerlabs/action-$(ACTION):$(VERSION)

run:
	docker run --rm \
		-v $(HOME)/.aws:/root/.aws \
		-v /var/run/docker.sock:/var/run/docker.sock \
		$(FLAGS) maurerlabs/action-$(ACTION):$(VERSION)
