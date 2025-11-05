package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffsettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffsettingsresponseDud struct { 
    


    


    


    


    

}

// Timeoffsettingsresponse
type Timeoffsettingsresponse struct { 
    // SubmissionRangeEnforced - Whether to enforce a submission range for agent time off requests
    SubmissionRangeEnforced bool `json:"submissionRangeEnforced"`


    // SubmissionRangeType - The type of the submission range
    SubmissionRangeType string `json:"submissionRangeType"`


    // SubmissionEarliestDaysFromNow - The earliest number of days from now for which an agent can submit a time off request.  Use negative numbers to indicate days in the past
    SubmissionEarliestDaysFromNow int `json:"submissionEarliestDaysFromNow"`


    // SubmissionLatestDaysFromNow - The latest number of days from now for which an agent can submit a time off request
    SubmissionLatestDaysFromNow int `json:"submissionLatestDaysFromNow"`


    // SubmissionLatestDate - The latest date for the time off request submission interpreted in the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    SubmissionLatestDate time.Time `json:"submissionLatestDate"`

}

// String returns a JSON representation of the model
func (o *Timeoffsettingsresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffsettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Timeoffsettingsresponse

    if TimeoffsettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    TimeoffsettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        SubmissionRangeEnforced bool `json:"submissionRangeEnforced"`
        
        SubmissionRangeType string `json:"submissionRangeType"`
        
        SubmissionEarliestDaysFromNow int `json:"submissionEarliestDaysFromNow"`
        
        SubmissionLatestDaysFromNow int `json:"submissionLatestDaysFromNow"`
        
        SubmissionLatestDate time.Time `json:"submissionLatestDate"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

