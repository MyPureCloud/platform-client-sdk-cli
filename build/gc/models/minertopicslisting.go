package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MinertopicslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MinertopicslistingDud struct { 
    


    


    


    

}

// Minertopicslisting
type Minertopicslisting struct { 
    // Entities
    Entities []Minertopic `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Minertopicslisting) String() string {
     o.Entities = []Minertopic{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Minertopicslisting) MarshalJSON() ([]byte, error) {
    type Alias Minertopicslisting

    if MinertopicslistingMarshalled {
        return []byte("{}"), nil
    }
    MinertopicslistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Minertopic `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Minertopic{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

