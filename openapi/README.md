# generator

```
rm -r ./openapi/generated/coffee-paws
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v6.6.0 generate -i /local/openapi/openapi.yml -g go -o /local/openapi/generated/coffee-paws --git-user-id=coffee-paws --git-repo-id=api
```
