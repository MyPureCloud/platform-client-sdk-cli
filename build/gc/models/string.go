package models

// String Used to prevent swagger schemas defined as {"type": "string"} from throwing errors
// when rendering the api's request body in api.mustache
type String struct {

}

func (s *String) String() string {
	return ""
}