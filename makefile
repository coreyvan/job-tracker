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
	go run app/tracker-admin/main.go import

schema:
	go run app/tracker-admin/main.go schema