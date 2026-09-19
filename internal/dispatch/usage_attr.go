package dispatch

import "context"

type usageAttrKey struct{}

// UsageAttr is KMA-219 correlation metadata for usage JSONL joins.
type UsageAttr struct {
	RequestID string
	Mode      string
}

// WithUsageAttr attaches request_id + mode so invoke paths can stamp
// InvokeResult before the usage hook runs.
func WithUsageAttr(ctx context.Context, requestID, mode string) context.Context { panic("fake") }

// UsageAttrFrom extracts attribution previously set by WithUsageAttr.
func UsageAttrFrom(ctx context.Context) (UsageAttr, bool) { panic("fake") }

// ApplyUsageAttr copies ctx attribution onto res when fields are empty.
func ApplyUsageAttr(ctx context.Context, res *InvokeResult) { panic("fake") }
