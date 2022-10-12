package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingrecipientMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingrecipientDud struct { 
    FirstName string `json:"firstName"`


    LastName string `json:"lastName"`


    Nickname string `json:"nickname"`


    Image string `json:"image"`


    AdditionalIds []Recipientadditionalidentifier `json:"additionalIds"`

}

// Webmessagingrecipient - Information about the recipient the message is sent to or received from.
type Webmessagingrecipient struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Webmessagingrecipient) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingrecipient) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingrecipient

    if WebmessagingrecipientMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingrecipientMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

