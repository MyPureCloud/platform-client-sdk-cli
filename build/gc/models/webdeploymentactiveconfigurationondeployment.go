package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymentactiveconfigurationondeploymentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymentactiveconfigurationondeploymentDud struct { 
    


    

}

// Webdeploymentactiveconfigurationondeployment - Details about the active configuration on a deployment
type Webdeploymentactiveconfigurationondeployment struct { 
    // ConfigurationVersion - The active configuration on a deployment
    ConfigurationVersion Webdeploymentconfigurationversion `json:"configurationVersion"`


    // Deployment - The web deployment associated with the active configuration
    Deployment Webdeployment `json:"deployment"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentactiveconfigurationondeployment) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeploymentactiveconfigurationondeployment) MarshalJSON() ([]byte, error) {
    type Alias Webdeploymentactiveconfigurationondeployment

    if WebdeploymentactiveconfigurationondeploymentMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymentactiveconfigurationondeploymentMarshalled = true

    return json.Marshal(&struct {
        
        ConfigurationVersion Webdeploymentconfigurationversion `json:"configurationVersion"`
        
        Deployment Webdeployment `json:"deployment"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

