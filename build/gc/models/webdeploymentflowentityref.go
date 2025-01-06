package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymentflowentityrefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymentflowentityrefDud struct { 
    


    


    


    


    

}

// Webdeploymentflowentityref
type Webdeploymentflowentityref struct { 
    // Id - The Flow ID
    Id string `json:"id"`


    // Name - The Flow name
    Name string `json:"name"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // FlowDescription - The flow description for the webdeployment
    FlowDescription string `json:"flowDescription"`


    // PublishVersion - The published config version for the associated deployment
    PublishVersion Flowversion `json:"publishVersion"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentflowentityref) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeploymentflowentityref) MarshalJSON() ([]byte, error) {
    type Alias Webdeploymentflowentityref

    if WebdeploymentflowentityrefMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymentflowentityrefMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        
        FlowDescription string `json:"flowDescription"`
        
        PublishVersion Flowversion `json:"publishVersion"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

