package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactsexportlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactsexportlistingDud struct { 
    


    


    


    

}

// Contactsexportlisting
type Contactsexportlisting struct { 
    // Entities
    Entities []Contactsexport `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Contactsexportlisting) String() string {
     o.Entities = []Contactsexport{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactsexportlisting) MarshalJSON() ([]byte, error) {
    type Alias Contactsexportlisting

    if ContactsexportlistingMarshalled {
        return []byte("{}"), nil
    }
    ContactsexportlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Contactsexport `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Contactsexport{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

