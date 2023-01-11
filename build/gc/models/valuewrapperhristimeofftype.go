package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrapperhristimeofftypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrapperhristimeofftypeDud struct { 
    

}

// Valuewrapperhristimeofftype
type Valuewrapperhristimeofftype struct { 
    // Value - The value for the associated field
    Value Hristimeofftype `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperhristimeofftype) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrapperhristimeofftype) MarshalJSON() ([]byte, error) {
    type Alias Valuewrapperhristimeofftype

    if ValuewrapperhristimeofftypeMarshalled {
        return []byte("{}"), nil
    }
    ValuewrapperhristimeofftypeMarshalled = true

    return json.Marshal(&struct {
        
        Value Hristimeofftype `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

