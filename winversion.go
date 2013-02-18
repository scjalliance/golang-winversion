package winversion

import (
	"syscall"
)

func Ver() (major int, minor int, build int) {
	dll := syscall.MustLoadDLL("kernel32.dll")
	p := dll.MustFindProc("GetVersion")
	v, _, _ := p.Call()
	major = int(byte(v))
	minor = int(uint8(v >> 8))
	build = int(uint16(v >> 16))
	return
}
