package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappintegrationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappintegrationrequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Whatsappintegrationrequest
type Whatsappintegrationrequest struct { 
    


    // Name - The name of the WhatsApp Integration
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting
    MessagingSetting Messagingsettingreference `json:"messagingSetting"`


    // PhoneNumber - The phone number associated to the whatsApp integration
    PhoneNumber string `json:"phoneNumber"`


    // WabaCertificate - The waba(WhatsApp Business Manager) certificate associated to the WhatsApp integration phone number
    WabaCertificate string `json:"wabaCertificate"`


    

}

// String returns a JSON representation of the model
func (o *Whatsappintegrationrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappintegrationrequest) MarshalJSON() ([]byte, error) {
    type Alias Whatsappintegrationrequest

    if WhatsappintegrationrequestMarshalled {
        return []byte("{}"), nil
    }
    WhatsappintegrationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingreference `json:"messagingSetting"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        WabaCertificate string `json:"wabaCertificate"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

