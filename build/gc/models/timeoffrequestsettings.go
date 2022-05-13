package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffrequestsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffrequestsettingsDud struct { 
    


    


    

}

// Timeoffrequestsettings - Time Off Request Settings
type Timeoffrequestsettings struct { 
    // SubmissionRangeEnforced - Whether to enforce a submission range for agent time off requests
    SubmissionRangeEnforced bool `json:"submissionRangeEnforced"`


    // SubmissionEarliestDaysFromNow - The earliest number of days from now for which an agent can submit a time off request.  Use negative numbers to indicate days in the past
    SubmissionEarliestDaysFromNow int `json:"submissionEarliestDaysFromNow"`


    // SubmissionLatestDaysFromNow - The latest number of days from now for which an agent can submit a time off request
    SubmissionLatestDaysFromNow int `json:"submissionLatestDaysFromNow"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestsettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffrequestsettings) MarshalJSON() ([]byte, error) {
    type Alias Timeoffrequestsettings

    if TimeoffrequestsettingsMarshalled {
        return []byte("{}"), nil
    }
    TimeoffrequestsettingsMarshalled = true

    return json.Marshal(&struct {
        
        SubmissionRangeEnforced bool `json:"submissionRangeEnforced"`
        
        SubmissionEarliestDaysFromNow int `json:"submissionEarliestDaysFromNow"`
        
        SubmissionLatestDaysFromNow int `json:"submissionLatestDaysFromNow"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

