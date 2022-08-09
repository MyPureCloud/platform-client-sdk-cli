package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CursorrelationshiplistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CursorrelationshiplistingDud struct { 
    


    


    


    


    

}

// Cursorrelationshiplisting
type Cursorrelationshiplisting struct { 
    // Entities
    Entities []Relationship `json:"entities"`


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
func (o *Cursorrelationshiplisting) String() string {
     o.Entities = []Relationship{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cursorrelationshiplisting) MarshalJSON() ([]byte, error) {
    type Alias Cursorrelationshiplisting

    if CursorrelationshiplistingMarshalled {
        return []byte("{}"), nil
    }
    CursorrelationshiplistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Relationship `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        Cursors Cursors `json:"cursors"`
        *Alias
    }{

        
        Entities: []Relationship{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

