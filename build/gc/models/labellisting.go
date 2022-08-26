package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LabellistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LabellistingDud struct { 
    


    


    


    

}

// Labellisting
type Labellisting struct { 
    // Entities
    Entities []Labelresponse `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Labellisting) String() string {
     o.Entities = []Labelresponse{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Labellisting) MarshalJSON() ([]byte, error) {
    type Alias Labellisting

    if LabellistingMarshalled {
        return []byte("{}"), nil
    }
    LabellistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Labelresponse `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Labelresponse{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

