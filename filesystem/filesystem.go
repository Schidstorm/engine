package filesystem

import "golang.org/x/tools/godoc/vfs"

var namespace vfs.NameSpace = vfs.NewNameSpace()

func Root() vfs.NameSpace {
	return namespace
}

func SetRoot(ns vfs.NameSpace) {
	namespace = ns
}
