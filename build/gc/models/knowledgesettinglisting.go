package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesettinglistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesettinglistingDud struct { 
    


    


    


    

}

// Knowledgesettinglisting
type Knowledgesettinglisting struct { 
    // Entities
    Entities []Knowledgesettingsresponse `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Knowledgesettinglisting) String() string {
     o.Entities = []Knowledgesettingsresponse{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesettinglisting) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesettinglisting

    if KnowledgesettinglistingMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesettinglistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Knowledgesettingsresponse `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Knowledgesettingsresponse{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

