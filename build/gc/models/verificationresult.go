package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VerificationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VerificationresultDud struct { 
    


    

}

// Verificationresult
type Verificationresult struct { 
    // Status - The verification status.
    Status string `json:"status"`


    // Records - The list of DNS records that pertain that need to exist for verification.
    Records []Record `json:"records"`

}

// String returns a JSON representation of the model
func (o *Verificationresult) String() string {
    
     o.Records = []Record{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Verificationresult) MarshalJSON() ([]byte, error) {
    type Alias Verificationresult

    if VerificationresultMarshalled {
        return []byte("{}"), nil
    }
    VerificationresultMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        Records []Record `json:"records"`
        *Alias
    }{

        


        
        Records: []Record{{}},
        

        Alias: (*Alias)(u),
    })
}

