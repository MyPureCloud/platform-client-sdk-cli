package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowobservationqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowobservationqueryfilterDud struct { 
    


    


    

}

// Flowobservationqueryfilter
type Flowobservationqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Flowobservationqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Flowobservationquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Flowobservationqueryfilter) String() string {
    
     o.Clauses = []Flowobservationqueryclause{{}} 
     o.Predicates = []Flowobservationquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowobservationqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Flowobservationqueryfilter

    if FlowobservationqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    FlowobservationqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Flowobservationqueryclause `json:"clauses"`
        
        Predicates []Flowobservationquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Flowobservationqueryclause{{}},
        


        
        Predicates: []Flowobservationquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

