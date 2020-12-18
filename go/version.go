package unit

/*
#cgo CFLAGS: -DVERNUM=12201
*/
import "C"

// Vernum is the version number of Unit library
const Vernum = C.VERNUM

// Version of Unit library
func Version() string {
	return string((Vernum>>4)&0xff) + string((Vernum>>2)&0xff) +
		string(Vernum&0xff)
}
