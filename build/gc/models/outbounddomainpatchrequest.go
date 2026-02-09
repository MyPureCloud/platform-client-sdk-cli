package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutbounddomainpatchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutbounddomainpatchrequestDud struct { 
    


    


    

}

// Outbounddomainpatchrequest
type Outbounddomainpatchrequest struct { 
    // CustomSMTPServer - The custom SMTP server integration to use when sending outbound emails from this domain.
    CustomSMTPServer Customsmtpserverrequest `json:"customSMTPServer"`


    // SenderType - Sender Type
    SenderType string `json:"senderType"`


    // EmailSetting - The email settings to associate with this domain.
    EmailSetting Emailsettingreference `json:"emailSetting"`

}

// String returns a JSON representation of the model
func (o *Outbounddomainpatchrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outbounddomainpatchrequest) MarshalJSON() ([]byte, error) {
    type Alias Outbounddomainpatchrequest

    if OutbounddomainpatchrequestMarshalled {
        return []byte("{}"), nil
    }
    OutbounddomainpatchrequestMarshalled = true

    return json.Marshal(&struct {
        
        CustomSMTPServer Customsmtpserverrequest `json:"customSMTPServer"`
        
        SenderType string `json:"senderType"`
        
        EmailSetting Emailsettingreference `json:"emailSetting"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

