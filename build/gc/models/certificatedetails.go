package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CertificatedetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CertificatedetailsDud struct { 
    


    


    


    


    


    


    

}

// Certificatedetails - Represents the details of a parsed certificate.
type Certificatedetails struct { 
    // Issuer - Information about the issuer of the certificate.  The value of this property is a comma separated key=value format.  Each key is one of the attribute names supported by X.500.
    Issuer string `json:"issuer"`


    // Subject - Information about the subject of the certificate.  The value of this property is a comma separated key=value format.  Each key is one of the attribute names supported by X.500.
    Subject string `json:"subject"`


    // ExpirationDate - The expiration date of the certificate. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ExpirationDate time.Time `json:"expirationDate"`


    // IssueDate - The issue date of the certificate. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    IssueDate time.Time `json:"issueDate"`


    // Expired - True if the certificate is expired, false otherwise.
    Expired bool `json:"expired"`


    // SignatureValid
    SignatureValid bool `json:"signatureValid"`


    // Valid
    Valid bool `json:"valid"`

}

// String returns a JSON representation of the model
func (o *Certificatedetails) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Certificatedetails) MarshalJSON() ([]byte, error) {
    type Alias Certificatedetails

    if CertificatedetailsMarshalled {
        return []byte("{}"), nil
    }
    CertificatedetailsMarshalled = true

    return json.Marshal(&struct {
        
        Issuer string `json:"issuer"`
        
        Subject string `json:"subject"`
        
        ExpirationDate time.Time `json:"expirationDate"`
        
        IssueDate time.Time `json:"issueDate"`
        
        Expired bool `json:"expired"`
        
        SignatureValid bool `json:"signatureValid"`
        
        Valid bool `json:"valid"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

