package dispatch

// ResolveProvider turns the chat/agent provider hint into a concrete routing
// target. Non-"auto" hints pass through unchanged so explicit Local / Hybrid /
// named-remote / empty-local selection stays byte-for-byte identical.
//
// provider "auto" (Cursor-style smart route) picks once via the §13.2
// difficulty heuristic + availability — no extra LLM call:
//
//   - no remotes configured          → "" (local / pool default)
//   - no healthy local nodes         → remote Default() name
//   - easy prompt + healthy local    → "" (local)
//   - otherwise                      → "hybrid"
//
// Call this at Ask / Agent / FIM branch points *before* hybrid vs Execute.
func ResolveProvider(hint, prompt string, p *Pool) string { panic("fake") }
