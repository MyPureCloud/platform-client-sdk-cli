package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementaggregatequeryclauseDud struct { 
    


    

}

// Taskmanagementaggregatequeryclause
type Taskmanagementaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Taskmanagementaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementaggregatequeryclause) String() string {
    
     o.Predicates = []Taskmanagementaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementaggregatequeryclause

    if TaskmanagementaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Taskmanagementaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Taskmanagementaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

