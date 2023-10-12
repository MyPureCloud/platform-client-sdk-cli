package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResolutionaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResolutionaggregatequeryclauseDud struct { 
    


    

}

// Resolutionaggregatequeryclause
type Resolutionaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Resolutionaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Resolutionaggregatequeryclause) String() string {
    
     o.Predicates = []Resolutionaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Resolutionaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Resolutionaggregatequeryclause

    if ResolutionaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    ResolutionaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Resolutionaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Resolutionaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

