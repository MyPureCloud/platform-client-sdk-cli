package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CursorsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CursorsDud struct { 
    


    

}

// Cursors
type Cursors struct { 
    // Before
    Before string `json:"before"`


    // After
    After string `json:"after"`

}

// String returns a JSON representation of the model
func (o *Cursors) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cursors) MarshalJSON() ([]byte, error) {
    type Alias Cursors

    if CursorsMarshalled {
        return []byte("{}"), nil
    }
    CursorsMarshalled = true

    return json.Marshal(&struct { 
        Before string `json:"before"`
        
        After string `json:"after"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

