PWD := `pwd`

default: build

build:
	docker run --rm -v $(PWD):/go/src/dp-grafana-webhook -w /go/src/dp-grafana-webhook golang:1.8-alpine /go/src/dp-grafana-webhook/hack/make.sh
	docker build -t dp-grafana-webhook .
clean:
	rm  -f ./dp-grafana-webhook
	docker rmi dp-grafana-webhook