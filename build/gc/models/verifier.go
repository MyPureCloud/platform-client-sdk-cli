package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VerifierMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VerifierDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Verifier
type Verifier struct { 
    


    // Name
    Name string `json:"name"`


    // Enabled - Indicates whether this verifier is enabled.
    Enabled bool `json:"enabled"`


    // VarDefault - Indicates whether this is the default verifier.
    VarDefault bool `json:"default"`


    

}

// String returns a JSON representation of the model
func (o *Verifier) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Verifier) MarshalJSON() ([]byte, error) {
    type Alias Verifier

    if VerifierMarshalled {
        return []byte("{}"), nil
    }
    VerifierMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Enabled bool `json:"enabled"`
        
        VarDefault bool `json:"default"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

