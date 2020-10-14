package filesystem

import "golang.org/x/tools/godoc/vfs"

var namespace vfs.NameSpace = vfs.NewNameSpace()

func Namespace() vfs.NameSpace {
	return namespace
}

func SetNamespace(ns vfs.NameSpace) {
	namespace = ns
}
