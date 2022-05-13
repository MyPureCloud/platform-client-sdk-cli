package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FromemailaddressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FromemailaddressDud struct { 
    


    


    

}

// Fromemailaddress
type Fromemailaddress struct { 
    // Domain - The OutboundDomain used for the email address.
    Domain Domainentityref `json:"domain"`


    // FriendlyName - The friendly name of the email address.
    FriendlyName string `json:"friendlyName"`


    // LocalPart - The local part of the email address.
    LocalPart string `json:"localPart"`

}

// String returns a JSON representation of the model
func (o *Fromemailaddress) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Fromemailaddress) MarshalJSON() ([]byte, error) {
    type Alias Fromemailaddress

    if FromemailaddressMarshalled {
        return []byte("{}"), nil
    }
    FromemailaddressMarshalled = true

    return json.Marshal(&struct {
        
        Domain Domainentityref `json:"domain"`
        
        FriendlyName string `json:"friendlyName"`
        
        LocalPart string `json:"localPart"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

