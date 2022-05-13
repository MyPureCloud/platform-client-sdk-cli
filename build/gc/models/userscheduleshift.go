package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserscheduleshiftMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserscheduleshiftDud struct { 
    WeekSchedule Weekschedulereference `json:"weekSchedule"`


    


    StartDate time.Time `json:"startDate"`


    LengthInMinutes int `json:"lengthInMinutes"`


    


    


    

}

// Userscheduleshift - Single shift in a user's schedule
type Userscheduleshift struct { 
    


    // Id - ID of the schedule shift. This is only for the case of updating and deleting an existing shift
    Id string `json:"id"`


    


    


    // Activities - List of activities in this shift
    Activities []Userscheduleactivity `json:"activities"`


    // Delete - If marked true for updating this schedule shift, it will be deleted
    Delete bool `json:"delete"`


    // ManuallyEdited - Whether the shift was set as manually edited
    ManuallyEdited bool `json:"manuallyEdited"`

}

// String returns a JSON representation of the model
func (o *Userscheduleshift) String() string {
    
     o.Activities = []Userscheduleactivity{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userscheduleshift) MarshalJSON() ([]byte, error) {
    type Alias Userscheduleshift

    if UserscheduleshiftMarshalled {
        return []byte("{}"), nil
    }
    UserscheduleshiftMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Activities []Userscheduleactivity `json:"activities"`
        
        Delete bool `json:"delete"`
        
        ManuallyEdited bool `json:"manuallyEdited"`
        *Alias
    }{

        


        


        


        


        
        Activities: []Userscheduleactivity{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

