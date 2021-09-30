package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContextDud struct { 
    

}

// Context
type Context struct { 
    // Patterns - A list of one or more patterns to match.
    Patterns []Contextpattern `json:"patterns"`

}

// String returns a JSON representation of the model
func (o *Context) String() string {
    
    
     o.Patterns = []Contextpattern{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Context) MarshalJSON() ([]byte, error) {
    type Alias Context

    if ContextMarshalled {
        return []byte("{}"), nil
    }
    ContextMarshalled = true

    return json.Marshal(&struct { 
        Patterns []Contextpattern `json:"patterns"`
        
        *Alias
    }{
        

        
        Patterns: []Contextpattern{{}},
        

        
        Alias: (*Alias)(u),
    })
}

