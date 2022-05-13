package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmessagingtorecipientMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmessagingtorecipientDud struct { 
    Nickname string `json:"nickname"`


    


    IdType string `json:"idType"`


    Image string `json:"image"`


    FirstName string `json:"firstName"`


    LastName string `json:"lastName"`


    Email string `json:"email"`


    AdditionalIds []Conversationrecipientadditionalidentifier `json:"additionalIds"`

}

// Conversationmessagingtorecipient - Information about the recipient the message is sent to.
type Conversationmessagingtorecipient struct { 
    


    // Id - The recipient ID specific to the provider.
    Id string `json:"id"`


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Conversationmessagingtorecipient) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmessagingtorecipient) MarshalJSON() ([]byte, error) {
    type Alias Conversationmessagingtorecipient

    if ConversationmessagingtorecipientMarshalled {
        return []byte("{}"), nil
    }
    ConversationmessagingtorecipientMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

