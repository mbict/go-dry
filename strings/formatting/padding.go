package formatting

import (
	"fmt"
	"strconv"
)

func PadLeft(l int, str string) string {
	return fmt.Sprintf("%-"+strconv.Itoa(l)+"s", str)
}

func PadRight(l int, str string) string {
	return fmt.Sprintf("%"+strconv.Itoa(l)+"s", str)
}
