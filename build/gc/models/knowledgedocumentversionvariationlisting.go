package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentversionvariationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentversionvariationlistingDud struct { 
    


    


    


    

}

// Knowledgedocumentversionvariationlisting
type Knowledgedocumentversionvariationlisting struct { 
    // Entities
    Entities []Knowledgedocumentversionvariation `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentversionvariationlisting) String() string {
     o.Entities = []Knowledgedocumentversionvariation{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentversionvariationlisting) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentversionvariationlisting

    if KnowledgedocumentversionvariationlistingMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentversionvariationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Knowledgedocumentversionvariation `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Knowledgedocumentversionvariation{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

