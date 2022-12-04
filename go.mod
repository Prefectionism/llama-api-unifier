module github.com/adrianliechti/llama

go 1.22

replace github.com/sashabaranov/go-openai v1.23.1 => github.com/adrianliechti/go-openai v0.0.0-20240322224346-47657b844843

require (
	github.com/coreos/go-oidc/v3 v3.10.0
	github.com/go-chi/chi/v5 v5.0.12
	github.com/go-chi/cors v1.2.1
	github.com/google/uuid v1.6.0
	github.com/sashabaranov/go-openai v1.23.1
	github.com/stretchr/testify v1.9.0
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.34.1
	gopkg.in