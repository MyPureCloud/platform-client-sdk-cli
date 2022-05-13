package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SessionsegmentassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SessionsegmentassignmentDud struct { 
    


    

}

// Sessionsegmentassignment
type Sessionsegmentassignment struct { 
    // Segment - The segment that was assigned.
    Segment Assignedsegment `json:"segment"`


    // AssignedDate - Timestamp indicating when the segment was assigned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    AssignedDate time.Time `json:"assignedDate"`

}

// String returns a JSON representation of the model
func (o *Sessionsegmentassignment) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sessionsegmentassignment) MarshalJSON() ([]byte, error) {
    type Alias Sessionsegmentassignment

    if SessionsegmentassignmentMarshalled {
        return []byte("{}"), nil
    }
    SessionsegmentassignmentMarshalled = true

    return json.Marshal(&struct {
        
        Segment Assignedsegment `json:"segment"`
        
        AssignedDate time.Time `json:"assignedDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

