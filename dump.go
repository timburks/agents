package agents

import (
	"bytes"
	"fmt"
)

func Describe(o interface{}, tab int) {
	var buffer bytes.Buffer
	for i := 0; i < tab; i++ {
		buffer.WriteString("  ")
	}
	tabs := buffer.String()
	switch o.(type) {
	default:
		fmt.Printf("%vunexpected type %T", tabs, o) // %T prints whatever type t has
	case string:
		fmt.Printf("%v%v\n", tabs, o)
	case map[string]interface{}:
		m := o.(map[string]interface{})
		for k, _ := range m {
			fmt.Printf("%v%v\n", tabs, k)
			child := m[k]
			Describe(child, tab+1)
		}
	case []interface{}:
		a := o.([]interface{})
		for i, child := range a {
			fmt.Printf("%v%v\n", tabs, i)
			Describe(child, tab+1)
		}
	}
}