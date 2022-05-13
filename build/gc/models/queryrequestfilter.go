package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryrequestfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryrequestfilterDud struct { 
    


    

}

// Queryrequestfilter
type Queryrequestfilter struct { 
    // VarType - The logic used to combine the clauses
    VarType string `json:"type"`


    // Clauses - The list of clauses used to filter the data
    Clauses []Queryrequestclause `json:"clauses"`

}

// String returns a JSON representation of the model
func (o *Queryrequestfilter) String() string {
    
     o.Clauses = []Queryrequestclause{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryrequestfilter) MarshalJSON() ([]byte, error) {
    type Alias Queryrequestfilter

    if QueryrequestfilterMarshalled {
        return []byte("{}"), nil
    }
    QueryrequestfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Queryrequestclause `json:"clauses"`
        *Alias
    }{

        


        
        Clauses: []Queryrequestclause{{}},
        

        Alias: (*Alias)(u),
    })
}

