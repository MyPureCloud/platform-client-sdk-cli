package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GroupsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GroupsettingsDud struct { 
    


    


    


    

}

// Groupsettings
type Groupsettings struct { 
    // MinimumGroupSize - The minimum size of a group for a session
    MinimumGroupSize int `json:"minimumGroupSize"`


    // MaximumGroupSize - The maximum size of a group for a session
    MaximumGroupSize int `json:"maximumGroupSize"`


    // MaximumTotalSessions - The maximum total number of sessions
    MaximumTotalSessions int `json:"maximumTotalSessions"`


    // MaximumConcurrentSessions - The maximum number of sessions that can be scheduled concurrently
    MaximumConcurrentSessions int `json:"maximumConcurrentSessions"`

}

// String returns a JSON representation of the model
func (o *Groupsettings) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Groupsettings) MarshalJSON() ([]byte, error) {
    type Alias Groupsettings

    if GroupsettingsMarshalled {
        return []byte("{}"), nil
    }
    GroupsettingsMarshalled = true

    return json.Marshal(&struct {
        
        MinimumGroupSize int `json:"minimumGroupSize"`
        
        MaximumGroupSize int `json:"maximumGroupSize"`
        
        MaximumTotalSessions int `json:"maximumTotalSessions"`
        
        MaximumConcurrentSessions int `json:"maximumConcurrentSessions"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

