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
    // Results
    Results []Bulkresponseresultnoteentity `json:"results"`


    // ErrorCount
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkfetchnotesresponse) String() string {
    
    
     o.Results = []Bulkresponseresultnoteentity{{}} 
    
    
    
    
    
    
    
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
        Results []Bulkresponseresultnoteentity `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        
        *Alias
    }{
        

        
        Results: []Bulkresponseresultnoteentity{{}},
        

        

        

        

        
        ErrorIndexes: []int{0},
        

        
        Alias: (*Alias)(u),
    })
}

