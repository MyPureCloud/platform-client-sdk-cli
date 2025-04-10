package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthzdivisioncursorlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthzdivisioncursorlistingDud struct { 
    


    


    


    

}

// Authzdivisioncursorlisting
type Authzdivisioncursorlisting struct { 
    // Entities
    Entities []Authzdivision `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Authzdivisioncursorlisting) String() string {
     o.Entities = []Authzdivision{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authzdivisioncursorlisting) MarshalJSON() ([]byte, error) {
    type Alias Authzdivisioncursorlisting

    if AuthzdivisioncursorlistingMarshalled {
        return []byte("{}"), nil
    }
    AuthzdivisioncursorlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Authzdivision `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Authzdivision{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

