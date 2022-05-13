package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentaggregatequeryrequestclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentaggregatequeryrequestclauseDud struct { 
    


    

}

// Learningassignmentaggregatequeryrequestclause
type Learningassignmentaggregatequeryrequestclause struct { 
    // VarType - The logic used to combine the predicates
    VarType string `json:"type"`


    // Predicates - The list of predicates used to filter the data
    Predicates []Learningassignmentaggregatequeryrequestpredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryrequestclause) String() string {
    
     o.Predicates = []Learningassignmentaggregatequeryrequestpredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentaggregatequeryrequestclause) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentaggregatequeryrequestclause

    if LearningassignmentaggregatequeryrequestclauseMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentaggregatequeryrequestclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Learningassignmentaggregatequeryrequestpredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Learningassignmentaggregatequeryrequestpredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

