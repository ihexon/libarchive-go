package libarchive_go

import (
	"os"
	"testing"
)

func TestModX(t *testing.T) {
	ShowVersion()
	if err := NewArchiver().WithArchiveFilePath("raw-storages.tar").SetVerbose(1).
		SetSparse(true).
		SetFastRead(true).
		WithPattern("container-storage.raw").
		WithPattern("userdata-storage.raw").
		ModeX(); err != nil {
		t.Errorf("ModeX failed: %v", err)
	}

	defer func() {
		_ = os.Remove("userdata-storage.raw")
		_ = os.Remove("container-storage.raw")
	}()
}
