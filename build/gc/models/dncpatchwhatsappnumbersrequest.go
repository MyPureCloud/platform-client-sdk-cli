package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DncpatchwhatsappnumbersrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DncpatchwhatsappnumbersrequestDud struct { 
    


    


    

}

// Dncpatchwhatsappnumbersrequest
type Dncpatchwhatsappnumbersrequest struct { 
    // Action - The action to perform
    Action string `json:"action"`


    // WhatsAppNumbers - The list of whatsApp numbers to Add to / Remove from the DNC list 
    WhatsAppNumbers []string `json:"whatsAppNumbers"`


    // ExpirationDateTime - Expiration date for DNC whatsApp numbers in yyyy-MM-ddTHH:mmZ format
    ExpirationDateTime string `json:"expirationDateTime"`

}

// String returns a JSON representation of the model
func (o *Dncpatchwhatsappnumbersrequest) String() string {
    
     o.WhatsAppNumbers = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dncpatchwhatsappnumbersrequest) MarshalJSON() ([]byte, error) {
    type Alias Dncpatchwhatsappnumbersrequest

    if DncpatchwhatsappnumbersrequestMarshalled {
        return []byte("{}"), nil
    }
    DncpatchwhatsappnumbersrequestMarshalled = true

    return json.Marshal(&struct {
        
        Action string `json:"action"`
        
        WhatsAppNumbers []string `json:"whatsAppNumbers"`
        
        ExpirationDateTime string `json:"expirationDateTime"`
        *Alias
    }{

        


        
        WhatsAppNumbers: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

