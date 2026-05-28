package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkopportunitiesenrollmentresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkopportunitiesenrollmentresultDud struct { 
    


    


    

}

// Bulkopportunitiesenrollmentresult
type Bulkopportunitiesenrollmentresult struct { 
    // Status - The status indicating the result of the bulk operation for this item
    Status string `json:"status"`


    // VarError - The error result if the operation failed
    VarError Bulkopportunitieserror `json:"error"`


    // Enrollment - The enrollment result
    Enrollment Opportunityenrollment `json:"enrollment"`

}

// String returns a JSON representation of the model
func (o *Bulkopportunitiesenrollmentresult) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkopportunitiesenrollmentresult) MarshalJSON() ([]byte, error) {
    type Alias Bulkopportunitiesenrollmentresult

    if BulkopportunitiesenrollmentresultMarshalled {
        return []byte("{}"), nil
    }
    BulkopportunitiesenrollmentresultMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        VarError Bulkopportunitieserror `json:"error"`
        
        Enrollment Opportunityenrollment `json:"enrollment"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

