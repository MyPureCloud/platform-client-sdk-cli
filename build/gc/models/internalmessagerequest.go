package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InternalmessagerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InternalmessagerequestDud struct { 
    

}

// Internalmessagerequest
type Internalmessagerequest struct { 
    // Text - The body of the text message.
    Text string `json:"text"`

}

// String returns a JSON representation of the model
func (o *Internalmessagerequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Internalmessagerequest) MarshalJSON() ([]byte, error) {
    type Alias Internalmessagerequest

    if InternalmessagerequestMarshalled {
        return []byte("{}"), nil
    }
    InternalmessagerequestMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

