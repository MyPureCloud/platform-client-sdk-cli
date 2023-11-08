package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DeploymentwebactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DeploymentwebactionDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Deploymentwebaction
type Deploymentwebaction struct { 
    // Id - System-generated UUID for the action.
    Id string `json:"id"`


    // MediaType - Action media type used to deliver the action.
    MediaType string `json:"mediaType"`


    // CustomerId - ID string of the customer that the action was triggered for.
    CustomerId string `json:"customerId"`


    // CustomerIdType - Type of the customer ID that the action was triggered for.
    CustomerIdType string `json:"customerIdType"`


    // ActionMapId - ID of the action map that triggered the action.
    ActionMapId string `json:"actionMapId"`


    // ActionMapVersion - Version of the action map that triggered the action.
    ActionMapVersion int `json:"actionMapVersion"`


    // SessionId - ID of the session that the action was triggered for.
    SessionId string `json:"sessionId"`


    // WebMessagingOfferProperties - Web messaging offer specific properties.
    WebMessagingOfferProperties Webmessagingofferproperties `json:"webMessagingOfferProperties"`


    // ContentOfferProperties - Content offer specific properties.
    ContentOfferProperties Contentoffer `json:"contentOfferProperties"`


    // OpenActionProperties - Open action specific properties.
    OpenActionProperties Openactionproperties `json:"openActionProperties"`

}

// String returns a JSON representation of the model
func (o *Deploymentwebaction) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Deploymentwebaction) MarshalJSON() ([]byte, error) {
    type Alias Deploymentwebaction

    if DeploymentwebactionMarshalled {
        return []byte("{}"), nil
    }
    DeploymentwebactionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        MediaType string `json:"mediaType"`
        
        CustomerId string `json:"customerId"`
        
        CustomerIdType string `json:"customerIdType"`
        
        ActionMapId string `json:"actionMapId"`
        
        ActionMapVersion int `json:"actionMapVersion"`
        
        SessionId string `json:"sessionId"`
        
        WebMessagingOfferProperties Webmessagingofferproperties `json:"webMessagingOfferProperties"`
        
        ContentOfferProperties Contentoffer `json:"contentOfferProperties"`
        
        OpenActionProperties Openactionproperties `json:"openActionProperties"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

