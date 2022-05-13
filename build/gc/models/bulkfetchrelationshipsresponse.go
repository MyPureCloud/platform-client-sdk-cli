package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkfetchrelationshipsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkfetchrelationshipsresponseDud struct { 
    


    


    

}

// Bulkfetchrelationshipsresponse
type Bulkfetchrelationshipsresponse struct { 
    // Results
    Results []Bulkresponseresultrelationshipentity `json:"results"`


    // ErrorCount
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkfetchrelationshipsresponse) String() string {
     o.Results = []Bulkresponseresultrelationshipentity{{}} 
    
     o.ErrorIndexes = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkfetchrelationshipsresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkfetchrelationshipsresponse

    if BulkfetchrelationshipsresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkfetchrelationshipsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkresponseresultrelationshipentity `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        *Alias
    }{

        
        Results: []Bulkresponseresultrelationshipentity{{}},
        


        


        
        ErrorIndexes: []int{0},
        

        Alias: (*Alias)(u),
    })
}

