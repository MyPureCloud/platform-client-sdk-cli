package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpeninboundmessagecontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpeninboundmessagecontentDud struct { 
    

}

// Openinboundmessagecontent - Message content element.
type Openinboundmessagecontent struct { 
    // Attachment - Attachment content.
    Attachment Opencontentattachment `json:"attachment"`

}

// String returns a JSON representation of the model
func (o *Openinboundmessagecontent) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openinboundmessagecontent) MarshalJSON() ([]byte, error) {
    type Alias Openinboundmessagecontent

    if OpeninboundmessagecontentMarshalled {
        return []byte("{}"), nil
    }
    OpeninboundmessagecontentMarshalled = true

    return json.Marshal(&struct {
        
        Attachment Opencontentattachment `json:"attachment"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

