package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkorganizationsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkorganizationsresponseDud struct { 
    


    


    

}

// Bulkorganizationsresponse
type Bulkorganizationsresponse struct { 
    // Results
    Results []Bulkresponseresultexternalorganizationexternalorganization `json:"results"`


    // ErrorCount
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkorganizationsresponse) String() string {
     o.Results = []Bulkresponseresultexternalorganizationexternalorganization{{}} 
    
     o.ErrorIndexes = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkorganizationsresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkorganizationsresponse

    if BulkorganizationsresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkorganizationsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkresponseresultexternalorganizationexternalorganization `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        *Alias
    }{

        
        Results: []Bulkresponseresultexternalorganizationexternalorganization{{}},
        


        


        
        ErrorIndexes: []int{0},
        

        Alias: (*Alias)(u),
    })
}

