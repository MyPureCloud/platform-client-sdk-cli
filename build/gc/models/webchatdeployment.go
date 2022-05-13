package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchatdeploymentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchatdeploymentDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Webchatdeployment
type Webchatdeployment struct { 
    


    // Name
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // AuthenticationRequired
    AuthenticationRequired bool `json:"authenticationRequired"`


    // AuthenticationUrl - URL for third party service authenticating web chat clients. See https://github.com/MyPureCloud/authenticated-web-chat-server-examples
    AuthenticationUrl string `json:"authenticationUrl"`


    // Disabled
    Disabled bool `json:"disabled"`


    // WebChatConfig
    WebChatConfig Webchatconfig `json:"webChatConfig"`


    // AllowedDomains
    AllowedDomains []string `json:"allowedDomains"`


    // Flow - The URI of the Inbound Chat Flow to run when new chats are initiated under this Deployment.
    Flow Domainentityref `json:"flow"`


    

}

// String returns a JSON representation of the model
func (o *Webchatdeployment) String() string {
    
    
    
    
    
    
     o.AllowedDomains = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchatdeployment) MarshalJSON() ([]byte, error) {
    type Alias Webchatdeployment

    if WebchatdeploymentMarshalled {
        return []byte("{}"), nil
    }
    WebchatdeploymentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        AuthenticationRequired bool `json:"authenticationRequired"`
        
        AuthenticationUrl string `json:"authenticationUrl"`
        
        Disabled bool `json:"disabled"`
        
        WebChatConfig Webchatconfig `json:"webChatConfig"`
        
        AllowedDomains []string `json:"allowedDomains"`
        
        Flow Domainentityref `json:"flow"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        AllowedDomains: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

