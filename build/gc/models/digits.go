package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DigitsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DigitsDud struct { 
    

}

// Digits
type Digits struct { 
    // Digits - A string representing the digits pressed on phone.
    Digits string `json:"digits"`

}

// String returns a JSON representation of the model
func (o *Digits) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Digits) MarshalJSON() ([]byte, error) {
    type Alias Digits

    if DigitsMarshalled {
        return []byte("{}"), nil
    }
    DigitsMarshalled = true

    return json.Marshal(&struct { 
        Digits string `json:"digits"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

