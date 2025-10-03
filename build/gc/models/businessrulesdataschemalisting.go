package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessrulesdataschemalistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessrulesdataschemalistingDud struct { 
    


    


    


    

}

// Businessrulesdataschemalisting
type Businessrulesdataschemalisting struct { 
    // Entities
    Entities []Businessrulesdataschema `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Businessrulesdataschemalisting) String() string {
     o.Entities = []Businessrulesdataschema{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessrulesdataschemalisting) MarshalJSON() ([]byte, error) {
    type Alias Businessrulesdataschemalisting

    if BusinessrulesdataschemalistingMarshalled {
        return []byte("{}"), nil
    }
    BusinessrulesdataschemalistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Businessrulesdataschema `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Businessrulesdataschema{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

