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


    Name string `json:"name"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Whatsappintegrationupdaterequest
type Whatsappintegrationupdaterequest struct { 
    


    


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting
    MessagingSetting Messagingsettingreference `json:"messagingSetting"`


    // Action - The action used to activate and then confirm a WhatsApp Integration.
    Action string `json:"action"`


    // AuthenticationMethod - The authentication method used to confirm a WhatsApp Integration activation. If action is set to Activate, then authenticationMethod is a required field. 
    AuthenticationMethod string `json:"authenticationMethod"`


    // ConfirmationCode - The confirmation code sent by Whatsapp to you during the activation step. If action is set to Confirm, then confirmationCode is a required field.
    ConfirmationCode string `json:"confirmationCode"`


    

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
        
        
        
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingreference `json:"messagingSetting"`
        
        Action string `json:"action"`
        
        AuthenticationMethod string `json:"authenticationMethod"`
        
        ConfirmationCode string `json:"confirmationCode"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

