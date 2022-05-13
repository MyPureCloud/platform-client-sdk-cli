package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssignmentDud struct { 
    


    


    

}

// Assignment
type Assignment struct { 
    // AssignedMembers - The list of users successfully assigned to the custom performance profile
    AssignedMembers []Userreference `json:"assignedMembers"`


    // RemovedMembers - The list of users successfully removed from the custom performance profile
    RemovedMembers []Userreference `json:"removedMembers"`


    // AssignmentErrors - The list of users failed assignment or removal for the custom performance profile
    AssignmentErrors []Assignmenterror `json:"assignmentErrors"`

}

// String returns a JSON representation of the model
func (o *Assignment) String() string {
     o.AssignedMembers = []Userreference{{}} 
     o.RemovedMembers = []Userreference{{}} 
     o.AssignmentErrors = []Assignmenterror{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assignment) MarshalJSON() ([]byte, error) {
    type Alias Assignment

    if AssignmentMarshalled {
        return []byte("{}"), nil
    }
    AssignmentMarshalled = true

    return json.Marshal(&struct {
        
        AssignedMembers []Userreference `json:"assignedMembers"`
        
        RemovedMembers []Userreference `json:"removedMembers"`
        
        AssignmentErrors []Assignmenterror `json:"assignmentErrors"`
        *Alias
    }{

        
        AssignedMembers: []Userreference{{}},
        


        
        RemovedMembers: []Userreference{{}},
        


        
        AssignmentErrors: []Assignmenterror{{}},
        

        Alias: (*Alias)(u),
    })
}

