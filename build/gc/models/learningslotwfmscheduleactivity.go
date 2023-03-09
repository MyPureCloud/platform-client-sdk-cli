package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningslotwfmscheduleactivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningslotwfmscheduleactivityDud struct { 
    


    


    

}

// Learningslotwfmscheduleactivity
type Learningslotwfmscheduleactivity struct { 
    // User - User that the schedule is for
    User Userreference `json:"user"`


    // Activities - List of user's scheduled activities
    Activities []Learningslotscheduleactivity `json:"activities"`


    // FullDayTimeOffMarkers - List of user's days off
    FullDayTimeOffMarkers []Learningslotfulldaytimeoffmarker `json:"fullDayTimeOffMarkers"`

}

// String returns a JSON representation of the model
func (o *Learningslotwfmscheduleactivity) String() string {
    
     o.Activities = []Learningslotscheduleactivity{{}} 
     o.FullDayTimeOffMarkers = []Learningslotfulldaytimeoffmarker{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningslotwfmscheduleactivity) MarshalJSON() ([]byte, error) {
    type Alias Learningslotwfmscheduleactivity

    if LearningslotwfmscheduleactivityMarshalled {
        return []byte("{}"), nil
    }
    LearningslotwfmscheduleactivityMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        Activities []Learningslotscheduleactivity `json:"activities"`
        
        FullDayTimeOffMarkers []Learningslotfulldaytimeoffmarker `json:"fullDayTimeOffMarkers"`
        *Alias
    }{

        


        
        Activities: []Learningslotscheduleactivity{{}},
        


        
        FullDayTimeOffMarkers: []Learningslotfulldaytimeoffmarker{{}},
        

        Alias: (*Alias)(u),
    })
}

