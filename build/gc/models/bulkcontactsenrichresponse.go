package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkcontactsenrichresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkcontactsenrichresponseDud struct { 
    


    


    

}

// Bulkcontactsenrichresponse
type Bulkcontactsenrichresponse struct { 
    // Results - A list of results for all of the Bulk operations specified in the request. Includes both successes and failures. Ordering is NOT guaranteed - may be in a different order from the request.
    Results []Bulkresponseresultexternalcontactcontactenrichrequestbulkentityerrorcontactenrichrequest `json:"results"`


    // ErrorCount - The number of failed operations in the results.
    ErrorCount int `json:"errorCount"`


    // ErrorIndexes - The indexes of all failed operations in the results field.
    ErrorIndexes []int `json:"errorIndexes"`

}

// String returns a JSON representation of the model
func (o *Bulkcontactsenrichresponse) String() string {
     o.Results = []Bulkresponseresultexternalcontactcontactenrichrequestbulkentityerrorcontactenrichrequest{{}} 
    
     o.ErrorIndexes = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkcontactsenrichresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkcontactsenrichresponse

    if BulkcontactsenrichresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkcontactsenrichresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkresponseresultexternalcontactcontactenrichrequestbulkentityerrorcontactenrichrequest `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorIndexes []int `json:"errorIndexes"`
        *Alias
    }{

        
        Results: []Bulkresponseresultexternalcontactcontactenrichrequestbulkentityerrorcontactenrichrequest{{}},
        


        


        
        ErrorIndexes: []int{0},
        

        Alias: (*Alias)(u),
    })
}

