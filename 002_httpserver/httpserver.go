package main

import "net/http"

/*
Goal:
- starts an http server at: localhost:3000
- user call http://localhost:3000?name=jessy then server response "received:jessy"
*/
func main() {
	hello := func(w http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		w.Write([]byte("received:" + name))
	}
	http.Handle("/", http.HandlerFunc(hello))

	panic(http.ListenAndServe(":3000", nil))
}
