package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailoutbounddomainresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailoutbounddomainresultDud struct { 
    


    


    


    


    


    

}

// Emailoutbounddomainresult
type Emailoutbounddomainresult struct { 
    // DnsCnameBounceRecord
    DnsCnameBounceRecord Dnsrecordentry `json:"dnsCnameBounceRecord"`


    // DnsTxtSendingRecord
    DnsTxtSendingRecord Dnsrecordentry `json:"dnsTxtSendingRecord"`


    // DomainName
    DomainName string `json:"domainName"`


    // SenderStatus
    SenderStatus string `json:"senderStatus"`


    // SenderType
    SenderType string `json:"senderType"`


    // EmailSetting - The email settings associated with this domain.
    EmailSetting Emailsetting `json:"emailSetting"`

}

// String returns a JSON representation of the model
func (o *Emailoutbounddomainresult) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailoutbounddomainresult) MarshalJSON() ([]byte, error) {
    type Alias Emailoutbounddomainresult

    if EmailoutbounddomainresultMarshalled {
        return []byte("{}"), nil
    }
    EmailoutbounddomainresultMarshalled = true

    return json.Marshal(&struct {
        
        DnsCnameBounceRecord Dnsrecordentry `json:"dnsCnameBounceRecord"`
        
        DnsTxtSendingRecord Dnsrecordentry `json:"dnsTxtSendingRecord"`
        
        DomainName string `json:"domainName"`
        
        SenderStatus string `json:"senderStatus"`
        
        SenderType string `json:"senderType"`
        
        EmailSetting Emailsetting `json:"emailSetting"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

