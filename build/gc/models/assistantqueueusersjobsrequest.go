package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistantqueueusersjobsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistantqueueusersjobsrequestDud struct { 
    

}

// Assistantqueueusersjobsrequest
type Assistantqueueusersjobsrequest struct { 
    // Action - Action to perform by the job.
    Action string `json:"action"`

}

// String returns a JSON representation of the model
func (o *Assistantqueueusersjobsrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistantqueueusersjobsrequest) MarshalJSON() ([]byte, error) {
    type Alias Assistantqueueusersjobsrequest

    if AssistantqueueusersjobsrequestMarshalled {
        return []byte("{}"), nil
    }
    AssistantqueueusersjobsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Action string `json:"action"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

