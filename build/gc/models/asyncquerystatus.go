package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AsyncquerystatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AsyncquerystatusDud struct { 
    


    


    


    


    

}

// Asyncquerystatus
type Asyncquerystatus struct { 
    // State - The current state of the asynchronous query
    State string `json:"state"`


    // ErrorMessage - The error associated with the current query, if the state is FAILED
    ErrorMessage string `json:"errorMessage"`


    // ExpirationDate - The time at which results for this query will expire. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ExpirationDate time.Time `json:"expirationDate"`


    // SubmissionDate - The time at which the query was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    SubmissionDate time.Time `json:"submissionDate"`


    // CompletionDate - The time at which the query completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CompletionDate time.Time `json:"completionDate"`

}

// String returns a JSON representation of the model
func (o *Asyncquerystatus) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Asyncquerystatus) MarshalJSON() ([]byte, error) {
    type Alias Asyncquerystatus

    if AsyncquerystatusMarshalled {
        return []byte("{}"), nil
    }
    AsyncquerystatusMarshalled = true

    return json.Marshal(&struct { 
        State string `json:"state"`
        
        ErrorMessage string `json:"errorMessage"`
        
        ExpirationDate time.Time `json:"expirationDate"`
        
        SubmissionDate time.Time `json:"submissionDate"`
        
        CompletionDate time.Time `json:"completionDate"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

