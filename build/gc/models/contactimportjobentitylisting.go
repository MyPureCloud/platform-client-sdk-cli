package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactimportjobentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactimportjobentitylistingDud struct { 
    


    


    


    


    

}

// Contactimportjobentitylisting
type Contactimportjobentitylisting struct { 
    // Entities
    Entities []Contactimportjobresponse `json:"entities"`


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
func (o *Contactimportjobentitylisting) String() string {
     o.Entities = []Contactimportjobresponse{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactimportjobentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Contactimportjobentitylisting

    if ContactimportjobentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ContactimportjobentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Contactimportjobresponse `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        Cursors Cursors `json:"cursors"`
        *Alias
    }{

        
        Entities: []Contactimportjobresponse{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

