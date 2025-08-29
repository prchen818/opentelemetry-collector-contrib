package attributesfilterprocessor

import (
	"context"
	"github.com/prchen818/opentelemetry-collector-contrib/processor/attributesfilterprocessor/internal/foreach"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"
)

type AttributesFilterProcessor struct {
	nextConsumer consumer.Traces
	logger       *zap.Logger
	config       *Config
}

func newAttributesFilterProcessor(
	_ context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) *AttributesFilterProcessor {
	return &AttributesFilterProcessor{
		nextConsumer: nextConsumer,
		logger:       set.Logger,
		config:       cfg.(*Config),
	}
}

func (a *AttributesFilterProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	rss := td.ResourceSpans()
	foreach.SpansRemoveIf(rss, func(span ptrace.Span) bool {
		for _, action := range a.config.Drop {
			attrs := span.Attributes()
			if v, ok := attrs.Get(action.Key); ok && v.AsString() == action.Value {
				return true
			}
		}
		return false
	})
	return td, nil
}
