package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DevelopmentactivityaggregatequeryrequestclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DevelopmentactivityaggregatequeryrequestclauseDud struct { 
    


    

}

// Developmentactivityaggregatequeryrequestclause
type Developmentactivityaggregatequeryrequestclause struct { 
    // VarType - The logic used to combine the predicates
    VarType string `json:"type"`


    // Predicates - The list of predicates used to filter the data
    Predicates []Developmentactivityaggregatequeryrequestpredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryrequestclause) String() string {
    
     o.Predicates = []Developmentactivityaggregatequeryrequestpredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Developmentactivityaggregatequeryrequestclause) MarshalJSON() ([]byte, error) {
    type Alias Developmentactivityaggregatequeryrequestclause

    if DevelopmentactivityaggregatequeryrequestclauseMarshalled {
        return []byte("{}"), nil
    }
    DevelopmentactivityaggregatequeryrequestclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Developmentactivityaggregatequeryrequestpredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Developmentactivityaggregatequeryrequestpredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

