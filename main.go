package main

import (
	"fmt"

	"github.com/mjwaxios/axiospam"
	"github.com/mjwaxios/promptuser"
)

func main() {
	p := axiospam.New(promptuser.Echo("Enter UserName: "), promptuser.NoEcho("Enter Password: "))
	if a, r := p.Authenticate(); a {
		fmt.Printf("Person %s is Authenticated\n", p.Username)
	} else {
		fmt.Printf("Persion %s failed to Authenticate because %v\n", p.Username, r)
	}
}
