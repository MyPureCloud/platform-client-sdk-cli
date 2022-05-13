package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotaggregatequeryclauseDud struct { 
    


    

}

// Botaggregatequeryclause
type Botaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Botaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Botaggregatequeryclause) String() string {
    
     o.Predicates = []Botaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Botaggregatequeryclause

    if BotaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    BotaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Botaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Botaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

