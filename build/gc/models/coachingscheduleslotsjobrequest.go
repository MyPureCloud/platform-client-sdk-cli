package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingscheduleslotsjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingscheduleslotsjobrequestDud struct { 
    


    


    


    


    


    

}

// Coachingscheduleslotsjobrequest
type Coachingscheduleslotsjobrequest struct { 
    // AttendeeIds - The attendee IDs to fetch the slots for.
    AttendeeIds []string `json:"attendeeIds"`


    // FacilitatorIds - The facilitator IDs to fetch the slots for.
    FacilitatorIds []string `json:"facilitatorIds"`


    // LengthInMinutes - The length in minutes of the slots, in 15 minutes granularity.
    LengthInMinutes int `json:"lengthInMinutes"`


    // ActivityCodeId - The Activity Code Id of the slots.
    ActivityCodeId string `json:"activityCodeId"`


    // Intervals - The intervals to fetch the slots for. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Intervals []string `json:"intervals"`


    // SlotsType - The type of slots to fetch in the job.
    SlotsType string `json:"slotsType"`

}

// String returns a JSON representation of the model
func (o *Coachingscheduleslotsjobrequest) String() string {
     o.AttendeeIds = []string{""} 
     o.FacilitatorIds = []string{""} 
    
    
     o.Intervals = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingscheduleslotsjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Coachingscheduleslotsjobrequest

    if CoachingscheduleslotsjobrequestMarshalled {
        return []byte("{}"), nil
    }
    CoachingscheduleslotsjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        AttendeeIds []string `json:"attendeeIds"`
        
        FacilitatorIds []string `json:"facilitatorIds"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Intervals []string `json:"intervals"`
        
        SlotsType string `json:"slotsType"`
        *Alias
    }{

        
        AttendeeIds: []string{""},
        


        
        FacilitatorIds: []string{""},
        


        


        


        
        Intervals: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

