package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserschedulecontainerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserschedulecontainerDud struct { 
    


    


    

}

// Userschedulecontainer - Container object to hold a map of user schedules
type Userschedulecontainer struct { 
    // ManagementUnitTimeZone - The reference time zone used for the management unit
    ManagementUnitTimeZone string `json:"managementUnitTimeZone"`


    // PublishedSchedules - References to all published week schedules overlapping the start/end date query parameters
    PublishedSchedules []Weekschedulereference `json:"publishedSchedules"`


    // UserSchedules - Map of user id to user schedule
    UserSchedules map[string]Userschedule `json:"userSchedules"`

}

// String returns a JSON representation of the model
func (o *Userschedulecontainer) String() string {
    
    
    
    
    
    
     o.PublishedSchedules = []Weekschedulereference{{}} 
    
    
    
     o.UserSchedules = map[string]Userschedule{"": {}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userschedulecontainer) MarshalJSON() ([]byte, error) {
    type Alias Userschedulecontainer

    if UserschedulecontainerMarshalled {
        return []byte("{}"), nil
    }
    UserschedulecontainerMarshalled = true

    return json.Marshal(&struct { 
        ManagementUnitTimeZone string `json:"managementUnitTimeZone"`
        
        PublishedSchedules []Weekschedulereference `json:"publishedSchedules"`
        
        UserSchedules map[string]Userschedule `json:"userSchedules"`
        
        *Alias
    }{
        

        

        

        
        PublishedSchedules: []Weekschedulereference{{}},
        

        

        
        UserSchedules: map[string]Userschedule{"": {}},
        

        
        Alias: (*Alias)(u),
    })
}

