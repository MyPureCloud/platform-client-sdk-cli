package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingtemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingtemplateDud struct { 
    

}

// Messagingtemplate - The messaging template identifies a structured message templates supported by a messaging channel.
type Messagingtemplate struct { 
    // WhatsApp - Defines a messaging template for a WhatsApp messaging channel
    WhatsApp Whatsappdefinition `json:"whatsApp"`

}

// String returns a JSON representation of the model
func (o *Messagingtemplate) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingtemplate) MarshalJSON() ([]byte, error) {
    type Alias Messagingtemplate

    if MessagingtemplateMarshalled {
        return []byte("{}"), nil
    }
    MessagingtemplateMarshalled = true

    return json.Marshal(&struct { 
        WhatsApp Whatsappdefinition `json:"whatsApp"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

