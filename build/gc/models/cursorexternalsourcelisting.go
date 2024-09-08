package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CursorexternalsourcelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CursorexternalsourcelistingDud struct { 
    


    


    


    


    

}

// Cursorexternalsourcelisting
type Cursorexternalsourcelisting struct { 
    // Entities
    Entities []Externalsource `json:"entities"`


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
func (o *Cursorexternalsourcelisting) String() string {
     o.Entities = []Externalsource{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cursorexternalsourcelisting) MarshalJSON() ([]byte, error) {
    type Alias Cursorexternalsourcelisting

    if CursorexternalsourcelistingMarshalled {
        return []byte("{}"), nil
    }
    CursorexternalsourcelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Externalsource `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        Cursors Cursors `json:"cursors"`
        *Alias
    }{

        
        Entities: []Externalsource{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

