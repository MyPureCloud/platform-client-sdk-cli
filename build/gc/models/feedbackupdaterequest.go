package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FeedbackupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FeedbackupdaterequestDud struct { 
    


    

}

// Feedbackupdaterequest
type Feedbackupdaterequest struct { 
    // Rating - Agent’s rating for the system-generated summary.
    Rating string `json:"rating"`


    // Summary - Agent's summary for the conversation
    Summary string `json:"summary"`

}

// String returns a JSON representation of the model
func (o *Feedbackupdaterequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Feedbackupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Feedbackupdaterequest

    if FeedbackupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    FeedbackupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Rating string `json:"rating"`
        
        Summary string `json:"summary"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

