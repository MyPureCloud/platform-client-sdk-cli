package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeammemberentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeammemberentitylistingDud struct { 
    


    


    


    

}

// Teammemberentitylisting
type Teammemberentitylisting struct { 
    // Entities
    Entities []Userreferencewithname `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Teammemberentitylisting) String() string {
     o.Entities = []Userreferencewithname{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teammemberentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Teammemberentitylisting

    if TeammemberentitylistingMarshalled {
        return []byte("{}"), nil
    }
    TeammemberentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Userreferencewithname `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Userreferencewithname{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

