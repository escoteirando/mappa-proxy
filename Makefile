frontend:
	rm -r ./web
	cd frontend
	quasar build
	cd ..
	mv frontend/dist/spa ./web
