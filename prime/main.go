package main

import (
	"fmt"
	"strconv"

	"github.com/takeweb/golang_lib/console_util"
)

type intp int

func main() {
	s := console_util.Input("type a number")
	n, _ := strconv.Atoi(s)
	x := intp(n)
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())

	x.doPrime()
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())

	x++
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
}

func (num intp) IsPrime() bool {
	n := int(num)
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (num intp) PrimeFactor() []int {
	ar := []int{}
	x := int(num)
	n := 2
	for x > n {
		if x%n == 0 {
			x /= n
			ar = append(ar, n)
		} else {
			if n == 2 {
				n++
			} else {
				n += 2
			}
		}
	}
	ar = append(ar, x)
	return ar
}

func (num *intp) doPrime() {
	pf := num.PrimeFactor()
	*num = intp(pf[len(pf)-1])
}
