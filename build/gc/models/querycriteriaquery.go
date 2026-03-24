package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuerycriteriaqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuerycriteriaqueryDud struct { 
    

}

// Querycriteriaquery - Query object for searching libraries based on criteria
type Querycriteriaquery struct { 
    // Query - List of criteria groups that will be AND'd together
    Query []Querycriteriagroup `json:"query"`

}

// String returns a JSON representation of the model
func (o *Querycriteriaquery) String() string {
     o.Query = []Querycriteriagroup{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Querycriteriaquery) MarshalJSON() ([]byte, error) {
    type Alias Querycriteriaquery

    if QuerycriteriaqueryMarshalled {
        return []byte("{}"), nil
    }
    QuerycriteriaqueryMarshalled = true

    return json.Marshal(&struct {
        
        Query []Querycriteriagroup `json:"query"`
        *Alias
    }{

        
        Query: []Querycriteriagroup{{}},
        

        Alias: (*Alias)(u),
    })
}

