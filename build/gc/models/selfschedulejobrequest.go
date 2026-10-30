package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelfschedulejobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelfschedulejobrequestDud struct { 
    

}

// Selfschedulejobrequest
type Selfschedulejobrequest struct { 
    // ActivityMoves - Activities based on the offer that the user wants to move
    ActivityMoves []Selfscheduleactivitymoverequestitem `json:"activityMoves"`

}

// String returns a JSON representation of the model
func (o *Selfschedulejobrequest) String() string {
     o.ActivityMoves = []Selfscheduleactivitymoverequestitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selfschedulejobrequest) MarshalJSON() ([]byte, error) {
    type Alias Selfschedulejobrequest

    if SelfschedulejobrequestMarshalled {
        return []byte("{}"), nil
    }
    SelfschedulejobrequestMarshalled = true

    return json.Marshal(&struct {
        
        ActivityMoves []Selfscheduleactivitymoverequestitem `json:"activityMoves"`
        *Alias
    }{

        
        ActivityMoves: []Selfscheduleactivitymoverequestitem{{}},
        

        Alias: (*Alias)(u),
    })
}

