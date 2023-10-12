package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatsendmessageresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatsendmessageresponseDud struct { 
    

}

// Chatsendmessageresponse
type Chatsendmessageresponse struct { 
    // Id - The id of the created message
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Chatsendmessageresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatsendmessageresponse) MarshalJSON() ([]byte, error) {
    type Alias Chatsendmessageresponse

    if ChatsendmessageresponseMarshalled {
        return []byte("{}"), nil
    }
    ChatsendmessageresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

