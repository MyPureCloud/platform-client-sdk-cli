package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyassignmentDud struct { 
    


    


    


    


    

}

// Surveyassignment
type Surveyassignment struct { 
    // SurveyForm - The survey form used for this survey.
    SurveyForm Publishedsurveyformreference `json:"surveyForm"`


    // Flow - The URI reference to the flow associated with this survey.
    Flow Domainentityref `json:"flow"`


    // InviteTimeInterval - An ISO 8601 repeated interval consisting of the number of repetitions, the start datetime, and the interval (e.g. R2/2018-03-01T13:00:00Z/P1M10DT2H30M). Total duration must not exceed 90 days.
    InviteTimeInterval string `json:"inviteTimeInterval"`


    // SendingUser - User together with sendingDomain used to send email, null to use no-reply
    SendingUser string `json:"sendingUser"`


    // SendingDomain - Validated email domain, required
    SendingDomain string `json:"sendingDomain"`

}

// String returns a JSON representation of the model
func (o *Surveyassignment) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyassignment) MarshalJSON() ([]byte, error) {
    type Alias Surveyassignment

    if SurveyassignmentMarshalled {
        return []byte("{}"), nil
    }
    SurveyassignmentMarshalled = true

    return json.Marshal(&struct { 
        SurveyForm Publishedsurveyformreference `json:"surveyForm"`
        
        Flow Domainentityref `json:"flow"`
        
        InviteTimeInterval string `json:"inviteTimeInterval"`
        
        SendingUser string `json:"sendingUser"`
        
        SendingDomain string `json:"sendingDomain"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

