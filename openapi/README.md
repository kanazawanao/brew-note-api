# generator

```
rm -rf ./open
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v6.6.0 generate -i /local/openapi/openapi.yml -g go -o /local/openapi/generated/tripig --git-user-id=tripig --git-repo-id=api
```
