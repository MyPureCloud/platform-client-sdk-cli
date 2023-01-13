package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymentsrefreshjwtrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymentsrefreshjwtrequestDud struct { 
    


    

}

// Webdeploymentsrefreshjwtrequest
type Webdeploymentsrefreshjwtrequest struct { 
    // RefreshToken - Refresh token used to issue a new JWT.
    RefreshToken string `json:"refreshToken"`


    // DeploymentId - The WebDeployment ID
    DeploymentId string `json:"deploymentId"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentsrefreshjwtrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeploymentsrefreshjwtrequest) MarshalJSON() ([]byte, error) {
    type Alias Webdeploymentsrefreshjwtrequest

    if WebdeploymentsrefreshjwtrequestMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymentsrefreshjwtrequestMarshalled = true

    return json.Marshal(&struct {
        
        RefreshToken string `json:"refreshToken"`
        
        DeploymentId string `json:"deploymentId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

