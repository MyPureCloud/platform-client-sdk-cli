package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KpiattributionaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KpiattributionaggregatequeryclauseDud struct { 
    


    

}

// Kpiattributionaggregatequeryclause
type Kpiattributionaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Kpiattributionaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Kpiattributionaggregatequeryclause) String() string {
    
     o.Predicates = []Kpiattributionaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Kpiattributionaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Kpiattributionaggregatequeryclause

    if KpiattributionaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    KpiattributionaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Kpiattributionaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Kpiattributionaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

