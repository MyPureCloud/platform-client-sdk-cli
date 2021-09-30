package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgebaselistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgebaselistingDud struct { 
    


    


    


    

}

// Knowledgebaselisting
type Knowledgebaselisting struct { 
    // Entities
    Entities []Knowledgebase `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Knowledgebaselisting) String() string {
    
    
     o.Entities = []Knowledgebase{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgebaselisting) MarshalJSON() ([]byte, error) {
    type Alias Knowledgebaselisting

    if KnowledgebaselistingMarshalled {
        return []byte("{}"), nil
    }
    KnowledgebaselistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Knowledgebase `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        *Alias
    }{
        

        
        Entities: []Knowledgebase{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

