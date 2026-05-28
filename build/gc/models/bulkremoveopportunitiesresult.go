package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkremoveopportunitiesresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkremoveopportunitiesresultDud struct { 
    


    


    

}

// Bulkremoveopportunitiesresult
type Bulkremoveopportunitiesresult struct { 
    // Status - The status indicating the result of the bulk operation for this item
    Status string `json:"status"`


    // VarError - The error result if the operation failed
    VarError Bulkopportunitieserror `json:"error"`


    // OpportunityId - The ID of the opportunity
    OpportunityId string `json:"opportunityId"`

}

// String returns a JSON representation of the model
func (o *Bulkremoveopportunitiesresult) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkremoveopportunitiesresult) MarshalJSON() ([]byte, error) {
    type Alias Bulkremoveopportunitiesresult

    if BulkremoveopportunitiesresultMarshalled {
        return []byte("{}"), nil
    }
    BulkremoveopportunitiesresultMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        VarError Bulkopportunitieserror `json:"error"`
        
        OpportunityId string `json:"opportunityId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

