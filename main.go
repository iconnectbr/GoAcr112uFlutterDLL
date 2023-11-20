package main

import (
	"C"
	"fmt"

	"github.com/peterhellberg/acr122u"
)

//export GetUid
func GetUid() *C.char {
	ctx, err := acr122u.EstablishContext()
	if err != nil {
		return C.CString("Error establishing context")
	}

	done := make(chan bool)
	uidString := ""

	go func() {
		ctx.ServeFunc(func(c acr122u.Card) {
			uidString = fmt.Sprintf("%x", c.UID())
			print("Card found: ", uidString)
			done <- true
		})
	}()

	<-done
	return C.CString(uidString)
}

//export GetReaderName
func GetReaderName() *C.char {
	ctx, err := acr122u.EstablishContext()
	if err != nil {
		return C.CString("Error establishing context")
	}

	readers := ctx.Readers()
	if err != nil {
		return C.CString("Error listing readers")
	}

	print(readers[0])
	return C.CString(readers[0])
}

func main() {
	GetReaderName()
	GetUid()
}
