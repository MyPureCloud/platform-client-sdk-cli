package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappintegrationupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappintegrationupdaterequestDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Whatsappintegrationupdaterequest
type Whatsappintegrationupdaterequest struct { 
    


    // Name - WhatsApp Integration name
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting - Defines the message settings to be applied for this integration
    MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`


    

}

// String returns a JSON representation of the model
func (o *Whatsappintegrationupdaterequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappintegrationupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Whatsappintegrationupdaterequest

    if WhatsappintegrationupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    WhatsappintegrationupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

