package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjobterminateresultsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjobterminateresultsresponseDud struct { 
    


    


    

}

// Bulkjobterminateresultsresponse
type Bulkjobterminateresultsresponse struct { 
    // Results - A list of the results from the bulk operation.
    Results []Bulkjobterminateresult `json:"results"`


    // ErrorCount - The number of errors from the bulk operation.
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes - An index of where the errors are in the listing.
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkjobterminateresultsresponse) String() string {
     o.Results = []Bulkjobterminateresult{{}} 
    
     o.ErrorIndexes = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjobterminateresultsresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkjobterminateresultsresponse

    if BulkjobterminateresultsresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkjobterminateresultsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkjobterminateresult `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        *Alias
    }{

        
        Results: []Bulkjobterminateresult{{}},
        


        


        
        ErrorIndexes: []int{0},
        

        Alias: (*Alias)(u),
    })
}

