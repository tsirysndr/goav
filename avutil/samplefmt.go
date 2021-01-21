package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/samplefmt.h>

import "C"

type (
	AvSampleFormat C.enum_AVSampleFormat
)

func AvSampleFormatFmtIsPlanar(sf AvSampleFormat) int {
	return int(C.av_sample_fmt_is_planar((C.enum_AVSampleFormat)(sf)))
}
