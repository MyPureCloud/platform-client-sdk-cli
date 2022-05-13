package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkfetchcontactsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkfetchcontactsresponseDud struct { 
    


    


    

}

// Bulkfetchcontactsresponse
type Bulkfetchcontactsresponse struct { 
    // Results
    Results []Bulkresponseresultexternalcontactentity `json:"results"`


    // ErrorCount
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkfetchcontactsresponse) String() string {
     o.Results = []Bulkresponseresultexternalcontactentity{{}} 
    
     o.ErrorIndexes = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkfetchcontactsresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkfetchcontactsresponse

    if BulkfetchcontactsresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkfetchcontactsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkresponseresultexternalcontactentity `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        *Alias
    }{

        
        Results: []Bulkresponseresultexternalcontactentity{{}},
        


        


        
        ErrorIndexes: []int{0},
        

        Alias: (*Alias)(u),
    })
}

