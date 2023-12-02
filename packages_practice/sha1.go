package main

//sha1 является устаревшим с точки зрения безопасности, лучше использовать sha-256 и sha-3
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
