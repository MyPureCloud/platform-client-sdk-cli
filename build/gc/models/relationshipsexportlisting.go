package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RelationshipsexportlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RelationshipsexportlistingDud struct { 
    


    


    


    

}

// Relationshipsexportlisting
type Relationshipsexportlisting struct { 
    // Entities
    Entities []Relationshipsexport `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Relationshipsexportlisting) String() string {
     o.Entities = []Relationshipsexport{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Relationshipsexportlisting) MarshalJSON() ([]byte, error) {
    type Alias Relationshipsexportlisting

    if RelationshipsexportlistingMarshalled {
        return []byte("{}"), nil
    }
    RelationshipsexportlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Relationshipsexport `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Relationshipsexport{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

