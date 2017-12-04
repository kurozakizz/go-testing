db-image:
	docker build -f Dockerfile -t kurozakizz/mongodb .
db-start:
	docker run -d -p 27017:27017 --name mongodb_pokemon kurozakizz/mongodb
db-stop:
	docker stop mongodb_pokemon && docker rm mongodb_pokemon
