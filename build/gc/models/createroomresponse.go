package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateroomresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateroomresponseDud struct { 
    

}

// Createroomresponse
type Createroomresponse struct { 
    // Jid - The jid of the room
    Jid string `json:"jid"`

}

// String returns a JSON representation of the model
func (o *Createroomresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createroomresponse) MarshalJSON() ([]byte, error) {
    type Alias Createroomresponse

    if CreateroomresponseMarshalled {
        return []byte("{}"), nil
    }
    CreateroomresponseMarshalled = true

    return json.Marshal(&struct {
        
        Jid string `json:"jid"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

