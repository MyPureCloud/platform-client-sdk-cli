package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CertificateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CertificateDud struct { 
    

}

// Certificate - Represents a certificate to parse.
type Certificate struct { 
    // Certificate - The certificate to parse.
    Certificate string `json:"certificate"`

}

// String returns a JSON representation of the model
func (o *Certificate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Certificate) MarshalJSON() ([]byte, error) {
    type Alias Certificate

    if CertificateMarshalled {
        return []byte("{}"), nil
    }
    CertificateMarshalled = true

    return json.Marshal(&struct {
        
        Certificate string `json:"certificate"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

