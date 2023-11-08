package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DeploymentpingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DeploymentpingDud struct { 
    


    

}

// Deploymentping
type Deploymentping struct { 
    // Actions - Collection of actions to be offered or displayed to the visitor.
    Actions []Deploymentwebaction `json:"actions"`


    // PollIntervalMilliseconds - Custom poll interval in milliseconds; when the return value is -1, disable pings.
    PollIntervalMilliseconds int `json:"pollIntervalMilliseconds"`

}

// String returns a JSON representation of the model
func (o *Deploymentping) String() string {
     o.Actions = []Deploymentwebaction{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Deploymentping) MarshalJSON() ([]byte, error) {
    type Alias Deploymentping

    if DeploymentpingMarshalled {
        return []byte("{}"), nil
    }
    DeploymentpingMarshalled = true

    return json.Marshal(&struct {
        
        Actions []Deploymentwebaction `json:"actions"`
        
        PollIntervalMilliseconds int `json:"pollIntervalMilliseconds"`
        *Alias
    }{

        
        Actions: []Deploymentwebaction{{}},
        


        

        Alias: (*Alias)(u),
    })
}

