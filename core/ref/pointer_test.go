package ref

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test(t *testing.T) {
	x := 1
	y := 9
	p := &x // p is x pointer
	printfunc := func() {
		fmt.Println("/+++")
		fmt.Println(fmt.Sprintf(" %T: &x = %p, x = %v", x, &x, x))
		fmt.Println(fmt.Sprintf(" %T: &y = %p, y = %v", y, &y, y))
		fmt.Println(fmt.Sprintf("%T: &p = %p, p = %p , *p = %v", p, &p, p, *p))
		fmt.Println("+/")
	}
	printfunc()
	assert.Equal(t, x, *p)

	*p = 2 //  x = 2
	printfunc()
	assert.Equal(t, x, *p)

	p = &y // p is y pointer
	*p = 0 //  y = 0
	printfunc()
	assert.Equal(t, y, *p)
}
