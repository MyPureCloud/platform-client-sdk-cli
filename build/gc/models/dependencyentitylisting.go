package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DependencyentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DependencyentitylistingDud struct { 
    


    


    


    

}

// Dependencyentitylisting
type Dependencyentitylisting struct { 
    // Entities
    Entities []Dependencyentity `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Dependencyentitylisting) String() string {
     o.Entities = []Dependencyentity{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dependencyentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Dependencyentitylisting

    if DependencyentitylistingMarshalled {
        return []byte("{}"), nil
    }
    DependencyentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Dependencyentity `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Dependencyentity{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

