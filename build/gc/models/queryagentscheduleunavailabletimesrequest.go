package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryagentscheduleunavailabletimesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryagentscheduleunavailabletimesrequestDud struct { 
    


    

}

// Queryagentscheduleunavailabletimesrequest
type Queryagentscheduleunavailabletimesrequest struct { 
    // AgentId - The ID of the agent for whom to fetch unavailable times
    AgentId string `json:"agentId"`


    // Schedule - The schedule for which to fetch unavailable times for the agent
    Schedule Buschedulereference `json:"schedule"`

}

// String returns a JSON representation of the model
func (o *Queryagentscheduleunavailabletimesrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryagentscheduleunavailabletimesrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryagentscheduleunavailabletimesrequest

    if QueryagentscheduleunavailabletimesrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryagentscheduleunavailabletimesrequestMarshalled = true

    return json.Marshal(&struct {
        
        AgentId string `json:"agentId"`
        
        Schedule Buschedulereference `json:"schedule"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

