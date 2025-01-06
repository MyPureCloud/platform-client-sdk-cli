package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappconfigDud struct { 
    


    


    

}

// Whatsappconfig
type Whatsappconfig struct { 
    // WhatsAppColumns - The contact list columns specifying the WhatsApp address(es) of the contact.
    WhatsAppColumns []string `json:"whatsAppColumns"`


    // WhatsAppIntegration - The WhatsApp integration used to send message to the contact.
    WhatsAppIntegration Addressableentityref `json:"whatsAppIntegration"`


    // ContentTemplate - The content template used to formulate the WhatsApp message to send to the contact.
    ContentTemplate Domainentityref `json:"contentTemplate"`

}

// String returns a JSON representation of the model
func (o *Whatsappconfig) String() string {
     o.WhatsAppColumns = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappconfig) MarshalJSON() ([]byte, error) {
    type Alias Whatsappconfig

    if WhatsappconfigMarshalled {
        return []byte("{}"), nil
    }
    WhatsappconfigMarshalled = true

    return json.Marshal(&struct {
        
        WhatsAppColumns []string `json:"whatsAppColumns"`
        
        WhatsAppIntegration Addressableentityref `json:"whatsAppIntegration"`
        
        ContentTemplate Domainentityref `json:"contentTemplate"`
        *Alias
    }{

        
        WhatsAppColumns: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

