package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NotesexportlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NotesexportlistingDud struct { 
    


    


    


    

}

// Notesexportlisting
type Notesexportlisting struct { 
    // Entities
    Entities []Notesexport `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Notesexportlisting) String() string {
     o.Entities = []Notesexport{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Notesexportlisting) MarshalJSON() ([]byte, error) {
    type Alias Notesexportlisting

    if NotesexportlistingMarshalled {
        return []byte("{}"), nil
    }
    NotesexportlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Notesexport `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Notesexport{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

