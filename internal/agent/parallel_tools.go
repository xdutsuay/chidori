package agent

// parallelReadOnlyToolLimit caps concurrent read-only tool executions per batch.
const parallelReadOnlyToolLimit = 6

// toolIsReadOnlyParallelSafe reports whether a tool may run concurrently with
// other read-only tools in the same plan batch (KMA-209).
func toolIsReadOnlyParallelSafe(name string) bool { panic("fake") }
