package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ParsedcertificateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ParsedcertificateDud struct { 
    

}

// Parsedcertificate - Represents the parsed certificate information.
type Parsedcertificate struct { 
    // CertificateDetails - The details of the certificates that were parsed correctly.
    CertificateDetails []Certificatedetails `json:"certificateDetails"`

}

// String returns a JSON representation of the model
func (o *Parsedcertificate) String() string {
    
    
     o.CertificateDetails = []Certificatedetails{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Parsedcertificate) MarshalJSON() ([]byte, error) {
    type Alias Parsedcertificate

    if ParsedcertificateMarshalled {
        return []byte("{}"), nil
    }
    ParsedcertificateMarshalled = true

    return json.Marshal(&struct { 
        CertificateDetails []Certificatedetails `json:"certificateDetails"`
        
        *Alias
    }{
        

        
        CertificateDetails: []Certificatedetails{{}},
        

        
        Alias: (*Alias)(u),
    })
}

