package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelfschedulequeryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelfschedulequeryDud struct { 
    

}

// Selfschedulequery
type Selfschedulequery struct { 
    // ActivityMoves - The self-schedule activity moves
    ActivityMoves []Selfscheduleactivitymovelistitem `json:"activityMoves"`

}

// String returns a JSON representation of the model
func (o *Selfschedulequery) String() string {
     o.ActivityMoves = []Selfscheduleactivitymovelistitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selfschedulequery) MarshalJSON() ([]byte, error) {
    type Alias Selfschedulequery

    if SelfschedulequeryMarshalled {
        return []byte("{}"), nil
    }
    SelfschedulequeryMarshalled = true

    return json.Marshal(&struct {
        
        ActivityMoves []Selfscheduleactivitymovelistitem `json:"activityMoves"`
        *Alias
    }{

        
        ActivityMoves: []Selfscheduleactivitymovelistitem{{}},
        

        Alias: (*Alias)(u),
    })
}

