package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkfetchnotesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkfetchnotesresponseDud struct { 
    


    


    

}

// Bulkfetchnotesresponse
type Bulkfetchnotesresponse struct { 
    // Results - A list of results for all of the Bulk operations specified in the request. Includes both successes and failures. Ordering is NOT guaranteed - may be in a different order from the request.
    Results []Bulkresponseresultnoteentitybulkentityerrorentity `json:"results"`


    // ErrorCount - The number of failed operations in the results.
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes - The indexes of all failed operations in the results field.
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkfetchnotesresponse) String() string {
     o.Results = []Bulkresponseresultnoteentitybulkentityerrorentity{{}} 
    
     o.ErrorIndexes = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkfetchnotesresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkfetchnotesresponse

    if BulkfetchnotesresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkfetchnotesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkresponseresultnoteentitybulkentityerrorentity `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        *Alias
    }{

        
        Results: []Bulkresponseresultnoteentitybulkentityerrorentity{{}},
        


        


        
        ErrorIndexes: []int{0},
        

        Alias: (*Alias)(u),
    })
}

