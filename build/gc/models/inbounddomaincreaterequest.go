package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InbounddomaincreaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InbounddomaincreaterequestDud struct { 
    


    


    


    


    

}

// Inbounddomaincreaterequest
type Inbounddomaincreaterequest struct { 
    // Id - Unique Id of the domain such as: example.com
    Id string `json:"id"`


    // SubDomain - Indicates if this a PureCloud sub-domain. If true, then the appropriate DNS records are created for sending/receiving email.
    SubDomain bool `json:"subDomain"`


    // MailFromSettings - The DNS settings if the inbound domain is using a custom Mail From. These settings can only be used on InboundDomains where subDomain is false.
    MailFromSettings Mailfromresult `json:"mailFromSettings"`


    // CustomSMTPServer - The custom SMTP server integration to use when sending outbound emails from this domain.
    CustomSMTPServer Domainentityref `json:"customSMTPServer"`


    // EmailSetting - The email settings to associate with this domain.
    EmailSetting Emailsettingreference `json:"emailSetting"`

}

// String returns a JSON representation of the model
func (o *Inbounddomaincreaterequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Inbounddomaincreaterequest) MarshalJSON() ([]byte, error) {
    type Alias Inbounddomaincreaterequest

    if InbounddomaincreaterequestMarshalled {
        return []byte("{}"), nil
    }
    InbounddomaincreaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SubDomain bool `json:"subDomain"`
        
        MailFromSettings Mailfromresult `json:"mailFromSettings"`
        
        CustomSMTPServer Domainentityref `json:"customSMTPServer"`
        
        EmailSetting Emailsettingreference `json:"emailSetting"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

