package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SubscriberresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SubscriberresponseDud struct { 
    


    

}

// Subscriberresponse
type Subscriberresponse struct { 
    // MessageReturned - Suggested valid addresses
    MessageReturned []string `json:"messageReturned"`


    // Status - http status
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Subscriberresponse) String() string {
     o.MessageReturned = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Subscriberresponse) MarshalJSON() ([]byte, error) {
    type Alias Subscriberresponse

    if SubscriberresponseMarshalled {
        return []byte("{}"), nil
    }
    SubscriberresponseMarshalled = true

    return json.Marshal(&struct {
        
        MessageReturned []string `json:"messageReturned"`
        
        Status string `json:"status"`
        *Alias
    }{

        
        MessageReturned: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

