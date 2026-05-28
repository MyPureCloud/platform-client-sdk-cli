package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkopportunitiesstatusupdateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkopportunitiesstatusupdateresponseDud struct { 
    


    

}

// Bulkopportunitiesstatusupdateresponse
type Bulkopportunitiesstatusupdateresponse struct { 
    // Results - The result for each requested item
    Results []Bulkopportunitiesreferenceresult `json:"results"`


    // ErrorCount - The count of failed operations in the bulk request
    ErrorCount int `json:"errorCount"`

}

// String returns a JSON representation of the model
func (o *Bulkopportunitiesstatusupdateresponse) String() string {
     o.Results = []Bulkopportunitiesreferenceresult{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkopportunitiesstatusupdateresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkopportunitiesstatusupdateresponse

    if BulkopportunitiesstatusupdateresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkopportunitiesstatusupdateresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkopportunitiesreferenceresult `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        *Alias
    }{

        
        Results: []Bulkopportunitiesreferenceresult{{}},
        


        

        Alias: (*Alias)(u),
    })
}

