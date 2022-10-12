package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentqueryclauseDud struct { 
    


    

}

// Documentqueryclause
type Documentqueryclause struct { 
    // Operator - Specifies how the predicates will be applied together.
    Operator string `json:"operator"`


    // Predicates - To apply multiple conditions. Limit of 10 predicates across all clauses.
    Predicates []Documentquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Documentqueryclause) String() string {
    
     o.Predicates = []Documentquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Documentqueryclause

    if DocumentqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    DocumentqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        Predicates []Documentquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Documentquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

