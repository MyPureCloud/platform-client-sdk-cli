package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatemuagentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatemuagentrequestDud struct { 
    


    

}

// Updatemuagentrequest
type Updatemuagentrequest struct { 
    // Schedulable - Whether the agent can be included in schedule generation
    Schedulable bool `json:"schedulable"`


    // UserId - User to be updated
    UserId string `json:"userId"`

}

// String returns a JSON representation of the model
func (o *Updatemuagentrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatemuagentrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatemuagentrequest

    if UpdatemuagentrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatemuagentrequestMarshalled = true

    return json.Marshal(&struct {
        
        Schedulable bool `json:"schedulable"`
        
        UserId string `json:"userId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

