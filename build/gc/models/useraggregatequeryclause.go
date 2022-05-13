package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseraggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseraggregatequeryclauseDud struct { 
    


    

}

// Useraggregatequeryclause
type Useraggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Useraggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Useraggregatequeryclause) String() string {
    
     o.Predicates = []Useraggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useraggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Useraggregatequeryclause

    if UseraggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    UseraggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Useraggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Useraggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

