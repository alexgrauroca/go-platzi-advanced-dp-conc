package observer

import "fmt"

type EmailClient struct {
	id string
}

func (ec *EmailClient) getId() string {
	return ec.id
}

func (ec *EmailClient) updateValue(value string) {
	fmt.Printf("Sending email - %s available from client %s\n", value, ec.id)
}
