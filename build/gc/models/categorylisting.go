package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CategorylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CategorylistingDud struct { 
    


    


    


    

}

// Categorylisting
type Categorylisting struct { 
    // Entities
    Entities []Knowledgecategory `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Categorylisting) String() string {
     o.Entities = []Knowledgecategory{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Categorylisting) MarshalJSON() ([]byte, error) {
    type Alias Categorylisting

    if CategorylistingMarshalled {
        return []byte("{}"), nil
    }
    CategorylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Knowledgecategory `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Knowledgecategory{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

