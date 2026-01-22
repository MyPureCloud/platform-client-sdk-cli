package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentrescheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentrescheduleDud struct { 
    


    


    

}

// Learningassignmentreschedule
type Learningassignmentreschedule struct { 
    // DateRecommendedForCompletion - The recommended completion date of the assignment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateRecommendedForCompletion time.Time `json:"dateRecommendedForCompletion"`


    // LengthInMinutes - The length in minutes of the assignment
    LengthInMinutes int `json:"lengthInMinutes"`


    // AddToSchedule - If True, adds the assignment to their schedule
    AddToSchedule bool `json:"addToSchedule"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentreschedule) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentreschedule) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentreschedule

    if LearningassignmentrescheduleMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentrescheduleMarshalled = true

    return json.Marshal(&struct {
        
        DateRecommendedForCompletion time.Time `json:"dateRecommendedForCompletion"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        AddToSchedule bool `json:"addToSchedule"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

