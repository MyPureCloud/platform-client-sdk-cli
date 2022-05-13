package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenmessagingtorecipientMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenmessagingtorecipientDud struct { 
    Nickname string `json:"nickname"`


    


    IdType string `json:"idType"`


    FirstName string `json:"firstName"`


    LastName string `json:"lastName"`


    Image string `json:"image"`


    

}

// Openmessagingtorecipient - Information about the recipient the message is sent to.
type Openmessagingtorecipient struct { 
    


    // Id - The recipient ID specific to the provider.
    Id string `json:"id"`


    


    


    


    


    // Email - E-mail address of the recipient.
    Email string `json:"email"`

}

// String returns a JSON representation of the model
func (o *Openmessagingtorecipient) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openmessagingtorecipient) MarshalJSON() ([]byte, error) {
    type Alias Openmessagingtorecipient

    if OpenmessagingtorecipientMarshalled {
        return []byte("{}"), nil
    }
    OpenmessagingtorecipientMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Email string `json:"email"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

