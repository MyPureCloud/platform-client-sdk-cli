package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkremoveopportunitiesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkremoveopportunitiesresponseDud struct { 
    


    

}

// Bulkremoveopportunitiesresponse
type Bulkremoveopportunitiesresponse struct { 
    // Results - The result for each requested item
    Results []Bulkremoveopportunitiesresult `json:"results"`


    // ErrorCount - The count of failed operations in the bulk request
    ErrorCount int `json:"errorCount"`

}

// String returns a JSON representation of the model
func (o *Bulkremoveopportunitiesresponse) String() string {
     o.Results = []Bulkremoveopportunitiesresult{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkremoveopportunitiesresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkremoveopportunitiesresponse

    if BulkremoveopportunitiesresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkremoveopportunitiesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkremoveopportunitiesresult `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        *Alias
    }{

        
        Results: []Bulkremoveopportunitiesresult{{}},
        


        

        Alias: (*Alias)(u),
    })
}

