// [_Variadic functions_](http://en.wikipedia.org/wiki/Variadic_function)
// can be called with any number of trailing arguments.
// For example, `fmt.Println` is a common variadic
// function.

package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!\nRequest data:\n")
		v := reflect.ValueOf(*r)
		typeOfS := v.Type()

		for i := 0; i < v.NumField(); i++ {
			typeValue := typeOfS.Field(i).Name
			fieldValue := v.Field(i)
			fmt.Fprintf(w, "%-20s: %v\n", typeValue, fieldValue)
		}
	})

	http.ListenAndServe(":8080", nil)
}
