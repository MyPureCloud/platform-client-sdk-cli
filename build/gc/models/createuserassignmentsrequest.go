package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateuserassignmentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateuserassignmentsrequestDud struct { 
    


    

}

// Createuserassignmentsrequest
type Createuserassignmentsrequest struct { 
    // UserId - The ID of the user assigned to the staffing group.
    UserId string `json:"userId"`


    // Assignments - Assignment effective date ranges for the user.
    Assignments []Assignmenteffectivedaterange `json:"assignments"`

}

// String returns a JSON representation of the model
func (o *Createuserassignmentsrequest) String() string {
    
     o.Assignments = []Assignmenteffectivedaterange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createuserassignmentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Createuserassignmentsrequest

    if CreateuserassignmentsrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateuserassignmentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        Assignments []Assignmenteffectivedaterange `json:"assignments"`
        *Alias
    }{

        


        
        Assignments: []Assignmenteffectivedaterange{{}},
        

        Alias: (*Alias)(u),
    })
}

