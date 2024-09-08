package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistantlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistantlistingDud struct { 
    


    


    


    

}

// Assistantlisting
type Assistantlisting struct { 
    // Entities
    Entities []Assistant `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Assistantlisting) String() string {
     o.Entities = []Assistant{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistantlisting) MarshalJSON() ([]byte, error) {
    type Alias Assistantlisting

    if AssistantlistingMarshalled {
        return []byte("{}"), nil
    }
    AssistantlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Assistant `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Assistant{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

