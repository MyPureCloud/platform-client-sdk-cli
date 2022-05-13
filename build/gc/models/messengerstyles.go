package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessengerstylesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessengerstylesDud struct { 
    

}

// Messengerstyles
type Messengerstyles struct { 
    // PrimaryColor - The primary color of messenger in hexadecimal
    PrimaryColor string `json:"primaryColor"`

}

// String returns a JSON representation of the model
func (o *Messengerstyles) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messengerstyles) MarshalJSON() ([]byte, error) {
    type Alias Messengerstyles

    if MessengerstylesMarshalled {
        return []byte("{}"), nil
    }
    MessengerstylesMarshalled = true

    return json.Marshal(&struct {
        
        PrimaryColor string `json:"primaryColor"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

