package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InbounddomainMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InbounddomainDud struct { 
    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Inbounddomain
type Inbounddomain struct { 
    // Id - Unique Id of the domain such as: example.com
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // MxRecordStatus - Mx Record Status
    MxRecordStatus string `json:"mxRecordStatus"`


    // SubDomain - Indicates if this a PureCloud sub-domain.  If true, then the appropriate DNS records are created for sending/receiving email.
    SubDomain bool `json:"subDomain"`


    // MailFromSettings - The DNS settings if the inbound domain is using a custom Mail From. These settings can only be used on InboundDomains where subDomain is false.
    MailFromSettings Mailfromresult `json:"mailFromSettings"`


    // CustomSMTPServer - The custom SMTP server integration to use when sending outbound emails from this domain.
    CustomSMTPServer Domainentityref `json:"customSMTPServer"`


    

}

// String returns a JSON representation of the model
func (o *Inbounddomain) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Inbounddomain) MarshalJSON() ([]byte, error) {
    type Alias Inbounddomain

    if InbounddomainMarshalled {
        return []byte("{}"), nil
    }
    InbounddomainMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        MxRecordStatus string `json:"mxRecordStatus"`
        
        SubDomain bool `json:"subDomain"`
        
        MailFromSettings Mailfromresult `json:"mailFromSettings"`
        
        CustomSMTPServer Domainentityref `json:"customSMTPServer"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

