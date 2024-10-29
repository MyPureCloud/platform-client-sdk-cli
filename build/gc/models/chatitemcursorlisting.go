package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatitemcursorlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatitemcursorlistingDud struct { 
    


    


    


    

}

// Chatitemcursorlisting
type Chatitemcursorlisting struct { 
    // Entities
    Entities []Chatitem `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Chatitemcursorlisting) String() string {
     o.Entities = []Chatitem{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatitemcursorlisting) MarshalJSON() ([]byte, error) {
    type Alias Chatitemcursorlisting

    if ChatitemcursorlistingMarshalled {
        return []byte("{}"), nil
    }
    ChatitemcursorlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Chatitem `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Chatitem{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

