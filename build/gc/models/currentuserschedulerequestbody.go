package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CurrentuserschedulerequestbodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CurrentuserschedulerequestbodyDud struct { 
    


    


    

}

// Currentuserschedulerequestbody - POST request body for fetching the current user's schedule over a given range
type Currentuserschedulerequestbody struct { 
    // StartDate - Beginning of the range of schedules to fetch, in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - End of the range of schedules to fetch, in ISO-8601 format
    EndDate time.Time `json:"endDate"`


    // LoadFullWeeks - Whether to load the full week's schedule (for the current user) of any week overlapping the start/end date query parameters, defaults to false
    LoadFullWeeks bool `json:"loadFullWeeks"`

}

// String returns a JSON representation of the model
func (o *Currentuserschedulerequestbody) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Currentuserschedulerequestbody) MarshalJSON() ([]byte, error) {
    type Alias Currentuserschedulerequestbody

    if CurrentuserschedulerequestbodyMarshalled {
        return []byte("{}"), nil
    }
    CurrentuserschedulerequestbodyMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        LoadFullWeeks bool `json:"loadFullWeeks"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

