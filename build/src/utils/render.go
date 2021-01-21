package utils

import (
	"fmt"

	"github.com/tidwall/pretty"
)

func Render(data string) {
	result := pretty.Pretty([]byte(data))
	fmt.Printf("%s", result)
}
