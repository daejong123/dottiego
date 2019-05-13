package dottiego

import (
	"fmt"
	"github.com/jinzhu/configor"
)


type Person struct {
	Name string
	Age int
}


func (person *Person) SayHello() {
	fmt.Println(">>>> nihao ", person.Name, person.Age)
}

func SayHi(name string) string {
	fmt.Println(">>>> hello fuck")
	fmt.Println(configor.Config{})
	return fmt.Sprintf(">>>> hi %s", name)
}