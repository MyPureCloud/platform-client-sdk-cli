package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjobDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Bulkjob
type Bulkjob struct { 
    


    // State - The bulk job state.
    State string `json:"state"`


    // Action - The bulk job action. This determines what the bulk job does, for example, terminate workitems.
    Action string `json:"action"`


    // TotalCount - Total count of items to be processed in the bulk job.
    TotalCount int `json:"totalCount"`


    // SuccessfulCount - Count of successfully processed items in the bulk job.
    SuccessfulCount int `json:"successfulCount"`


    // FailedCount - Count of failed processed items in the bulk job.
    FailedCount int `json:"failedCount"`


    // DateStarted - The bulk job start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStarted time.Time `json:"dateStarted"`


    // DateFinished - The bulk job finished date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateFinished time.Time `json:"dateFinished"`


    // DateModified - The bulk job modification date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Bulkjob) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjob) MarshalJSON() ([]byte, error) {
    type Alias Bulkjob

    if BulkjobMarshalled {
        return []byte("{}"), nil
    }
    BulkjobMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        Action string `json:"action"`
        
        TotalCount int `json:"totalCount"`
        
        SuccessfulCount int `json:"successfulCount"`
        
        FailedCount int `json:"failedCount"`
        
        DateStarted time.Time `json:"dateStarted"`
        
        DateFinished time.Time `json:"dateFinished"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

