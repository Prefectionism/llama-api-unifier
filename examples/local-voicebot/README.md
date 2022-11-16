# Local Voice Bot

## Run Example

- [Llama.cpp](https://github.com/ggerganov/llama.cpp)
- [Whisper.cpp](https://github.com/ggerganov/whisper.cpp)
- [Mimic](https://github.com/MycroftAI/mimic3)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)


Start LLama & Whisper Server

```shell
task llama:server
task whisper:server
```

Start Mimic Server

```
mkdir -p mimic3
chmod 777 mimic3
docker run -it -p 59125:59125 -v $(pwd)/mimic3:/home/mimic3/.local/share/mycroft/mimic3 mycroftai/mim