package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CursornotelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CursornotelistingDud struct { 
    


    


    


    

}

// Cursornotelisting
type Cursornotelisting struct { 
    // Entities
    Entities []Note `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Cursornotelisting) String() string {
     o.Entities = []Note{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cursornotelisting) MarshalJSON() ([]byte, error) {
    type Alias Cursornotelisting

    if CursornotelistingMarshalled {
        return []byte("{}"), nil
    }
    CursornotelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Note `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Note{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

