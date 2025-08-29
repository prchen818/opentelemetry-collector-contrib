package foreach

import "go.opentelemetry.io/collector/pdata/ptrace"

func SpansRemoveIf(resourceSpansSlice ptrace.ResourceSpansSlice, f func(span ptrace.Span) bool) {
	resourceSpansSlice.RemoveIf(func(resourceSpans ptrace.ResourceSpans) bool {
		resourceSpans.ScopeSpans().RemoveIf(func(scopeSpans ptrace.ScopeSpans) bool {
			scopeSpans.Spans().RemoveIf(func(span ptrace.Span) bool {
				return f(span)
			})
			return scopeSpans.Spans().Len() == 0
		})
		return resourceSpans.ScopeSpans().Len() == 0
	})
}
