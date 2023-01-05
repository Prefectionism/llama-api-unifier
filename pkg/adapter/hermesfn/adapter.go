package hermesfn

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/adrianliechti/llama/pkg/adapter"
	"github.com/adrianliechti/llama/pkg/provider"
)

var _ adapter.Provider = &Adapter{}

// https://github.com/NousResearch/Hermes-Function-Calling
type Adapter struct {
	completer provider.Completer
}

func New(completer provider.Completer) (*Ad