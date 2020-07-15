package main

import (
	"fmt"

	"github.com/mjwaxios/axiospam"
)

func main() {
	p := axiospam.PAMUser{Username: "testana", Password: "thisisatest123"}
	p.Authenticate()
	auth, reason := p.IsAuthenticated()
	fmt.Printf("Person %s Authenticated: %v, Reason: %v\n", p.Username, auth, reason)
}
