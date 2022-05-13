package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingslotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingslotDud struct { 
    DateStart time.Time `json:"dateStart"`


    LengthInMinutes int `json:"lengthInMinutes"`


    StaffingDifference float64 `json:"staffingDifference"`


    DifferenceRating string `json:"differenceRating"`


    WfmSchedule Wfmschedulereference `json:"wfmSchedule"`

}

// Coachingslot
type Coachingslot struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Coachingslot) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingslot) MarshalJSON() ([]byte, error) {
    type Alias Coachingslot

    if CoachingslotMarshalled {
        return []byte("{}"), nil
    }
    CoachingslotMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

