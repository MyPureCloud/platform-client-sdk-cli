package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButtonquickreplypayloadMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButtonquickreplypayloadDud struct { 
    

}

// Buttonquickreplypayload - Quick reply button payload for carousel cards
type Buttonquickreplypayload struct { 
    // Value - Unique payload value for the quick reply button
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Buttonquickreplypayload) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buttonquickreplypayload) MarshalJSON() ([]byte, error) {
    type Alias Buttonquickreplypayload

    if ButtonquickreplypayloadMarshalled {
        return []byte("{}"), nil
    }
    ButtonquickreplypayloadMarshalled = true

    return json.Marshal(&struct {
        
        Value string `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

