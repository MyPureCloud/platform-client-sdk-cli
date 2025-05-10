package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RowsearchfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RowsearchfilterDud struct { 
    

}

// Rowsearchfilter
type Rowsearchfilter struct { 
    // Predicates - The predicates that each row matches in the search results must match
    Predicates []Rowsearchpredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Rowsearchfilter) String() string {
     o.Predicates = []Rowsearchpredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Rowsearchfilter) MarshalJSON() ([]byte, error) {
    type Alias Rowsearchfilter

    if RowsearchfilterMarshalled {
        return []byte("{}"), nil
    }
    RowsearchfilterMarshalled = true

    return json.Marshal(&struct {
        
        Predicates []Rowsearchpredicate `json:"predicates"`
        *Alias
    }{

        
        Predicates: []Rowsearchpredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

