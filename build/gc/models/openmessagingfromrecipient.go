package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenmessagingfromrecipientMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenmessagingfromrecipientDud struct { 
    


    


    


    


    


    Image string `json:"image"`


    

}

// Openmessagingfromrecipient - Information about the recipient the message is received from.
type Openmessagingfromrecipient struct { 
    // Nickname - Nickname or display name of the recipient.
    Nickname string `json:"nickname"`


    // Id - The recipient ID specific to the provider.
    Id string `json:"id"`


    // IdType - The recipient ID type. This is used to indicate the format used for the ID.
    IdType string `json:"idType"`


    // FirstName - First name of the recipient.
    FirstName string `json:"firstName"`


    // LastName - Last name of the recipient.
    LastName string `json:"lastName"`


    


    // Email - E-mail address of the recipient.
    Email string `json:"email"`

}

// String returns a JSON representation of the model
func (o *Openmessagingfromrecipient) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openmessagingfromrecipient) MarshalJSON() ([]byte, error) {
    type Alias Openmessagingfromrecipient

    if OpenmessagingfromrecipientMarshalled {
        return []byte("{}"), nil
    }
    OpenmessagingfromrecipientMarshalled = true

    return json.Marshal(&struct { 
        Nickname string `json:"nickname"`
        
        Id string `json:"id"`
        
        IdType string `json:"idType"`
        
        FirstName string `json:"firstName"`
        
        LastName string `json:"lastName"`
        
        
        
        Email string `json:"email"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

