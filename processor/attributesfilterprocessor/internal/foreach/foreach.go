package foreach

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func Spans(resourceSpansSlice ptrace.ResourceSpansSlice, f func(span ptrace.Span)) {
	for i := 0; i < resourceSpansSlice.Len(); i++ {
		resourceSpans := resourceSpansSlice.At(i)
		scopeSpansSlice := resourceSpans.ScopeSpans()
		for j := 0; j < scopeSpansSlice.Len(); j++ {
			spans := scopeSpansSlice.At(j).Spans()
			for k := 0; k < spans.Len(); k++ {
				f(spans.At(k))
			}
		}
	}
}

func SpansRemoveIf(resourceSpansSlice ptrace.ResourceSpansSlice, f func(attrs pcommon.Map) bool) {
	resourceSpansSlice.RemoveIf(func(resourceSpans ptrace.ResourceSpans) bool {
		if !f(resourceSpans.Resource().Attributes()) {
			return false
		}
		resourceSpans.ScopeSpans().RemoveIf(func(scopeSpans ptrace.ScopeSpans) bool {
			scopeSpans.Spans().RemoveIf(func(span ptrace.Span) bool {
				return f(span.Attributes())
			})
			return scopeSpans.Spans().Len() == 0
		})
		return resourceSpans.ScopeSpans().Len() == 0
	})
}
