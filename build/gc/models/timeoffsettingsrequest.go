package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffsettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffsettingsrequestDud struct { 
    


    


    

}

// Timeoffsettingsrequest
type Timeoffsettingsrequest struct { 
    // SubmissionRangeEnforced - Whether to enforce a submission range for agent time off requests
    SubmissionRangeEnforced bool `json:"submissionRangeEnforced"`


    // SubmissionEarliestDaysFromNow - The earliest number of days from now for which an agent can submit a time off request.  Use negative numbers to indicate days in the past
    SubmissionEarliestDaysFromNow int `json:"submissionEarliestDaysFromNow"`


    // SubmissionLatestDaysFromNow - The latest number of days from now for which an agent can submit a time off request
    SubmissionLatestDaysFromNow int `json:"submissionLatestDaysFromNow"`

}

// String returns a JSON representation of the model
func (o *Timeoffsettingsrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffsettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Timeoffsettingsrequest

    if TimeoffsettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    TimeoffsettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        SubmissionRangeEnforced bool `json:"submissionRangeEnforced"`
        
        SubmissionEarliestDaysFromNow int `json:"submissionEarliestDaysFromNow"`
        
        SubmissionLatestDaysFromNow int `json:"submissionLatestDaysFromNow"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

