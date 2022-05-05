package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingrecipientMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingrecipientDud struct { 
    Nickname string `json:"nickname"`


    


    


    Image string `json:"image"`


    FirstName string `json:"firstName"`


    LastName string `json:"lastName"`


    Email string `json:"email"`


    

}

// Messagingrecipient - Information about the recipient the message is sent to or received from.
type Messagingrecipient struct { 
    


    // Id - The recipient ID specific to the provider.
    Id string `json:"id"`


    // IdType - The recipient ID type. This is used to indicate the format used for the ID.
    IdType string `json:"idType"`


    


    


    


    


    // AdditionalIds - List of recipient additional identifiers
    AdditionalIds []Recipientadditionalidentifier `json:"additionalIds"`

}

// String returns a JSON representation of the model
func (o *Messagingrecipient) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.AdditionalIds = []Recipientadditionalidentifier{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingrecipient) MarshalJSON() ([]byte, error) {
    type Alias Messagingrecipient

    if MessagingrecipientMarshalled {
        return []byte("{}"), nil
    }
    MessagingrecipientMarshalled = true

    return json.Marshal(&struct { 
        
        
        Id string `json:"id"`
        
        IdType string `json:"idType"`
        
        
        
        
        
        
        
        
        
        AdditionalIds []Recipientadditionalidentifier `json:"additionalIds"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        AdditionalIds: []Recipientadditionalidentifier{{}},
        

        
        Alias: (*Alias)(u),
    })
}

