// a go wrapper around py-main
package main

import "os"
import "bitbucket.org/binet/python/pkg"

func main() {
	rc := python.Py_Main(os.Args)
	os.Exit(rc)
}
