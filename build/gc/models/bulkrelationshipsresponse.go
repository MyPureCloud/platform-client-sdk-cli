package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkrelationshipsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkrelationshipsresponseDud struct { 
    


    


    

}

// Bulkrelationshipsresponse
type Bulkrelationshipsresponse struct { 
    // Results - A list of results for all of the Bulk operations specified in the request. Includes both successes and failures. Ordering is NOT guaranteed - may be in a different order from the request.
    Results []Bulkresponseresultrelationshiprelationshipbulkentityerrorrelationship `json:"results"`


    // ErrorCount - The number of failed operations in the results.
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes - The indexes of all failed operations in the results field.
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkrelationshipsresponse) String() string {
     o.Results = []Bulkresponseresultrelationshiprelationshipbulkentityerrorrelationship{{}} 
    
     o.ErrorIndexes = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkrelationshipsresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkrelationshipsresponse

    if BulkrelationshipsresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkrelationshipsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkresponseresultrelationshiprelationshipbulkentityerrorrelationship `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        *Alias
    }{

        
        Results: []Bulkresponseresultrelationshiprelationshipbulkentityerrorrelationship{{}},
        


        


        
        ErrorIndexes: []int{0},
        

        Alias: (*Alias)(u),
    })
}

