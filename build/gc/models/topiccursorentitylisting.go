package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TopiccursorentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TopiccursorentitylistingDud struct { 
    


    


    


    

}

// Topiccursorentitylisting - Cursor listing of Topics.
type Topiccursorentitylisting struct { 
    // Entities
    Entities []string `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Topiccursorentitylisting) String() string {
     o.Entities = []string{""} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Topiccursorentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Topiccursorentitylisting

    if TopiccursorentitylistingMarshalled {
        return []byte("{}"), nil
    }
    TopiccursorentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []string `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

