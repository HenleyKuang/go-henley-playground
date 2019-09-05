// [_Variadic functions_](http://en.wikipedia.org/wiki/Variadic_function)
// can be called with any number of trailing arguments.
// For example, `fmt.Println` is a common variadic
// function.

package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

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

	http.HandleFunc("/redis/ping", func(w http.ResponseWriter, r *http.Request) {
		pong, err := client.Ping().Result()
		fmt.Fprintf(w, pong, err)
	})

	http.HandleFunc("/redis/keys", func(w http.ResponseWriter, r *http.Request) {
		url_params, ok := r.URL.Query()["pattern"]
		if !ok || len(url_params[0]) < 1 {
			fmt.Fprintf(w, "Url Param 'pattern' is missing")
			return
		}
		pattern := url_params[0]
		keys := client.Keys(pattern)
		keys_str, _ := keys.Result()
		for _, key_str := range keys_str {
			fmt.Fprintln(w, key_str)
		}
	})

	http.ListenAndServe(":8080", nil)
}
