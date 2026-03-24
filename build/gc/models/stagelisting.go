package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StagelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StagelistingDud struct { 
    


    


    


    

}

// Stagelisting
type Stagelisting struct { 
    // Entities
    Entities []Stage `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Stagelisting) String() string {
     o.Entities = []Stage{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stagelisting) MarshalJSON() ([]byte, error) {
    type Alias Stagelisting

    if StagelistingMarshalled {
        return []byte("{}"), nil
    }
    StagelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Stage `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Stage{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

