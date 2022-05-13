package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkfetchorganizationsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkfetchorganizationsresponseDud struct { 
    


    


    

}

// Bulkfetchorganizationsresponse
type Bulkfetchorganizationsresponse struct { 
    // Results
    Results []Bulkresponseresultexternalorganizationentity `json:"results"`


    // ErrorCount
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkfetchorganizationsresponse) String() string {
     o.Results = []Bulkresponseresultexternalorganizationentity{{}} 
    
     o.ErrorIndexes = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkfetchorganizationsresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkfetchorganizationsresponse

    if BulkfetchorganizationsresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkfetchorganizationsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkresponseresultexternalorganizationentity `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        *Alias
    }{

        
        Results: []Bulkresponseresultexternalorganizationentity{{}},
        


        


        
        ErrorIndexes: []int{0},
        

        Alias: (*Alias)(u),
    })
}

