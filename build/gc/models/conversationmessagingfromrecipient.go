package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmessagingfromrecipientMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmessagingfromrecipientDud struct { 
    Nickname string `json:"nickname"`


    


    IdType string `json:"idType"`


    


    


    


    Email string `json:"email"`


    AdditionalIds []Conversationrecipientadditionalidentifier `json:"additionalIds"`

}

// Conversationmessagingfromrecipient - Information about the recipient the message is received from.
type Conversationmessagingfromrecipient struct { 
    


    // Id - The recipient ID specific to the provider.
    Id string `json:"id"`


    


    // Image - URL of an image that represents the recipient.
    Image string `json:"image"`


    // FirstName - First name of the recipient.
    FirstName string `json:"firstName"`


    // LastName - Last name of the recipient.
    LastName string `json:"lastName"`


    


    

}

// String returns a JSON representation of the model
func (o *Conversationmessagingfromrecipient) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmessagingfromrecipient) MarshalJSON() ([]byte, error) {
    type Alias Conversationmessagingfromrecipient

    if ConversationmessagingfromrecipientMarshalled {
        return []byte("{}"), nil
    }
    ConversationmessagingfromrecipientMarshalled = true

    return json.Marshal(&struct { 
        
        
        Id string `json:"id"`
        
        
        
        Image string `json:"image"`
        
        FirstName string `json:"firstName"`
        
        LastName string `json:"lastName"`
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

