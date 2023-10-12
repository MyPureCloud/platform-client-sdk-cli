package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementaggregatequeryfilterDud struct { 
    


    


    

}

// Taskmanagementaggregatequeryfilter
type Taskmanagementaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Taskmanagementaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Taskmanagementaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementaggregatequeryfilter) String() string {
    
     o.Clauses = []Taskmanagementaggregatequeryclause{{}} 
     o.Predicates = []Taskmanagementaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementaggregatequeryfilter

    if TaskmanagementaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Taskmanagementaggregatequeryclause `json:"clauses"`
        
        Predicates []Taskmanagementaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Taskmanagementaggregatequeryclause{{}},
        


        
        Predicates: []Taskmanagementaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

