package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SignatureMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SignatureDud struct { 
    


    


    


    

}

// Signature
type Signature struct { 
    // Enabled - A toggle to enable the signature on email send.
    Enabled bool `json:"enabled"`


    // CannedResponseId - The identifier referring to an email signature canned response.
    CannedResponseId string `json:"cannedResponseId"`


    // AlwaysIncluded - A toggle that defines if a signature is always included or only set on the first email in an email chain.
    AlwaysIncluded bool `json:"alwaysIncluded"`


    // InclusionType - The configuration to indicate when the signature of a conversation has to be included
    InclusionType string `json:"inclusionType"`

}

// String returns a JSON representation of the model
func (o *Signature) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Signature) MarshalJSON() ([]byte, error) {
    type Alias Signature

    if SignatureMarshalled {
        return []byte("{}"), nil
    }
    SignatureMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        CannedResponseId string `json:"cannedResponseId"`
        
        AlwaysIncluded bool `json:"alwaysIncluded"`
        
        InclusionType string `json:"inclusionType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

