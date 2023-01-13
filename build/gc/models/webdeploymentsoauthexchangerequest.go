package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymentsoauthexchangerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymentsoauthexchangerequestDud struct { 
    


    


    

}

// Webdeploymentsoauthexchangerequest
type Webdeploymentsoauthexchangerequest struct { 
    // DeploymentId - The WebDeployment ID
    DeploymentId string `json:"deploymentId"`


    // JourneyContext - A Customer journey context.
    JourneyContext Webdeploymentsjourneycontext `json:"journeyContext"`


    // Oauth
    Oauth Webdeploymentsoauthrequestparameters `json:"oauth"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentsoauthexchangerequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeploymentsoauthexchangerequest) MarshalJSON() ([]byte, error) {
    type Alias Webdeploymentsoauthexchangerequest

    if WebdeploymentsoauthexchangerequestMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymentsoauthexchangerequestMarshalled = true

    return json.Marshal(&struct {
        
        DeploymentId string `json:"deploymentId"`
        
        JourneyContext Webdeploymentsjourneycontext `json:"journeyContext"`
        
        Oauth Webdeploymentsoauthrequestparameters `json:"oauth"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

