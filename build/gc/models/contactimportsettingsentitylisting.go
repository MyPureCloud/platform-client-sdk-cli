package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactimportsettingsentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactimportsettingsentitylistingDud struct { 
    


    


    


    


    

}

// Contactimportsettingsentitylisting
type Contactimportsettingsentitylisting struct { 
    // Entities
    Entities []Contactimportsettings `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // Cursors - The cursor that points to the next set of entities being returned.
    Cursors Cursors `json:"cursors"`

}

// String returns a JSON representation of the model
func (o *Contactimportsettingsentitylisting) String() string {
     o.Entities = []Contactimportsettings{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactimportsettingsentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Contactimportsettingsentitylisting

    if ContactimportsettingsentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ContactimportsettingsentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Contactimportsettings `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        Cursors Cursors `json:"cursors"`
        *Alias
    }{

        
        Entities: []Contactimportsettings{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

