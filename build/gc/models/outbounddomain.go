package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutbounddomainMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutbounddomainDud struct { 
    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Outbounddomain
type Outbounddomain struct { 
    // Id - Unique Id of the domain such as: example.com
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // CnameVerificationResult - CNAME registration Status
    CnameVerificationResult Verificationresult `json:"cnameVerificationResult"`


    // DkimVerificationResult - DKIM registration Status
    DkimVerificationResult Verificationresult `json:"dkimVerificationResult"`


    // SenderType - Sender Type
    SenderType string `json:"senderType"`


    

}

// String returns a JSON representation of the model
func (o *Outbounddomain) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outbounddomain) MarshalJSON() ([]byte, error) {
    type Alias Outbounddomain

    if OutbounddomainMarshalled {
        return []byte("{}"), nil
    }
    OutbounddomainMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        CnameVerificationResult Verificationresult `json:"cnameVerificationResult"`
        
        DkimVerificationResult Verificationresult `json:"dkimVerificationResult"`
        
        SenderType string `json:"senderType"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

