package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkopportunitiesreferenceresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkopportunitiesreferenceresultDud struct { 
    


    


    

}

// Bulkopportunitiesreferenceresult
type Bulkopportunitiesreferenceresult struct { 
    // Status - The status indicating the result of the bulk operation for this item
    Status string `json:"status"`


    // VarError - The error result if the operation failed
    VarError Bulkopportunitieserror `json:"error"`


    // Opportunity - Reference to the opportunity
    Opportunity Opportunityreference `json:"opportunity"`

}

// String returns a JSON representation of the model
func (o *Bulkopportunitiesreferenceresult) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkopportunitiesreferenceresult) MarshalJSON() ([]byte, error) {
    type Alias Bulkopportunitiesreferenceresult

    if BulkopportunitiesreferenceresultMarshalled {
        return []byte("{}"), nil
    }
    BulkopportunitiesreferenceresultMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        VarError Bulkopportunitieserror `json:"error"`
        
        Opportunity Opportunityreference `json:"opportunity"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

