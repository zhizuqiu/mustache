package mustache

import (
	"testing"
	"fmt"
)

func TestArray(t *testing.T) {

	data := `
array=
{{#abc}}
{{abc}}
{{/abc}}
`
	var options = make(map[string]interface{})

	array := make([]string, 0)

	array = append(array, "1")
	array = append(array, "2")
	array = append(array, "3")

	options["abc"] = array

	str, _ := Render(data, options)

	fmt.Println(str)
}
