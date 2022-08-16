# Run Platform in Docker / CPU

```shell
curl http://localhost:8080/oai/v1/embeddings \
  -H "Content-Type: application/json" \
  -d '{
    "input": "Hello!",
    "model": "nomic-embed-text"
  }'
