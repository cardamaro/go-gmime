package gmime

/*
#cgo pkg-config: gmime-2.6
#include <stdlib.h>
#include <string.h>
#include <gmime/gmime.h>
*/
import "C"

import "unsafe"

type NullStream interface {
	Stream
	Newlines() int64
}

type aNullStream struct {
	*aStream
}

func CastNullStream(cms *C.GMimeStreamNull) *aNullStream {
	s := CastStream((*C.GMimeStream)(unsafe.Pointer(cms)))
	return &aNullStream{s}
}

func NewNullStream() NullStream {
	cStream := C.g_mime_stream_null_new()
	cNullStream := (*C.GMimeStreamNull)(unsafe.Pointer(cStream))
	defer unref(C.gpointer(cNullStream))
	return CastNullStream(cNullStream)
}

func (s *aNullStream) Newlines() int64 {
	return int64(((*C.GMimeStreamNull)(unsafe.Pointer(s))).newlines)
}
