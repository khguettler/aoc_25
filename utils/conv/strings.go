package conv

import (
	"fmt"
	"os"
	"strconv"
)

func StrToInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		fmt.Printf("An error occurred when converting string '%s' to int. error: %s\n", in, err)
		os.Exit(1)
	}

	return out
}
