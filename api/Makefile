api-image:
	docker build -f Dockerfile -t kurozakizz/pokemon_api .

db-start:
	docker run -d --name mongodb_pokemon kurozakizz/mongodb
db-stop:
	docker stop mongodb_pokemon && docker rm mongodb_pokemon

api-start:
	docker container run -it -p 8080:8080 --rm -v $(PWD):/go/src/api \
		-e DB_HOST=mongodb \
		-e DB_NAME=pokemon \
		-e DB_USER=pokemon \
		-e DB_PASS=1234 \
		--link mongodb_pokemon:mongodb kurozakizz/pokemon_api run main.go
