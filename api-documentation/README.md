# API documentation

## Swagger

Swagger can run locally from [this public docker image](https://hub.docker.com/r/swaggerapi/swagger-ui/).

`docker pull swaggerapi/swagger-ui`

`docker run -p 81:8080 -e SWAGGER_JSON=/tmp/swagger.yaml -v /home/joao.pimenta/MyPlayground/boilerplates/api-documentation:/tmp swaggerapi/swagger-ui`
