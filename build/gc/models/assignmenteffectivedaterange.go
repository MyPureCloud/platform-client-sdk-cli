package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssignmenteffectivedaterangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssignmenteffectivedaterangeDud struct { 
    


    

}

// Assignmenteffectivedaterange
type Assignmenteffectivedaterange struct { 
    // StartDate - Effective start date of the user assignment in ISO-8601 format or empty value. Empty value means no limit on start-date.
    StartDate time.Time `json:"startDate"`


    // EndDate - Effective end date of the user assignment in ISO-8601 format or empty value. Empty value means no limit on end-date.
    EndDate time.Time `json:"endDate"`

}

// String returns a JSON representation of the model
func (o *Assignmenteffectivedaterange) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assignmenteffectivedaterange) MarshalJSON() ([]byte, error) {
    type Alias Assignmenteffectivedaterange

    if AssignmenteffectivedaterangeMarshalled {
        return []byte("{}"), nil
    }
    AssignmenteffectivedaterangeMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

