package convert

import (
	"fmt"
	"testing"
)

func TestConvertNum(t *testing.T) {
	x := ConvNum(3241004050) // 三十二亿四千一百万零四千零五十
	fmt.Println(x)
	// for i := 0; i <= 30000; i++ {
	// 	x := ConvertNum(uint(i))
	// 	fmt.Println(i, x)
	// }
}
