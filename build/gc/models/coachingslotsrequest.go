package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingslotsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingslotsrequestDud struct { 
    


    


    


    

}

// Coachingslotsrequest
type Coachingslotsrequest struct { 
    // Interval - Range of time to get slots for scheduling coaching appointments. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // LengthInMinutes - The duration of coaching appointment to schedule in 15 minutes granularity up to maximum of 60 minutes
    LengthInMinutes int `json:"lengthInMinutes"`


    // AttendeeIds - List of attendees to determine coaching appointment slots
    AttendeeIds []string `json:"attendeeIds"`


    // FacilitatorIds - List of facilitators to determine coaching appointment slots
    FacilitatorIds []string `json:"facilitatorIds"`

}

// String returns a JSON representation of the model
func (o *Coachingslotsrequest) String() string {
    
    
    
    
    
    
    
    
    
    
     o.AttendeeIds = []string{""} 
    
    
    
     o.FacilitatorIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingslotsrequest) MarshalJSON() ([]byte, error) {
    type Alias Coachingslotsrequest

    if CoachingslotsrequestMarshalled {
        return []byte("{}"), nil
    }
    CoachingslotsrequestMarshalled = true

    return json.Marshal(&struct { 
        Interval string `json:"interval"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        AttendeeIds []string `json:"attendeeIds"`
        
        FacilitatorIds []string `json:"facilitatorIds"`
        
        *Alias
    }{
        

        

        

        

        

        
        AttendeeIds: []string{""},
        

        

        
        FacilitatorIds: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

