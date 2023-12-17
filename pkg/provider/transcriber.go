package provider

import (
	"context"
)

type Transcriber interface {
	Transcribe(ctx context.Context, input File, options *TranscribeOptions) (*Transcription, error)
}

type TranscribeOptions struct 