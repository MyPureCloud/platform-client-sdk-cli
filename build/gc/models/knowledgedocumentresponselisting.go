package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentresponselistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentresponselistingDud struct { 
    


    


    


    

}

// Knowledgedocumentresponselisting
type Knowledgedocumentresponselisting struct { 
    // Entities
    Entities []Knowledgedocumentresponse `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentresponselisting) String() string {
     o.Entities = []Knowledgedocumentresponse{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentresponselisting) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentresponselisting

    if KnowledgedocumentresponselistingMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentresponselistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Knowledgedocumentresponse `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Knowledgedocumentresponse{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

