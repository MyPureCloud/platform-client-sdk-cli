package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaselistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaselistingDud struct { 
    


    


    


    

}

// Caselisting
type Caselisting struct { 
    // Entities
    Entities []Case `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Caselisting) String() string {
     o.Entities = []Case{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caselisting) MarshalJSON() ([]byte, error) {
    type Alias Caselisting

    if CaselistingMarshalled {
        return []byte("{}"), nil
    }
    CaselistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Case `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Case{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

