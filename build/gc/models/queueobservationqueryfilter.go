package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueueobservationqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueueobservationqueryfilterDud struct { 
    


    


    

}

// Queueobservationqueryfilter
type Queueobservationqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Queueobservationqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Queueobservationquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Queueobservationqueryfilter) String() string {
    
     o.Clauses = []Queueobservationqueryclause{{}} 
     o.Predicates = []Queueobservationquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queueobservationqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Queueobservationqueryfilter

    if QueueobservationqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    QueueobservationqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Queueobservationqueryclause `json:"clauses"`
        
        Predicates []Queueobservationquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Queueobservationqueryclause{{}},
        


        
        Predicates: []Queueobservationquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

