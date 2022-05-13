package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryrequestclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryrequestclauseDud struct { 
    


    

}

// Queryrequestclause
type Queryrequestclause struct { 
    // VarType - The logic used to combine the predicates
    VarType string `json:"type"`


    // Predicates - The list of predicates used to filter the data
    Predicates []Queryrequestpredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Queryrequestclause) String() string {
    
     o.Predicates = []Queryrequestpredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryrequestclause) MarshalJSON() ([]byte, error) {
    type Alias Queryrequestclause

    if QueryrequestclauseMarshalled {
        return []byte("{}"), nil
    }
    QueryrequestclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Queryrequestpredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Queryrequestpredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

