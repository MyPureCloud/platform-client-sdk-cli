package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FeedbackaddrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FeedbackaddrequestDud struct { 
    


    

}

// Feedbackaddrequest
type Feedbackaddrequest struct { 
    // Summary - Agent's summary for the conversation
    Summary string `json:"summary"`


    // Rating - Agent’s rating for the system-generated summary.
    Rating string `json:"rating"`

}

// String returns a JSON representation of the model
func (o *Feedbackaddrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Feedbackaddrequest) MarshalJSON() ([]byte, error) {
    type Alias Feedbackaddrequest

    if FeedbackaddrequestMarshalled {
        return []byte("{}"), nil
    }
    FeedbackaddrequestMarshalled = true

    return json.Marshal(&struct {
        
        Summary string `json:"summary"`
        
        Rating string `json:"rating"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

