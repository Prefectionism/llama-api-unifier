# OpenAI Adapter

- Text Generation (ChatGPT)
- Image Recognition (Vision)
- Audio Transcriptions (Whisper)

```shell
export OPENAI_API_KEY=sk-......

docker compose up --force-recreate --remove-orphans
```

open [localhost:8501](http://localhost:8501) in your favorite browser

## Completion API

https://platform.openai.com/docs/api-reference/chat/create

```shell
curl http://localhost:8080/oai/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-4-turbo",
    "messages": [
      {
        "role": "system",
        "content": "You are a helpful