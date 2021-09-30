package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotuseragentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotuseragentDud struct { 
    

}

// Textbotuseragent - Information about the caller executing a bot flow.
type Textbotuseragent struct { 
    // Name - The name of the user agent.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Textbotuseragent) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotuseragent) MarshalJSON() ([]byte, error) {
    type Alias Textbotuseragent

    if TextbotuseragentMarshalled {
        return []byte("{}"), nil
    }
    TextbotuseragentMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

