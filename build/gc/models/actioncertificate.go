package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActioncertificateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActioncertificateDud struct { 
    


    


    


    

}

// Actioncertificate - Details for an mTLS certificate
type Actioncertificate struct { 
    // SigningAuthority - The Signing Authority for the certificate
    SigningAuthority string `json:"signingAuthority"`


    // Certificate - The certificate string
    Certificate string `json:"certificate"`


    // Status - The certificate status
    Status string `json:"status"`


    // VarType - The certificate type
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Actioncertificate) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actioncertificate) MarshalJSON() ([]byte, error) {
    type Alias Actioncertificate

    if ActioncertificateMarshalled {
        return []byte("{}"), nil
    }
    ActioncertificateMarshalled = true

    return json.Marshal(&struct {
        
        SigningAuthority string `json:"signingAuthority"`
        
        Certificate string `json:"certificate"`
        
        Status string `json:"status"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

