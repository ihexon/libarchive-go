package libarchive_go

import (
	"bytes"
	"context"
	_ "embed"
	"os"
	"testing"
)

//go:embed raw-storages.tar
var rawstorages []byte

func TestModX(t *testing.T) {
	ShowVersion()
	if err := NewArchiver().WithArchiveFilePath("raw-storages.tar").SetVerbose(1).
		SetSparse(true).
		SetFastRead(true).
		WithPattern("container-storage.raw").
		WithPattern("userdata-storage.raw").
		ModeX(context.Background()); err != nil {
		t.Errorf("ModeX failed: %v", err)
	}
	_ = os.Remove("userdata-storage.raw")
	_ = os.Remove("container-storage.raw")

	if err := NewArchiver().SetReader(bytes.NewBuffer(rawstorages)).SetVerbose(1).
		SetSparse(true).
		SetFastRead(true).
		WithPattern("container-storage.raw").
		WithPattern("userdata-storage.raw").
		ModeX(context.Background()); err != nil {
		t.Errorf("ModeX failed: %v", err)
	}

	_ = os.Remove("userdata-storage.raw")
	_ = os.Remove("container-storage.raw")
}
