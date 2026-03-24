package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SteplistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SteplistingDud struct { 
    


    


    


    

}

// Steplisting
type Steplisting struct { 
    // Entities
    Entities []Step `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Steplisting) String() string {
     o.Entities = []Step{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Steplisting) MarshalJSON() ([]byte, error) {
    type Alias Steplisting

    if SteplistingMarshalled {
        return []byte("{}"), nil
    }
    SteplistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Step `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Step{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

