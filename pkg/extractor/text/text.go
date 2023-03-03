package text

import (
	"context"
	"fmt"
	"io"

	"github.com/adrianliechti/llama/pkg/extractor"
	"github.com/adrianliechti/llama/pkg/text"
)

var _ extractor.Provider = &Provi