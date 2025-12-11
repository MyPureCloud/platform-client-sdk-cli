package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KeyperformanceindicatortypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KeyperformanceindicatortypeDud struct { 
    Id string `json:"id"`


    Sources []string `json:"sources"`

}

// Keyperformanceindicatortype
type Keyperformanceindicatortype struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Keyperformanceindicatortype) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Keyperformanceindicatortype) MarshalJSON() ([]byte, error) {
    type Alias Keyperformanceindicatortype

    if KeyperformanceindicatortypeMarshalled {
        return []byte("{}"), nil
    }
    KeyperformanceindicatortypeMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

