package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentsortitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentsortitemDud struct { 
    


    

}

// Contentsortitem
type Contentsortitem struct { 
    // Name
    Name string `json:"name"`


    // Ascending
    Ascending bool `json:"ascending"`

}

// String returns a JSON representation of the model
func (o *Contentsortitem) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentsortitem) MarshalJSON() ([]byte, error) {
    type Alias Contentsortitem

    if ContentsortitemMarshalled {
        return []byte("{}"), nil
    }
    ContentsortitemMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Ascending bool `json:"ascending"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

