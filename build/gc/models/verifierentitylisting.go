package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VerifierentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VerifierentitylistingDud struct { 
    


    


    

}

// Verifierentitylisting
type Verifierentitylisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Verifier `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Verifierentitylisting) String() string {
    
     o.Entities = []Verifier{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Verifierentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Verifierentitylisting

    if VerifierentitylistingMarshalled {
        return []byte("{}"), nil
    }
    VerifierentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Verifier `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Verifier{{}},
        


        

        Alias: (*Alias)(u),
    })
}

