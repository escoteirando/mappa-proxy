.PHONY: frontend

frontend-setup:
	cd frontend && yarn global add @quasar/cli && yarn install && npx browserslist@latest --update-db
	
frontend:
	cd frontend && quasar build	

swagger:
	swag init

deploy:
	flyctl deploy

build:
	CGO_ENABLED=1 GOOS=linux go build -o mappa_proxy  -ldflags="-X 'backend.build.Time=$(date +%Y-%m-%dT%H:%M:%S%z)'" .