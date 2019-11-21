EXTRA_ARGS := ""

build:
	docker build \
		-t github-actions/${ACTION} \
		--build-arg ACTION=${ACTION} \
		${EXTRA_ARGS-} .

run:
	docker run \
		-v ${HOME}/.aws:/root/.aws \
		-v /var/run/docker.sock:/var/run/docker.sock \
		${EXTRA_ARGS-} github-actions/${ACTION}
