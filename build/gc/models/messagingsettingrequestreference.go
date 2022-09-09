package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingsettingrequestreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingsettingrequestreferenceDud struct { 
    

}

// Messagingsettingrequestreference - Messaging Setting for messaging platform integrations
type Messagingsettingrequestreference struct { 
    // Id - The messaging Setting unique identifier associated with this integration
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Messagingsettingrequestreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingsettingrequestreference) MarshalJSON() ([]byte, error) {
    type Alias Messagingsettingrequestreference

    if MessagingsettingrequestreferenceMarshalled {
        return []byte("{}"), nil
    }
    MessagingsettingrequestreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

