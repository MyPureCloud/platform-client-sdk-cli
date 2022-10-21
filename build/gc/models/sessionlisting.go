package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SessionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SessionlistingDud struct { 
    


    


    


    

}

// Sessionlisting
type Sessionlisting struct { 
    // Entities
    Entities []Session `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Sessionlisting) String() string {
     o.Entities = []Session{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sessionlisting) MarshalJSON() ([]byte, error) {
    type Alias Sessionlisting

    if SessionlistingMarshalled {
        return []byte("{}"), nil
    }
    SessionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Session `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Session{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

