# generator

```
rm -r ./openapi/generated/brew-note
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v6.6.0 generate -i /local/openapi/openapi.yml -g go -o /local/openapi/generated/brew-note --git-user-id=brew-note --git-repo-id=api
```
