package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelfscheduleoffersresulttemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelfscheduleoffersresulttemplateDud struct { 
    

}

// Selfscheduleoffersresulttemplate
type Selfscheduleoffersresulttemplate struct { 
    // ActivityMoves - Self-schedule activity moves
    ActivityMoves []Selfscheduleoffersactivitymovetemplate `json:"activityMoves"`

}

// String returns a JSON representation of the model
func (o *Selfscheduleoffersresulttemplate) String() string {
     o.ActivityMoves = []Selfscheduleoffersactivitymovetemplate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selfscheduleoffersresulttemplate) MarshalJSON() ([]byte, error) {
    type Alias Selfscheduleoffersresulttemplate

    if SelfscheduleoffersresulttemplateMarshalled {
        return []byte("{}"), nil
    }
    SelfscheduleoffersresulttemplateMarshalled = true

    return json.Marshal(&struct {
        
        ActivityMoves []Selfscheduleoffersactivitymovetemplate `json:"activityMoves"`
        *Alias
    }{

        
        ActivityMoves: []Selfscheduleoffersactivitymovetemplate{{}},
        

        Alias: (*Alias)(u),
    })
}

