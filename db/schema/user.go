package schema

import "encoding/json"

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}
