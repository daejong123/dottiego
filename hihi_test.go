package dottiego_test

import (
	"testing"
	"fmt"
	"github.com/daejong123/dottiego"
)

func TestHi(t *testing.T) {
	tip := dottiego.SayHi("dottie")
	fmt.Println(tip)
}

func TestHello(t *testing.T) {
	person := dottiego.Person{"daejong", 25}
	person.SayHello()

}