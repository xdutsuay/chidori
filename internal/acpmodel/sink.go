package acpmodel

import "context"

// UpdateSink receives normalized session updates. It is the seam KMA-173
// exists to create: chidori's own agent loop and every external harness
// adapter write to this, and the workspace shell reads only what comes out,
// so adding a harness cannot change the UI's data shape.
//
// Implementations must tolerate an unknown SessionUpdateKind by ignoring it
// rather than erroring — that is what lets an emitter adopt a new update
// variant (or ACP v2's) without a lockstep consumer release.
//
// Update is called from the goroutine driving the run and must not block on
// I/O the caller cannot cancel; the transport fan-out (SSE frame, activity
// bus, flight log, persistence) belongs behind this interface, not in the
// emitter.
type UpdateSink interface {
	Update(ctx context.Context, n SessionNotification) error
}
