package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommonruleconditionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommonruleconditionsDud struct { 
    


    


    


    

}

// Commonruleconditions
type Commonruleconditions struct { 
    // Clauses - The list of predicates groups to be evaluated
    Clauses []Commonruleconditions `json:"clauses"`


    // Predicates - The list of rule metric predicates to be evaluated.
    Predicates []Commonrulepredicate `json:"predicates"`


    // VarType - the logic operator performed.
    VarType string `json:"type"`


    // Id - The id.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Commonruleconditions) String() string {
     o.Clauses = []Commonruleconditions{{}} 
     o.Predicates = []Commonrulepredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commonruleconditions) MarshalJSON() ([]byte, error) {
    type Alias Commonruleconditions

    if CommonruleconditionsMarshalled {
        return []byte("{}"), nil
    }
    CommonruleconditionsMarshalled = true

    return json.Marshal(&struct {
        
        Clauses []Commonruleconditions `json:"clauses"`
        
        Predicates []Commonrulepredicate `json:"predicates"`
        
        VarType string `json:"type"`
        
        Id string `json:"id"`
        *Alias
    }{

        
        Clauses: []Commonruleconditions{{}},
        


        
        Predicates: []Commonrulepredicate{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

