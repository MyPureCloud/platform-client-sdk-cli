package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasemanagementaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasemanagementaggregatequeryclauseDud struct { 
    


    

}

// Casemanagementaggregatequeryclause
type Casemanagementaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Casemanagementaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Casemanagementaggregatequeryclause) String() string {
    
     o.Predicates = []Casemanagementaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casemanagementaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Casemanagementaggregatequeryclause

    if CasemanagementaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    CasemanagementaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Casemanagementaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Casemanagementaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

