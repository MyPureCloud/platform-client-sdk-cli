package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserstaffinggroupresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserstaffinggroupresponseDud struct { 
    


    


    


    

}

// Userstaffinggroupresponse
type Userstaffinggroupresponse struct { 
    // StartDate - Effective start date of the user assignment in ISO-8601 format or empty value. Empty value means no limit on start-date.
    StartDate time.Time `json:"startDate"`


    // EndDate - Effective end date of the user assignment in ISO-8601 format or empty value. Empty value means no limit on end-date.
    EndDate time.Time `json:"endDate"`


    // User - The user associated with the staffing group
    User Userreference `json:"user"`


    // StaffingGroup - The staffing group associated with the user
    StaffingGroup Staffinggroupreference `json:"staffingGroup"`

}

// String returns a JSON representation of the model
func (o *Userstaffinggroupresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userstaffinggroupresponse) MarshalJSON() ([]byte, error) {
    type Alias Userstaffinggroupresponse

    if UserstaffinggroupresponseMarshalled {
        return []byte("{}"), nil
    }
    UserstaffinggroupresponseMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        User Userreference `json:"user"`
        
        StaffingGroup Staffinggroupreference `json:"staffingGroup"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

