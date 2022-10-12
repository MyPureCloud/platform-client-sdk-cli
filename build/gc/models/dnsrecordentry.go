package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DnsrecordentryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DnsrecordentryDud struct { 
    


    


    

}

// Dnsrecordentry
type Dnsrecordentry struct { 
    // Host - the hostname of the DNS entry
    Host string `json:"host"`


    // RecordContents - the payload of the DNS entry
    RecordContents string `json:"recordContents"`


    // VerificationStatus - the current status of the related verification process
    VerificationStatus string `json:"verificationStatus"`

}

// String returns a JSON representation of the model
func (o *Dnsrecordentry) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dnsrecordentry) MarshalJSON() ([]byte, error) {
    type Alias Dnsrecordentry

    if DnsrecordentryMarshalled {
        return []byte("{}"), nil
    }
    DnsrecordentryMarshalled = true

    return json.Marshal(&struct {
        
        Host string `json:"host"`
        
        RecordContents string `json:"recordContents"`
        
        VerificationStatus string `json:"verificationStatus"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

