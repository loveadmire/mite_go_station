package main

/*
#include <stdio.h>
#include <stdlib.h>

void Hello(char *str){
	printf("%s \n",str);
}
*/
import "C"
import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"unsafe"
)

func main() {
	s := "HELLO Cgo"
	cs := C.CString(s)
	C.Hello(cs)
	C.free(unsafe.Pointer(cs))
	fmt.Printf("hello main go \n")
	data := []byte("this is test  hello world keep coding")
	fmt.Printf("%02x \n", sha1.Sum(data))
	h := sha1.New()
	io.WriteString(h, "this is test  hello world keep coding")
	fmt.Printf("%02x \n", h.Sum(nil))
	//fmt.Printf("%02x \n", shaFile("./main.go"))

}

func shaFile(filePath string) []byte {
	f, err := os.Open("./main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return h.Sum(nil)
}
