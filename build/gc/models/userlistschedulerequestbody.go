package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserlistschedulerequestbodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserlistschedulerequestbodyDud struct { 
    


    


    


    

}

// Userlistschedulerequestbody - Request body for fetching the schedule for a group of users over a given time range
type Userlistschedulerequestbody struct { 
    // UserIds - The user ids for which to fetch schedules
    UserIds []string `json:"userIds"`


    // StartDate - Beginning of the range of schedules to fetch, in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - End of the range of schedules to fetch, in ISO-8601 format
    EndDate time.Time `json:"endDate"`


    // LoadFullWeeks - Whether to load the full week's schedule (for the requested users) of any week overlapping the start/end date query parameters, defaults to false
    LoadFullWeeks bool `json:"loadFullWeeks"`

}

// String returns a JSON representation of the model
func (o *Userlistschedulerequestbody) String() string {
    
    
     o.UserIds = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userlistschedulerequestbody) MarshalJSON() ([]byte, error) {
    type Alias Userlistschedulerequestbody

    if UserlistschedulerequestbodyMarshalled {
        return []byte("{}"), nil
    }
    UserlistschedulerequestbodyMarshalled = true

    return json.Marshal(&struct { 
        UserIds []string `json:"userIds"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        LoadFullWeeks bool `json:"loadFullWeeks"`
        
        *Alias
    }{
        

        
        UserIds: []string{""},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

