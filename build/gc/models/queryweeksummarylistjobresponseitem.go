package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryweeksummarylistjobresponseitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryweeksummarylistjobresponseitemDud struct { 
    

}

// Queryweeksummarylistjobresponseitem
type Queryweeksummarylistjobresponseitem struct { 
    // Weeks - Weekly summary counts of the trades for the requested weeks
    Weeks []Shifttradeweeksummaryresponseitem `json:"weeks"`

}

// String returns a JSON representation of the model
func (o *Queryweeksummarylistjobresponseitem) String() string {
     o.Weeks = []Shifttradeweeksummaryresponseitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryweeksummarylistjobresponseitem) MarshalJSON() ([]byte, error) {
    type Alias Queryweeksummarylistjobresponseitem

    if QueryweeksummarylistjobresponseitemMarshalled {
        return []byte("{}"), nil
    }
    QueryweeksummarylistjobresponseitemMarshalled = true

    return json.Marshal(&struct {
        
        Weeks []Shifttradeweeksummaryresponseitem `json:"weeks"`
        *Alias
    }{

        
        Weeks: []Shifttradeweeksummaryresponseitem{{}},
        

        Alias: (*Alias)(u),
    })
}

