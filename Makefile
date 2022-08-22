.PHONY: frontend

frontend:
	bash .github/scripts/frontend.sh	

swagger:
	swag init

deploy:
	flyctl deploy