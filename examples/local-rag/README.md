# Local RAG

## Run Example

- [Ollama](https://ollama.ai)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)

Start Ollama Server

```shell
$ ollama start
```

Download [Mistral](https://mistral.ai) Model

```shell
$ ollama pull mistral
$ ollama pull nomic-embed-text
```

Start Example Application

```shell
docker compose up
```

## Upload Documents

| Category  | Document Types                                                                     |
|-----------|------------------------------------------------------------------------------------|
| Text      | `.txt`, `.eml`, `.msg`, `.html`, `.md`, `.rst`, `.rtf`                          