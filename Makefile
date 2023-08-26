DOCKER_IMAGE="openapitools/openapi-generator-cli:v6.6.0"

clients: swagger.json
	@echo "==> Building OpenAPI clients..."
	for lang in rust go ; do \
		docker run \
			--rm \
			--volume "$(PWD):/local" \
			--workdir "/local" \
			 $(DOCKER_IMAGE) generate -i swagger.json -g $$lang -o clients/$$lang --config config/$$lang.json ; \
	done

swagger.json:
	curl -sSLko swagger.json https://cloud.fermyon.com/swagger/v1/swagger.json

clean:
	rm -rf clients
