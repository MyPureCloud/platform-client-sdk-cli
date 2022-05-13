package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DraftlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DraftlistingDud struct { 
    


    


    


    

}

// Draftlisting
type Draftlisting struct { 
    // Entities
    Entities []Draft `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Draftlisting) String() string {
     o.Entities = []Draft{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Draftlisting) MarshalJSON() ([]byte, error) {
    type Alias Draftlisting

    if DraftlistingMarshalled {
        return []byte("{}"), nil
    }
    DraftlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Draft `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Draft{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

