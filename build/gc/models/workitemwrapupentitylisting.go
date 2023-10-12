package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemwrapupentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemwrapupentitylistingDud struct { 
    


    


    


    

}

// Workitemwrapupentitylisting
type Workitemwrapupentitylisting struct { 
    // Entities
    Entities []Workitemwrapup `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Workitemwrapupentitylisting) String() string {
     o.Entities = []Workitemwrapup{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemwrapupentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Workitemwrapupentitylisting

    if WorkitemwrapupentitylistingMarshalled {
        return []byte("{}"), nil
    }
    WorkitemwrapupentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workitemwrapup `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Workitemwrapup{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

