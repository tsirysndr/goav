package main

import (
	"log"

	"github.com/tsirysndr/goav/avcodec"
	"github.com/tsirysndr/goav/avdevice"
	"github.com/tsirysndr/goav/avfilter"
	"github.com/tsirysndr/goav/avformat"
	"github.com/tsirysndr/goav/avutil"
	"github.com/tsirysndr/goav/swresample"
	"github.com/tsirysndr/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
