package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningslotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningslotDud struct { 
    


    


    


    

}

// Learningslot
type Learningslot struct { 
    // DateStart - Start date and time of scheduled Learning activity slot. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // LengthInMinutes - Length of Learning activity slot in minutes
    LengthInMinutes int `json:"lengthInMinutes"`


    // StaffingDifference - Difference between scheduled and forecast headcount for this slot after scheduling the Learning activity
    StaffingDifference float64 `json:"staffingDifference"`


    // DifferenceRating - Rating based on the staffing difference for scheduled slot
    DifferenceRating string `json:"differenceRating"`

}

// String returns a JSON representation of the model
func (o *Learningslot) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningslot) MarshalJSON() ([]byte, error) {
    type Alias Learningslot

    if LearningslotMarshalled {
        return []byte("{}"), nil
    }
    LearningslotMarshalled = true

    return json.Marshal(&struct {
        
        DateStart time.Time `json:"dateStart"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        StaffingDifference float64 `json:"staffingDifference"`
        
        DifferenceRating string `json:"differenceRating"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

