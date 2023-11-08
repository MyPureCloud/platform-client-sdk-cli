package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappembeddedsignupintegrationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappembeddedsignupintegrationrequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Whatsappembeddedsignupintegrationrequest
type Whatsappembeddedsignupintegrationrequest struct { 
    


    // Name - The name of the WhatsApp Integration
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting - Defines the message settings to be applied for this integration
    MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`


    // EmbeddedSignupAccessToken - The access token returned from the embedded signup flow
    EmbeddedSignupAccessToken string `json:"embeddedSignupAccessToken"`


    

}

// String returns a JSON representation of the model
func (o *Whatsappembeddedsignupintegrationrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappembeddedsignupintegrationrequest) MarshalJSON() ([]byte, error) {
    type Alias Whatsappembeddedsignupintegrationrequest

    if WhatsappembeddedsignupintegrationrequestMarshalled {
        return []byte("{}"), nil
    }
    WhatsappembeddedsignupintegrationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`
        
        EmbeddedSignupAccessToken string `json:"embeddedSignupAccessToken"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

