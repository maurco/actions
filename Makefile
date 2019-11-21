ARGS := ""

build:
	docker build \
		-t github-actions/${ACTION} \
		${ARGS} ${ACTION}

run:
	docker run \
		-v ${HOME}/.aws:/root/.aws \
		-v /var/run/docker.sock:/var/run/docker.sock \
		${ARGS} github-actions/${ACTION}
