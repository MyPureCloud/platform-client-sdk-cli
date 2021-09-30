package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AsyncqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AsyncqueryresponseDud struct { 
    

}

// Asyncqueryresponse
type Asyncqueryresponse struct { 
    // JobId - Unique identifier for the async query execution. Can be used to check the status of the query and retrieve results.
    JobId string `json:"jobId"`

}

// String returns a JSON representation of the model
func (o *Asyncqueryresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Asyncqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Asyncqueryresponse

    if AsyncqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    AsyncqueryresponseMarshalled = true

    return json.Marshal(&struct { 
        JobId string `json:"jobId"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

