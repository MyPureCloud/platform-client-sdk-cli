package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkupdateopportunityenrollmentsstatusresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkupdateopportunityenrollmentsstatusresponseDud struct { 
    


    

}

// Bulkupdateopportunityenrollmentsstatusresponse
type Bulkupdateopportunityenrollmentsstatusresponse struct { 
    // Results - The result for each requested item
    Results []Bulkopportunitiesenrollmentresult `json:"results"`


    // ErrorCount - The count of failed operations in the bulk request
    ErrorCount int `json:"errorCount"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateopportunityenrollmentsstatusresponse) String() string {
     o.Results = []Bulkopportunitiesenrollmentresult{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkupdateopportunityenrollmentsstatusresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkupdateopportunityenrollmentsstatusresponse

    if BulkupdateopportunityenrollmentsstatusresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkupdateopportunityenrollmentsstatusresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkopportunitiesenrollmentresult `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        *Alias
    }{

        
        Results: []Bulkopportunitiesenrollmentresult{{}},
        


        

        Alias: (*Alias)(u),
    })
}

