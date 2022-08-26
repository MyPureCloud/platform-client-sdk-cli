package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CategoryresponselistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CategoryresponselistingDud struct { 
    


    


    


    

}

// Categoryresponselisting
type Categoryresponselisting struct { 
    // Entities
    Entities []Categoryresponse `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Categoryresponselisting) String() string {
     o.Entities = []Categoryresponse{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Categoryresponselisting) MarshalJSON() ([]byte, error) {
    type Alias Categoryresponselisting

    if CategoryresponselistingMarshalled {
        return []byte("{}"), nil
    }
    CategoryresponselistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Categoryresponse `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Categoryresponse{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

