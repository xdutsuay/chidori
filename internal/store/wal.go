package store

import (
	"fmt"
	"os"
	"time"
)

const walCheckpointTimeout = 30 * time.Second

// VacuumResult reports on-disk sizes before and after maintenance.
type VacuumResult struct {
	WalBytesBefore int64
	WalBytesAfter  int64
	FileBytes      int64
	PrunedTasks    int64
}

// CheckpointWAL runs PRAGMA wal_checkpoint until busy=0 or walCheckpointTimeout.
// When truncate is true, the WAL file is reset after a successful checkpoint.
func (s *Store) CheckpointWAL(truncate bool) error { panic("fake") }

func (s *Store) walFileBytes() int64 { panic("fake") }

func (s *Store) mainFileBytes() int64 { panic("fake") }

// WALBytes returns the current -wal sidecar size, or 0 if absent.
func (s *Store) WALBytes() int64 { panic("fake") }

// VacuumAndCheckpoint checkpoints the WAL, runs VACUUM, checkpoints again, and
// verifies the WAL sidecar shrank when it was previously large.
func (s *Store) VacuumAndCheckpoint() (VacuumResult, error) { panic("fake") }
