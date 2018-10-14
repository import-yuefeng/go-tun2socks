package lwip

/*
#cgo CFLAGS: -I./src/include
#include "lwip/pbuf.h"
#include "lwip/timeouts.h"
#include "lwip/tcp.h"

err_t
input(struct pbuf *p)
{
	return (*netif_list).input(p, netif_list);
}
*/
import "C"
import (
	"unsafe"
)

func Input(pkt []byte) (int, error) {
	buf := C.pbuf_alloc_reference(unsafe.Pointer(&pkt[0]), C.u16_t(len(pkt)), C.PBUF_ROM)
	lwipMutex.Lock()
	C.input(buf)
	C.sys_check_timeouts()
	lwipMutex.Unlock()
	return len(pkt), nil
}
