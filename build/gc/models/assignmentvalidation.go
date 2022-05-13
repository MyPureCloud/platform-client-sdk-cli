package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssignmentvalidationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssignmentvalidationDud struct { 
    


    


    


    

}

// Assignmentvalidation
type Assignmentvalidation struct { 
    // MembersNotAssigned - The list of users that are not assigned to any custom performance profile
    MembersNotAssigned []Userreference `json:"membersNotAssigned"`


    // MembersAlreadyAssigned - The list of users that are already assigned to the requesting custom performance profile
    MembersAlreadyAssigned []Userreference `json:"membersAlreadyAssigned"`


    // MembersAlreadyAssignedToOther - The list of users that are already assigned to other custom performance profiles
    MembersAlreadyAssignedToOther []Otherprofileassignment `json:"membersAlreadyAssignedToOther"`


    // InvalidMemberAssignments - The list of user id that are invalid for the gamfication service to handle
    InvalidMemberAssignments []Invalidassignment `json:"invalidMemberAssignments"`

}

// String returns a JSON representation of the model
func (o *Assignmentvalidation) String() string {
     o.MembersNotAssigned = []Userreference{{}} 
     o.MembersAlreadyAssigned = []Userreference{{}} 
     o.MembersAlreadyAssignedToOther = []Otherprofileassignment{{}} 
     o.InvalidMemberAssignments = []Invalidassignment{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assignmentvalidation) MarshalJSON() ([]byte, error) {
    type Alias Assignmentvalidation

    if AssignmentvalidationMarshalled {
        return []byte("{}"), nil
    }
    AssignmentvalidationMarshalled = true

    return json.Marshal(&struct {
        
        MembersNotAssigned []Userreference `json:"membersNotAssigned"`
        
        MembersAlreadyAssigned []Userreference `json:"membersAlreadyAssigned"`
        
        MembersAlreadyAssignedToOther []Otherprofileassignment `json:"membersAlreadyAssignedToOther"`
        
        InvalidMemberAssignments []Invalidassignment `json:"invalidMemberAssignments"`
        *Alias
    }{

        
        MembersNotAssigned: []Userreference{{}},
        


        
        MembersAlreadyAssigned: []Userreference{{}},
        


        
        MembersAlreadyAssignedToOther: []Otherprofileassignment{{}},
        


        
        InvalidMemberAssignments: []Invalidassignment{{}},
        

        Alias: (*Alias)(u),
    })
}

