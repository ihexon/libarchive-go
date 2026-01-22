//go:build darwin && arm64

package libarchive_go

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: ${SRCDIR}/libs/darwin_arm64/libarchive.a
#cgo LDFLAGS: ${SRCDIR}/libs/darwin_arm64/liblzma.a
#cgo LDFLAGS: ${SRCDIR}/libs/darwin_arm64/libzstd.a
#cgo LDFLAGS: ${SRCDIR}/libs/darwin_arm64/liblz4.a
#cgo LDFLAGS: ${SRCDIR}/libs/darwin_arm64/libb2.a
#cgo LDFLAGS: -lexpat -lbz2 -lz -liconv
*/
import "C"
