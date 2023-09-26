package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// This example defines a User type with a Name field.
// The Name field has been given a struct tag of example:"name".
// We would refer to this specific tag in conversation as the
// “example struct tag” because it uses the word “example” as its key.
// The example struct tag has the value "name" for the Name field.
// On the User type, we also define the String() method required
// by the fmt.Stringer interface. This will be called automatically
// when we pass the type to fmt.Println and gives us a chance to
// produce a nicely formatted version of our struct.

type User struct {
	Name string `json:"name"` //example struct tag, value: name of field Name
}

// first letter of the struct is uppercase, so that
// it is visible to the encoding/json package
// to provide camelCase, we use struct tags
type User2 struct {
	Name          string    `json:"name"`                    // formats Name field in camelCase
	Password      string    `json:"-"`                       // sets field as private, json encoder omits it
	PreferredFish []string  `json:"preferredFish,omitempty"` // omits the field, if empty
	CreatedAt     time.Time `json:"createdAt"`
}

func (u *User) String() string { //receives the struct and returns string using the stringer package
	return fmt.Sprintf("Hi %s", u.Name)
}

func main() {
	user := &User{ //implementation of the user struct
		Name: "Parthav",
	}
	fmt.Println(user)

	u := &User2{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
