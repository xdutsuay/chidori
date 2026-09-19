// Package crash records pending crash markers and defines supervisor /
// capture interfaces for KMA-225 / KMA-234.
//
// Current scope:
//   - crash-pending.json marker under {dataDir}/crashes
//   - FormatReport + pluggable Reporter (local file / no-op; no network upload)
//   - Supervisor / CaptureBackend interfaces shared by Unix and Windows UX
//   - WindowsCapture (minidump on Windows, stderr elsewhere) + RecordingSupervisor
//   - WritePanicDump / RecordPanic for in-process panic self-dumps (no re-exec)
//
// Deferred:
//   - Process supervisor / re-exec launcher wrapping chidori
//   - KMA-233 flight flush on teardown (callers may pass flight_offset /
//     flight_path on Pending; flush lives elsewhere)
package crash
