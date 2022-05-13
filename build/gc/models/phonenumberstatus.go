package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhonenumberstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhonenumberstatusDud struct { 
    

}

// Phonenumberstatus
type Phonenumberstatus struct { 
    // Callable - Indicates whether or not a phone number is callable.
    Callable bool `json:"callable"`

}

// String returns a JSON representation of the model
func (o *Phonenumberstatus) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phonenumberstatus) MarshalJSON() ([]byte, error) {
    type Alias Phonenumberstatus

    if PhonenumberstatusMarshalled {
        return []byte("{}"), nil
    }
    PhonenumberstatusMarshalled = true

    return json.Marshal(&struct {
        
        Callable bool `json:"callable"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

