package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AttendancestatuslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AttendancestatuslistingDud struct { 
    

}

// Attendancestatuslisting
type Attendancestatuslisting struct { 
    // Entities
    Entities []Attendancestatus `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Attendancestatuslisting) String() string {
    
    
     o.Entities = []Attendancestatus{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Attendancestatuslisting) MarshalJSON() ([]byte, error) {
    type Alias Attendancestatuslisting

    if AttendancestatuslistingMarshalled {
        return []byte("{}"), nil
    }
    AttendancestatuslistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Attendancestatus `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Attendancestatus{{}},
        

        
        Alias: (*Alias)(u),
    })
}

