package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InbounddomainpatchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InbounddomainpatchrequestDud struct { 
    


    


    


    

}

// Inbounddomainpatchrequest
type Inbounddomainpatchrequest struct { 
    // MailFromSettings - The DNS settings if the inbound domain is using a custom Mail From. These settings can only be used on InboundDomains where subDomain is false.
    MailFromSettings Mailfromresult `json:"mailFromSettings"`


    // CustomSMTPServer - The custom SMTP server integration to use when sending outbound emails from this domain.
    CustomSMTPServer Domainentityref `json:"customSMTPServer"`


    // ImapSettings - The IMAP server integration and settings to use for processing inbound emails.
    ImapSettings Imapsettings `json:"imapSettings"`


    // EmailSetting - The email settings to associate with this domain.
    EmailSetting Emailsettingreference `json:"emailSetting"`

}

// String returns a JSON representation of the model
func (o *Inbounddomainpatchrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Inbounddomainpatchrequest) MarshalJSON() ([]byte, error) {
    type Alias Inbounddomainpatchrequest

    if InbounddomainpatchrequestMarshalled {
        return []byte("{}"), nil
    }
    InbounddomainpatchrequestMarshalled = true

    return json.Marshal(&struct {
        
        MailFromSettings Mailfromresult `json:"mailFromSettings"`
        
        CustomSMTPServer Domainentityref `json:"customSMTPServer"`
        
        ImapSettings Imapsettings `json:"imapSettings"`
        
        EmailSetting Emailsettingreference `json:"emailSetting"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

