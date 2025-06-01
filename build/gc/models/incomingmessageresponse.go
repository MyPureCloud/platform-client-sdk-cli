package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IncomingmessageresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IncomingmessageresponseDud struct { 
    

}

// Incomingmessageresponse - Incoming Message response
type Incomingmessageresponse struct { 
    // MessageId - The message ID.
    MessageId string `json:"messageId"`

}

// String returns a JSON representation of the model
func (o *Incomingmessageresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Incomingmessageresponse) MarshalJSON() ([]byte, error) {
    type Alias Incomingmessageresponse

    if IncomingmessageresponseMarshalled {
        return []byte("{}"), nil
    }
    IncomingmessageresponseMarshalled = true

    return json.Marshal(&struct {
        
        MessageId string `json:"messageId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

