package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappdefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappdefinitionDud struct { 
    


    


    


    


    


    

}

// Whatsappdefinition - A WhatsApp messaging template definition as defined in the WhatsApp Business Manager
type Whatsappdefinition struct { 
    // Name - The messaging template name.
    Name string `json:"name"`


    // Namespace - The messaging template namespace.
    Namespace string `json:"namespace"`


    // Language - The messaging template language configured for this template. This is a WhatsApp specific value. For example, 'en_US'
    Language string `json:"language"`


    // Buttons - List of buttons to be included in the WhatsApp messages channel
    Buttons []Button `json:"buttons"`


    // MessageFooter - Footer for the message in the WhatsApp messages channel
    MessageFooter Messagefooter `json:"messageFooter"`


    // Header - Header for the message in the WhatsApp messages channel
    Header Messageheader `json:"header"`

}

// String returns a JSON representation of the model
func (o *Whatsappdefinition) String() string {
    
    
    
     o.Buttons = []Button{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappdefinition) MarshalJSON() ([]byte, error) {
    type Alias Whatsappdefinition

    if WhatsappdefinitionMarshalled {
        return []byte("{}"), nil
    }
    WhatsappdefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Namespace string `json:"namespace"`
        
        Language string `json:"language"`
        
        Buttons []Button `json:"buttons"`
        
        MessageFooter Messagefooter `json:"messageFooter"`
        
        Header Messageheader `json:"header"`
        *Alias
    }{

        


        


        


        
        Buttons: []Button{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

