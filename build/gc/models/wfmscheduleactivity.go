package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmscheduleactivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmscheduleactivityDud struct { 
    UserReference Userreference `json:"userReference"`


    Activities []Scheduleactivity `json:"activities"`


    FullDayTimeOffMarkers []Fulldaytimeoffmarker `json:"fullDayTimeOffMarkers"`

}

// Wfmscheduleactivity
type Wfmscheduleactivity struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Wfmscheduleactivity) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmscheduleactivity) MarshalJSON() ([]byte, error) {
    type Alias Wfmscheduleactivity

    if WfmscheduleactivityMarshalled {
        return []byte("{}"), nil
    }
    WfmscheduleactivityMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

