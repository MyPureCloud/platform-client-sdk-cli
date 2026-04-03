package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotaggregatequeryclauseDud struct { 
    


    

}

// Copilotaggregatequeryclause
type Copilotaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Copilotaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Copilotaggregatequeryclause) String() string {
    
     o.Predicates = []Copilotaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Copilotaggregatequeryclause

    if CopilotaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    CopilotaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Copilotaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Copilotaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

