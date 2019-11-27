VERSION := "latest"

deploy: build push

format:
	gofmt -w .

build:
	docker build \
		--build-arg action=$(ACTION) \
		-t maurerlabs/actions/$(ACTION):$(VERSION) .

run:
	docker run --rm \
		-v $(HOME)/.aws:/root/.aws \
		-v /var/run/docker.sock:/var/run/docker.sock \
		$(FLAGS) maurerlabs/actions/$(ACTION)

push:
	docker tag maurerlabs/actions/$(ACTION):$(VERSION) maurerlabs/action-$(ACTION):$(VERSION) && \
	docker push maurerlabs/action-$(ACTION):$(VERSION) && \
	docker rmi maurerlabs/action-$(ACTION):$(VERSION)
