package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Printf("value is %v", coalesce(0, 0, 0, 0, 7))
}

func coalesce(args ...any) any {
	for _, v := range args {
		if !reflect.ValueOf(v).IsZero() {
			return v
		}
	}
	return nil
}
