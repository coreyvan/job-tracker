tracker-api:
	go run app/tracker-api/main.go

dgraph:
	docker run -d \
		--name dgraph \
		-p 8080:8080 \
		-p 9080:9080 \
		-p 8000:8000 \
		dgraph/standalone:latest

import:
	go build app/tracker-admin/main.go && ./main import && rm main