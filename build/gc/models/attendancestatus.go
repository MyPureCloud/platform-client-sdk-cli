package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AttendancestatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AttendancestatusDud struct { 
    DateWorkday time.Time `json:"dateWorkday"`


    AttendanceStatusType string `json:"attendanceStatusType"`

}

// Attendancestatus
type Attendancestatus struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Attendancestatus) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Attendancestatus) MarshalJSON() ([]byte, error) {
    type Alias Attendancestatus

    if AttendancestatusMarshalled {
        return []byte("{}"), nil
    }
    AttendancestatusMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

