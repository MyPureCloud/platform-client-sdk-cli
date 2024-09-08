package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistantqueuelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistantqueuelistingDud struct { 
    


    


    


    

}

// Assistantqueuelisting
type Assistantqueuelisting struct { 
    // Entities
    Entities []Assistantqueue `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Assistantqueuelisting) String() string {
     o.Entities = []Assistantqueue{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistantqueuelisting) MarshalJSON() ([]byte, error) {
    type Alias Assistantqueuelisting

    if AssistantqueuelistingMarshalled {
        return []byte("{}"), nil
    }
    AssistantqueuelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Assistantqueue `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Assistantqueue{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

