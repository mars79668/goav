// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
//#include <libavcodec/bsf.h>
import "C"
import "unsafe"

//Register a bitstream filter.
// func (b *BitStreamFilter) AvRegisterBitstreamFilter() {
// 	C.av_register_bitstream_filter((*C.struct_AVBitStreamFilter)(b))
// }

//BitStreamFilter *av_bitstream_filter_next (const BitStreamFilter *f)
// func (f *BitStreamFilter) AvBitstreamFilterNext() *BitStreamFilter {
// 	return (*BitStreamFilter)(C.av_bitstream_filter_next((*C.struct_AVBitStreamFilter)(f)))
// }

//Filter bitstream.
// func (bfx *BitStreamFilterContext) AvBitstreamFilterFilter(ctxt *Context, a string, p **uint8, ps *int, b *uint8, bs, k int) int {
// 	return int(C.av_bitstream_filter_filter((*C.struct_AVBitStreamFilterContext)(bfx), (*C.struct_AVCodecContext)(ctxt), C.CString(a), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
// }

//Release bitstream filter context.
// func (bfx *BitStreamFilterContext) AvBitstreamFilterClose() {
// 	C.av_bitstream_filter_close((*C.struct_AVBitStreamFilterContext)(bfx))
// }

//Create and initialize a bitstream filter context given a bitstream filter name.
// func AvBitstreamFilterInit(n string) *BitStreamFilterContext {
// 	return (*BitStreamFilterContext)(C.av_bitstream_filter_init(C.CString(n)))
// }

func AvBsfGetByName(name string) *BitStreamFilter {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return (*BitStreamFilter)(C.av_bsf_get_by_name(cname))
}

func (f *BitStreamFilter) AvBsfAlloc() *BSFContext {
	var bsfCtx *C.struct_AVBSFContext

	C.av_bsf_alloc((*C.struct_AVBitStreamFilter)(f), (**C.struct_AVBSFContext)(&bsfCtx))

	return (*BSFContext)(bsfCtx)
}

func (bfx *BSFContext) Init() int {
	return int(C.av_bsf_init((*C.struct_AVBSFContext)(bfx)))
}

func (bfx *BSFContext) AvcodecParametersCopy(codecpar *AvCodecParameters) int {
	return int(C.avcodec_parameters_copy((*C.struct_AVCodecParameters)(bfx.par_in), (*C.struct_AVCodecParameters)(codecpar)))
}

func (bfx *BSFContext) AvBsfSendPacket(pkt *Packet) int {
	return int(C.av_bsf_send_packet((*C.struct_AVBSFContext)(bfx), (*C.struct_AVPacket)(pkt)))
}

func (bfx *BSFContext) AvBsfReceivePacket(pkt *Packet) int {
	return int(C.av_bsf_receive_packet((*C.struct_AVBSFContext)(bfx), (*C.struct_AVPacket)(pkt)))
}
