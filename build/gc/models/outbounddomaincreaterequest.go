package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutbounddomaincreaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutbounddomaincreaterequestDud struct { 
    


    


    


    

}

// Outbounddomaincreaterequest
type Outbounddomaincreaterequest struct { 
    // Id - Unique Id of the domain such as: example.com
    Id string `json:"id"`


    // SenderType - Sender Type
    SenderType string `json:"senderType"`


    // EmailSetting - The email settings to associate with this domain.
    EmailSetting Emailsettingreference `json:"emailSetting"`


    // Name - The domain such as: example.com
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Outbounddomaincreaterequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outbounddomaincreaterequest) MarshalJSON() ([]byte, error) {
    type Alias Outbounddomaincreaterequest

    if OutbounddomaincreaterequestMarshalled {
        return []byte("{}"), nil
    }
    OutbounddomaincreaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SenderType string `json:"senderType"`
        
        EmailSetting Emailsettingreference `json:"emailSetting"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

