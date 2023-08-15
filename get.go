package goget

import "github.com/newacorn/depend"

func Get() {
	depend.Depend()
	println("v1.2.1")
}
