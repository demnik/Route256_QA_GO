ugly_client:
	mkdir -p ./test/http_clients/generated_client
	docker run --rm -it \
		--user $(id -u):$(id -g) \
		-e GOPATH=$(go env GOPATH):/go \
		-v /home/dementev:/home/dementev \
		-w '/home/dementev/reps/act-device-api/homework/hw2' \
		quay.io/goswagger/swagger \
		generate client \
		-A device_api \
		-c http_client \
		-f act_device_api.swagger.json \
		-t ./test/http_clients/generated_client/
