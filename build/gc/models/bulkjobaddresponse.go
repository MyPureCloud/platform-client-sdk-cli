package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjobaddresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjobaddresponseDud struct { 
    


    


    

}

// Bulkjobaddresponse
type Bulkjobaddresponse struct { 
    // Results - A list of the results from the bulk operation.
    Results []Bulkjobaddresult `json:"results"`


    // ErrorCount - The number of errors from the bulk operation.
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes - An index of where the errors are in the listing.
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkjobaddresponse) String() string {
     o.Results = []Bulkjobaddresult{{}} 
    
     o.ErrorIndexes = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjobaddresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkjobaddresponse

    if BulkjobaddresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkjobaddresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkjobaddresult `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        *Alias
    }{

        
        Results: []Bulkjobaddresult{{}},
        


        


        
        ErrorIndexes: []int{0},
        

        Alias: (*Alias)(u),
    })
}

