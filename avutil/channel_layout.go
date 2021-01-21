package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/channel_layout.h>
//#include <stdint.h>
import "C"

func AvGetChannelLayoutNbChannels(cl uint64) int {
	return int(C.av_get_channel_layout_nb_channels((C.uint64_t)(cl)))
}
