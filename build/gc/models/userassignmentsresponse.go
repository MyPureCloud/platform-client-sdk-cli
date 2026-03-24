package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserassignmentsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserassignmentsresponseDud struct { 
    


    

}

// Userassignmentsresponse
type Userassignmentsresponse struct { 
    // User - User assigned to the staffing group.
    User Userreference `json:"user"`


    // Assignments - Date pairs representing the assignments for the user.
    Assignments []Assignmenteffectivedaterange `json:"assignments"`

}

// String returns a JSON representation of the model
func (o *Userassignmentsresponse) String() string {
    
     o.Assignments = []Assignmenteffectivedaterange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userassignmentsresponse) MarshalJSON() ([]byte, error) {
    type Alias Userassignmentsresponse

    if UserassignmentsresponseMarshalled {
        return []byte("{}"), nil
    }
    UserassignmentsresponseMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        Assignments []Assignmenteffectivedaterange `json:"assignments"`
        *Alias
    }{

        


        
        Assignments: []Assignmenteffectivedaterange{{}},
        

        Alias: (*Alias)(u),
    })
}

