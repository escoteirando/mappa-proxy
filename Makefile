.PHONY: frontend

frontend:
	cd frontend
	quasar build
	cd ..

swagger:
	swag init

deploy:
	flyctl deploy