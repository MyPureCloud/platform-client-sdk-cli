package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsconfigDud struct { 
    


    


    


    

}

// Smsconfig
type Smsconfig struct { 
    // MessageColumn - The Contact List column specifying the message to send to the contact.
    MessageColumn string `json:"messageColumn"`


    // PhoneColumn - The Contact List column specifying the phone number to send a message to.
    PhoneColumn string `json:"phoneColumn"`


    // SenderSmsPhoneNumber - A reference to the SMS Phone Number that will be used as the sender of a message.
    SenderSmsPhoneNumber Smsphonenumberref `json:"senderSmsPhoneNumber"`


    // ContentTemplate - The content template used to formulate the message to send to the contact.
    ContentTemplate Domainentityref `json:"contentTemplate"`

}

// String returns a JSON representation of the model
func (o *Smsconfig) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsconfig) MarshalJSON() ([]byte, error) {
    type Alias Smsconfig

    if SmsconfigMarshalled {
        return []byte("{}"), nil
    }
    SmsconfigMarshalled = true

    return json.Marshal(&struct {
        
        MessageColumn string `json:"messageColumn"`
        
        PhoneColumn string `json:"phoneColumn"`
        
        SenderSmsPhoneNumber Smsphonenumberref `json:"senderSmsPhoneNumber"`
        
        ContentTemplate Domainentityref `json:"contentTemplate"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

