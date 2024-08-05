package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrustedcertificateinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrustedcertificateinfoDud struct { 
    


    


    

}

// Trustedcertificateinfo
type Trustedcertificateinfo struct { 
    // Description - The description of the certificate
    Description string `json:"description"`


    // SerialNumber - The serial number of the certificate
    SerialNumber string `json:"serialNumber"`


    // Signature - The signature of the certificate
    Signature string `json:"signature"`

}

// String returns a JSON representation of the model
func (o *Trustedcertificateinfo) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trustedcertificateinfo) MarshalJSON() ([]byte, error) {
    type Alias Trustedcertificateinfo

    if TrustedcertificateinfoMarshalled {
        return []byte("{}"), nil
    }
    TrustedcertificateinfoMarshalled = true

    return json.Marshal(&struct {
        
        Description string `json:"description"`
        
        SerialNumber string `json:"serialNumber"`
        
        Signature string `json:"signature"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

