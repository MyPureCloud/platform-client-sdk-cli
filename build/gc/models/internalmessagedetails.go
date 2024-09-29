package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InternalmessagedetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InternalmessagedetailsDud struct { 
    


    

}

// Internalmessagedetails
type Internalmessagedetails struct { 
    // MessageId - UUID identifying the internal message media.
    MessageId string `json:"messageId"`


    // MessageTime - The time when the message was sent or received. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    MessageTime time.Time `json:"messageTime"`

}

// String returns a JSON representation of the model
func (o *Internalmessagedetails) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Internalmessagedetails) MarshalJSON() ([]byte, error) {
    type Alias Internalmessagedetails

    if InternalmessagedetailsMarshalled {
        return []byte("{}"), nil
    }
    InternalmessagedetailsMarshalled = true

    return json.Marshal(&struct {
        
        MessageId string `json:"messageId"`
        
        MessageTime time.Time `json:"messageTime"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

