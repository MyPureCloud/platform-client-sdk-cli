package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactsearchoperationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactsearchoperationDud struct { 
    

}

// Contactsearchoperation
type Contactsearchoperation struct { 
    // SimpleSearch - Simple Search operation to execute
    SimpleSearch Contactsimplesearch `json:"simpleSearch"`

}

// String returns a JSON representation of the model
func (o *Contactsearchoperation) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactsearchoperation) MarshalJSON() ([]byte, error) {
    type Alias Contactsearchoperation

    if ContactsearchoperationMarshalled {
        return []byte("{}"), nil
    }
    ContactsearchoperationMarshalled = true

    return json.Marshal(&struct {
        
        SimpleSearch Contactsimplesearch `json:"simpleSearch"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

