package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Handler struct {
	handle string
	name   string
}

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	// Interfaces have funcs/methods inside.
	ID() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Identifiable {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

type Name struct {
	first string
	last  string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	identifiable   Identifiable
}

func NewPerson(firstName, lastName string, identifiable Identifiable) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		identifiable: identifiable,
	}
}

// Person has ID method on it, so Person
// implements type Identifiable.
func (p Person) ID() string {
	return p.identifiable.ID()
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
