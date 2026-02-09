package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistantcopilotvariationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistantcopilotvariationlistingDud struct { 
    


    


    


    

}

// Assistantcopilotvariationlisting
type Assistantcopilotvariationlisting struct { 
    // Entities
    Entities []Assistantcopilotvariation `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Assistantcopilotvariationlisting) String() string {
     o.Entities = []Assistantcopilotvariation{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistantcopilotvariationlisting) MarshalJSON() ([]byte, error) {
    type Alias Assistantcopilotvariationlisting

    if AssistantcopilotvariationlistingMarshalled {
        return []byte("{}"), nil
    }
    AssistantcopilotvariationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Assistantcopilotvariation `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Assistantcopilotvariation{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

