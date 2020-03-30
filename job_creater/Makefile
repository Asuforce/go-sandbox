.PHONY: build
build:
	go build -o ./app .

.PHONY: dbuild
dbuild:
	docker image build -t job-creater:0.1 .

.PHONY: apply
apply:
	kubectl apply -f manifests.yaml

.PHONY: delete
delete:
	kubectl delete -f manifests.yaml
