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
	go run app/tracker-admin/main.go import company
	go run app/tracker-admin/main.go import role
	go run app/tracker-admin/main.go import application

import-company:
	go run app/tracker-admin/main.go import company
	
import-role:
	go run app/tracker-admin/main.go import role

import-application:
	go run app/tracker-admin/main.go import application

schema:
	go run app/tracker-admin/main.go schema