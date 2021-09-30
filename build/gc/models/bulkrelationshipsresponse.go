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
    // Results
    Results []Bulkresponseresultrelationshiprelationship `json:"results"`


    // ErrorCount
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkrelationshipsresponse) String() string {
    
    
     o.Results = []Bulkresponseresultrelationshiprelationship{{}} 
    
    
    
    
    
    
    
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
        Results []Bulkresponseresultrelationshiprelationship `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        
        *Alias
    }{
        

        
        Results: []Bulkresponseresultrelationshiprelationship{{}},
        

        

        

        

        
        ErrorIndexes: []int{0},
        

        
        Alias: (*Alias)(u),
    })
}

