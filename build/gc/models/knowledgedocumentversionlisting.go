package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentversionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentversionlistingDud struct { 
    


    


    


    

}

// Knowledgedocumentversionlisting
type Knowledgedocumentversionlisting struct { 
    // Entities
    Entities []Knowledgedocumentversion `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentversionlisting) String() string {
     o.Entities = []Knowledgedocumentversion{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentversionlisting) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentversionlisting

    if KnowledgedocumentversionlistingMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentversionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Knowledgedocumentversion `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Knowledgedocumentversion{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

