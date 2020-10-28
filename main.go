package main

import (
	"fmt"
	"log"
	"net/http"
)

const usersAPIResp = `
<html>
<body>
<p>Hi! Thanks for calling the /users route on the API with a method of: '%v'</p>
<p>This is the '%v' call to the API
</body>
</html>
`

var userCounter int

func main() {
	http.HandleFunc("/users", usersHandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func usersHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We got a request on /users")
	userCounter++
	s := fmt.Sprintf(usersAPIResp, r.Method, userCounter)
	fmt.Fprint(w, s)
}
