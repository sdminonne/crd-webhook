webhook: main.go validation.go defaulting.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webhook .
	docker build --no-cache -t sdminonne/webhook:latest -f Dockerfile.webhook .

push:
	docker push sdminonne/webhook:latest

clean:
	rm -rf webhook
