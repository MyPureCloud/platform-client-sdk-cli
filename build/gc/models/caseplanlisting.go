package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseplanlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseplanlistingDud struct { 
    


    


    


    

}

// Caseplanlisting
type Caseplanlisting struct { 
    // Entities
    Entities []Caseplan `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Caseplanlisting) String() string {
     o.Entities = []Caseplan{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseplanlisting) MarshalJSON() ([]byte, error) {
    type Alias Caseplanlisting

    if CaseplanlistingMarshalled {
        return []byte("{}"), nil
    }
    CaseplanlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Caseplan `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Caseplan{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

