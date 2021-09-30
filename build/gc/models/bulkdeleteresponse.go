package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkdeleteresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkdeleteresponseDud struct { 
    


    


    

}

// Bulkdeleteresponse
type Bulkdeleteresponse struct { 
    // Results
    Results []Bulkresponseresultvoidentity `json:"results"`


    // ErrorCount
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkdeleteresponse) String() string {
    
    
     o.Results = []Bulkresponseresultvoidentity{{}} 
    
    
    
    
    
    
    
     o.ErrorIndexes = []int{0} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkdeleteresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkdeleteresponse

    if BulkdeleteresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkdeleteresponseMarshalled = true

    return json.Marshal(&struct { 
        Results []Bulkresponseresultvoidentity `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        
        *Alias
    }{
        

        
        Results: []Bulkresponseresultvoidentity{{}},
        

        

        

        

        
        ErrorIndexes: []int{0},
        

        
        Alias: (*Alias)(u),
    })
}

