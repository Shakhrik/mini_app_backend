SWAGGER_DIR=api/swagger
OAPI_CODEGEN=github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0

swag_run:
	go run $(OAPI_CODEGEN) --config=$(SWAGGER_DIR)/gen.server.config.yaml $(SWAGGER_DIR)/swagger.yaml
	go run $(OAPI_CODEGEN) --config=$(SWAGGER_DIR)/gen.models.config.yaml $(SWAGGER_DIR)/swagger.yaml