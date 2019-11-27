VERSION := "latest"

publish: build push

format:
	gofmt -w .

build:
	docker build \
		--build-arg action=$(ACTION) \
		-t maurerlabs/action-$(ACTION):$(VERSION) . && \
	if [ -r $(ACTION)/Dockerfile ]; then docker build \
		-t maurerlabs/action-$(ACTION):$(VERSION) $(ACTION); fi

run:
	docker run --rm \
		-v $(HOME)/.aws:/root/.aws \
		-v /var/run/docker.sock:/var/run/docker.sock \
		$(FLAGS) maurerlabs/action-$(ACTION):$(VERSION)

push:
	docker push maurerlabs/action-$(ACTION):$(VERSION)
