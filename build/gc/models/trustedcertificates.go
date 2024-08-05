package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrustedcertificatesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrustedcertificatesDud struct { 
    


    

}

// Trustedcertificates - Information about trusted certificates
type Trustedcertificates struct { 
    // Entities - The list of trusted certificates
    Entities []Trustedcertificateinfo `json:"entities"`


    // Total - The total number of trusted certificates
    Total int `json:"total"`

}

// String returns a JSON representation of the model
func (o *Trustedcertificates) String() string {
     o.Entities = []Trustedcertificateinfo{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trustedcertificates) MarshalJSON() ([]byte, error) {
    type Alias Trustedcertificates

    if TrustedcertificatesMarshalled {
        return []byte("{}"), nil
    }
    TrustedcertificatesMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Trustedcertificateinfo `json:"entities"`
        
        Total int `json:"total"`
        *Alias
    }{

        
        Entities: []Trustedcertificateinfo{{}},
        


        

        Alias: (*Alias)(u),
    })
}

