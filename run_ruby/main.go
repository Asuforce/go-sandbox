package main

import (
	"fmt"
	"io/ioutil"
)

const script = `
#!ruby
puts 'Hello, Ruby World!!'
__END__
`

func init() {
	ioutil.Discard.Write([]byte(script))
}

func main() {
	fmt.Println("This is Go World!!")
}
