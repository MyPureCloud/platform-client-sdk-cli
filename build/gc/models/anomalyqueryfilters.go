package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalyqueryfiltersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalyqueryfiltersDud struct { 
    


    


    

}

// Anomalyqueryfilters
type Anomalyqueryfilters struct { 
    // Operator - Boolean operation to apply to the provided predicates and clauses
    Operator string `json:"operator"`


    // Clauses - Boolean 'And/Or' logic with up to two levels of nesting. At least one of clauses or predicates is required
    Clauses []Anomalyqueryclause `json:"clauses"`


    // Predicates - Predicates to include data in the resulting query. At least one of clauses or predicates is required
    Predicates []Anomalyquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Anomalyqueryfilters) String() string {
    
     o.Clauses = []Anomalyqueryclause{{}} 
     o.Predicates = []Anomalyquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalyqueryfilters) MarshalJSON() ([]byte, error) {
    type Alias Anomalyqueryfilters

    if AnomalyqueryfiltersMarshalled {
        return []byte("{}"), nil
    }
    AnomalyqueryfiltersMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        Clauses []Anomalyqueryclause `json:"clauses"`
        
        Predicates []Anomalyquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Anomalyqueryclause{{}},
        


        
        Predicates: []Anomalyquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

