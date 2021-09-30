package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowaggregatequeryclauseDud struct { 
    


    

}

// Flowaggregatequeryclause
type Flowaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Flowaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Flowaggregatequeryclause) String() string {
    
    
    
    
    
    
     o.Predicates = []Flowaggregatequerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Flowaggregatequeryclause

    if FlowaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    FlowaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Predicates []Flowaggregatequerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Predicates: []Flowaggregatequerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

