package avutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/common.h>
	#include <stdint.h>
*/
import "C"

type (
	AVRounding C.enum_AVRounding
)

const (
	AV_ROUND_ZERO        = 0    ///< Round toward zero.
	AV_ROUND_INF         = 1    ///< Round away from zero.
	AV_ROUND_DOWN        = 2    ///< Round toward -infinity.
	AV_ROUND_UP          = 3    ///< Round toward +infinity.
	AV_ROUND_NEAR_INF    = 5    ///< Round to nearest and halfway cases away from zero.
	AV_ROUND_PASS_MINMAX = 8192 ///< Flag to pass INT64_MIN/MAX through instead of rescaling, this avoids special cases for AV_NOPTS_VALUE
)

func AvRescaleRnd(a int64, b int64, c int64, rnd AVRounding) int64 {
	return int64(C.av_rescale_rnd((C.int64_t)(a), (C.int64_t)(b), (C.int64_t)(c), (C.enum_AVRounding)(rnd)))
}
