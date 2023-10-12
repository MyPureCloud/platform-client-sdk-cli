package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomeattributionjobstateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomeattributionjobstateresponseDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`


    

}

// Outcomeattributionjobstateresponse
type Outcomeattributionjobstateresponse struct { 
    


    // State - State of job.
    State string `json:"state"`


    // ResultsUri - URI where the query results can be retrieved.  Populated when job has reached a terminal state, i.e. no longer Running.
    ResultsUri string `json:"resultsUri"`


    // PercentFailedThreshold - Optional percent failed threshold for validation errors; if reached will halt the job. Default is 5 percent, allowed values 0 to 100.
    PercentFailedThreshold int `json:"percentFailedThreshold"`


    


    // CreatedDate - Date when the job was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`

}

// String returns a JSON representation of the model
func (o *Outcomeattributionjobstateresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomeattributionjobstateresponse) MarshalJSON() ([]byte, error) {
    type Alias Outcomeattributionjobstateresponse

    if OutcomeattributionjobstateresponseMarshalled {
        return []byte("{}"), nil
    }
    OutcomeattributionjobstateresponseMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        ResultsUri string `json:"resultsUri"`
        
        PercentFailedThreshold int `json:"percentFailedThreshold"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

