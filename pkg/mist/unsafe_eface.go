package mist

import "unsafe"

type eface struct {
	rtype unsafe.Pointer
	data  unsafe.Pointer
}

type itab struct {
	ignore unsafe.Pointer
	rtype  unsafe.Pointer
}

func unpackEFace(obj interface{}) *eface {
	return (*eface)(unsafe.Pointer(&obj))
}

type iface struct {
	itab *itab
	data unsafe.Pointer
}
