package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutgoingmessageresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutgoingmessageresponseDud struct { 
    

}

// Outgoingmessageresponse - Outgoing Message response
type Outgoingmessageresponse struct { 
    // MessageId - The message ID.
    MessageId string `json:"messageId"`

}

// String returns a JSON representation of the model
func (o *Outgoingmessageresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outgoingmessageresponse) MarshalJSON() ([]byte, error) {
    type Alias Outgoingmessageresponse

    if OutgoingmessageresponseMarshalled {
        return []byte("{}"), nil
    }
    OutgoingmessageresponseMarshalled = true

    return json.Marshal(&struct {
        
        MessageId string `json:"messageId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

