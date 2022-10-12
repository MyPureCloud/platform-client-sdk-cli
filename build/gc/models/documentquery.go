package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentqueryDud struct { 
    


    

}

// Documentquery
type Documentquery struct { 
    // Clauses - Documents filter clauses/criteria. Limit of 20 clauses.
    Clauses []Documentqueryclause `json:"clauses"`


    // Operator - Specifies how the filter clauses will be applied together.
    Operator string `json:"operator"`

}

// String returns a JSON representation of the model
func (o *Documentquery) String() string {
     o.Clauses = []Documentqueryclause{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentquery) MarshalJSON() ([]byte, error) {
    type Alias Documentquery

    if DocumentqueryMarshalled {
        return []byte("{}"), nil
    }
    DocumentqueryMarshalled = true

    return json.Marshal(&struct {
        
        Clauses []Documentqueryclause `json:"clauses"`
        
        Operator string `json:"operator"`
        *Alias
    }{

        
        Clauses: []Documentqueryclause{{}},
        


        

        Alias: (*Alias)(u),
    })
}

