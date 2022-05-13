package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyaggregatequeryfilterDud struct { 
    


    


    

}

// Surveyaggregatequeryfilter
type Surveyaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Surveyaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Surveyaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Surveyaggregatequeryfilter) String() string {
    
     o.Clauses = []Surveyaggregatequeryclause{{}} 
     o.Predicates = []Surveyaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Surveyaggregatequeryfilter

    if SurveyaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    SurveyaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Surveyaggregatequeryclause `json:"clauses"`
        
        Predicates []Surveyaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Surveyaggregatequeryclause{{}},
        


        
        Predicates: []Surveyaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

