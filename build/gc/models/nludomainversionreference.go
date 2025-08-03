package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NludomainversionreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NludomainversionreferenceDud struct { 
    Id string `json:"id"`


    Domain *Nludomain `json:"domain"`


    


    SelfUri string `json:"selfUri"`

}

// Nludomainversionreference
type Nludomainversionreference struct { 
    


    


    // Intents - The intents defined for this NLU domain version.
    Intents []Intentreference `json:"intents"`


    

}

// String returns a JSON representation of the model
func (o *Nludomainversionreference) String() string {
     o.Intents = []Intentreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nludomainversionreference) MarshalJSON() ([]byte, error) {
    type Alias Nludomainversionreference

    if NludomainversionreferenceMarshalled {
        return []byte("{}"), nil
    }
    NludomainversionreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Intents []Intentreference `json:"intents"`
        *Alias
    }{

        


        


        
        Intents: []Intentreference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

