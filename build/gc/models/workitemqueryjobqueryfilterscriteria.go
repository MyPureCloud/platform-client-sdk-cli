package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemqueryjobqueryfilterscriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemqueryjobqueryfilterscriteriaDud struct { 
    


    

}

// Workitemqueryjobqueryfilterscriteria
type Workitemqueryjobqueryfilterscriteria struct { 
    // Operator - Query filter logical operator to join predicates.
    Operator string `json:"operator"`


    // Predicates - Query filter predicates. Number of predicates within the query filter should be between 1 and 5.
    Predicates []Workitemqueryjobqueryfilterspredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Workitemqueryjobqueryfilterscriteria) String() string {
    
     o.Predicates = []Workitemqueryjobqueryfilterspredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemqueryjobqueryfilterscriteria) MarshalJSON() ([]byte, error) {
    type Alias Workitemqueryjobqueryfilterscriteria

    if WorkitemqueryjobqueryfilterscriteriaMarshalled {
        return []byte("{}"), nil
    }
    WorkitemqueryjobqueryfilterscriteriaMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        Predicates []Workitemqueryjobqueryfilterspredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Workitemqueryjobqueryfilterspredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

