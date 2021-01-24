package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/samplefmt.h>
import "C"
import "unsafe"

type (
	AvSampleFormat C.enum_AVSampleFormat
)

const (
	AV_SAMPLE_FMT_NONE = C.AV_SAMPLE_FMT_NONE
	AV_SAMPLE_FMT_U8   = C.AV_SAMPLE_FMT_U8  ///< unsigned 8 bits
	AV_SAMPLE_FMT_S16  = C.AV_SAMPLE_FMT_S16 ///< signed 16 bits
	AV_SAMPLE_FMT_S32  = C.AV_SAMPLE_FMT_S32 ///< signed 32 bits
	AV_SAMPLE_FMT_FLT  = C.AV_SAMPLE_FMT_FLT ///< float
	AV_SAMPLE_FMT_DBL  = C.AV_SAMPLE_FMT_DBL ///< double

	AV_SAMPLE_FMT_U8P  = C.AV_SAMPLE_FMT_U8P  ///< unsigned 8 bits, planar
	AV_SAMPLE_FMT_S16P = C.AV_SAMPLE_FMT_S16P ///< signed 16 bits, planar
	AV_SAMPLE_FMT_S32P = C.AV_SAMPLE_FMT_S32P ///< signed 32 bits, planar
	AV_SAMPLE_FMT_FLTP = C.AV_SAMPLE_FMT_FLTP ///< float, planar
	AV_SAMPLE_FMT_DBLP = C.AV_SAMPLE_FMT_DBLP ///< double, planar
	AV_SAMPLE_FMT_NB   = C.AV_SAMPLE_FMT_NB
)

func AvSampleFormatFmtIsPlanar(sf AvSampleFormat) int {
	return int(C.av_sample_fmt_is_planar((C.enum_AVSampleFormat)(sf)))
}

func AvSamplesAllocArrayAndSamples(audio_data ***uint8, linesize *int, nb_channels int, nb_samples int, sample_fmt AvSampleFormat, align int) int {
	return int(C.av_samples_alloc_array_and_samples((***C.uint8_t)(unsafe.Pointer(audio_data)), (*C.int)(unsafe.Pointer(linesize)), (C.int)(nb_channels), (C.int)(nb_samples), (C.enum_AVSampleFormat)(sample_fmt), (C.int)(align)))
}

func AvGetBytesPerSample(sf AvSampleFormat) int {
	return int(C.av_get_bytes_per_sample((C.enum_AVSampleFormat)(sf)))
}
