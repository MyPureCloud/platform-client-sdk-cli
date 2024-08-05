package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshiftagentscheduledshiftMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshiftagentscheduledshiftDud struct { 
    


    


    


    


    

}

// Alternativeshiftagentscheduledshift
type Alternativeshiftagentscheduledshift struct { 
    // DayIndex - The number of days since start of schedule
    DayIndex int `json:"dayIndex"`


    // ReferenceKey - A key generated for an offer to help facilitate alternative shift trading
    ReferenceKey string `json:"referenceKey"`


    // StartDate - The start date of this shift in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of this shift in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Activities - A list of activities in this shift
    Activities []Buagentscheduleactivity `json:"activities"`

}

// String returns a JSON representation of the model
func (o *Alternativeshiftagentscheduledshift) String() string {
    
    
    
    
     o.Activities = []Buagentscheduleactivity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshiftagentscheduledshift) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshiftagentscheduledshift

    if AlternativeshiftagentscheduledshiftMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshiftagentscheduledshiftMarshalled = true

    return json.Marshal(&struct {
        
        DayIndex int `json:"dayIndex"`
        
        ReferenceKey string `json:"referenceKey"`
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Activities []Buagentscheduleactivity `json:"activities"`
        *Alias
    }{

        


        


        


        


        
        Activities: []Buagentscheduleactivity{{}},
        

        Alias: (*Alias)(u),
    })
}

