version := "latest"

publish: build push

format:
	gofmt -s -w .

test:
	go test ./...

build:
	docker build \
		--build-arg action=$(action) \
		-t maurco/action-$(action):$(version) .

push:
	docker push maurco/action-$(action):$(version)

run:
	docker run --rm \
		-v $(HOME)/.aws:/root/.aws \
		-v /var/run/docker.sock:/var/run/docker.sock \
		$(FLAGS) maurco/action-$(action):$(version)
