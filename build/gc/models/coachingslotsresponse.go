package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingslotsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingslotsresponseDud struct { 
    SuggestedSlots []Coachingslot `json:"suggestedSlots"`


    AttendeeSchedules []Useravailabletimes `json:"attendeeSchedules"`


    FacilitatorSchedules []Useravailabletimes `json:"facilitatorSchedules"`

}

// Coachingslotsresponse
type Coachingslotsresponse struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Coachingslotsresponse) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingslotsresponse) MarshalJSON() ([]byte, error) {
    type Alias Coachingslotsresponse

    if CoachingslotsresponseMarshalled {
        return []byte("{}"), nil
    }
    CoachingslotsresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

