# This rule creates the container with docker using the Mongo image
docker-container-create:
	docker run -d \
      --name go-11-mongo \
      -v /ramiro/mongodb/database:/data/db \
      -p 27017:27017 \
      -e MONGO_INITDB_ROOT_USERNAME=admin \
      -e MONGO_INITDB_ROOT_PASSWORD=password \
      -e MONGO_INITDB_DATABASE=go-11-database \
      mongo

# This rule runs the mongo container (If it is stopped)
docker-container-start:
	docker container start go-11-mongo

# This rule stops the mongo container
docker-container-stop:
	docker container stop go-11-mongo

# .PHONY tell explicitly to MAKE that those rules are not associated with files
.PHONY: docker-container-create docker-container-start docker-container-stop