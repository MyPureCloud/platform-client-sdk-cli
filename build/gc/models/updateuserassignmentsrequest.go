package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateuserassignmentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateuserassignmentsrequestDud struct { 
    


    

}

// Updateuserassignmentsrequest
type Updateuserassignmentsrequest struct { 
    // UserId - The ID of the user assigned to the staffing group.
    UserId string `json:"userId"`


    // Assignments - Assignment effective date ranges for the user. Empty list removes all assignments.
    Assignments Setwrapperassignmenteffectivedaterange `json:"assignments"`

}

// String returns a JSON representation of the model
func (o *Updateuserassignmentsrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateuserassignmentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateuserassignmentsrequest

    if UpdateuserassignmentsrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateuserassignmentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        Assignments Setwrapperassignmenteffectivedaterange `json:"assignments"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

