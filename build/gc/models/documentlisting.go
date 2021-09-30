package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentlistingDud struct { 
    


    


    


    

}

// Documentlisting
type Documentlisting struct { 
    // Entities
    Entities []Knowledgedocument `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Documentlisting) String() string {
    
    
     o.Entities = []Knowledgedocument{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentlisting) MarshalJSON() ([]byte, error) {
    type Alias Documentlisting

    if DocumentlistingMarshalled {
        return []byte("{}"), nil
    }
    DocumentlistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Knowledgedocument `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        *Alias
    }{
        

        
        Entities: []Knowledgedocument{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

