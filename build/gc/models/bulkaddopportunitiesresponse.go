package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkaddopportunitiesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkaddopportunitiesresponseDud struct { 
    

}

// Bulkaddopportunitiesresponse
type Bulkaddopportunitiesresponse struct { 
    // Opportunities - The list of opportunities
    Opportunities []Opportunityresult `json:"opportunities"`

}

// String returns a JSON representation of the model
func (o *Bulkaddopportunitiesresponse) String() string {
     o.Opportunities = []Opportunityresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkaddopportunitiesresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkaddopportunitiesresponse

    if BulkaddopportunitiesresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkaddopportunitiesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Opportunities []Opportunityresult `json:"opportunities"`
        *Alias
    }{

        
        Opportunities: []Opportunityresult{{}},
        

        Alias: (*Alias)(u),
    })
}

