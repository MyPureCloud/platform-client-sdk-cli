package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StageplanlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StageplanlistingDud struct { 
    


    


    


    

}

// Stageplanlisting
type Stageplanlisting struct { 
    // Entities
    Entities []Stageplan `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Stageplanlisting) String() string {
     o.Entities = []Stageplan{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stageplanlisting) MarshalJSON() ([]byte, error) {
    type Alias Stageplanlisting

    if StageplanlistingMarshalled {
        return []byte("{}"), nil
    }
    StageplanlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Stageplan `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Stageplan{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

