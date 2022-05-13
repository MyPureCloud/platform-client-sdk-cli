package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulknotesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulknotesresponseDud struct { 
    


    


    

}

// Bulknotesresponse
type Bulknotesresponse struct { 
    // Results
    Results []Bulkresponseresultnotenote `json:"results"`


    // ErrorCount
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulknotesresponse) String() string {
     o.Results = []Bulkresponseresultnotenote{{}} 
    
     o.ErrorIndexes = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulknotesresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulknotesresponse

    if BulknotesresponseMarshalled {
        return []byte("{}"), nil
    }
    BulknotesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkresponseresultnotenote `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        *Alias
    }{

        
        Results: []Bulkresponseresultnotenote{{}},
        


        


        
        ErrorIndexes: []int{0},
        

        Alias: (*Alias)(u),
    })
}

