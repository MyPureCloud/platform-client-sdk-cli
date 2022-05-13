package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WidgetdeploymentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WidgetdeploymentDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Widgetdeployment
type Widgetdeployment struct { 
    


    // Name
    Name string `json:"name"`


    // Description - A human-readable description of this Deployment.
    Description string `json:"description"`


    // AuthenticationRequired - When true, the customer members starting a chat must be authenticated by supplying their JWT to the create operation.
    AuthenticationRequired bool `json:"authenticationRequired"`


    // Disabled - When true, all create chat operations using this Deployment will be rejected.
    Disabled bool `json:"disabled"`


    // Flow - The URI of the Inbound Chat Flow to run when new chats are initiated under this Deployment.
    Flow Domainentityref `json:"flow"`


    // AllowedDomains - The list of domains that are approved to use this Deployment; the list will be added to CORS headers for ease of web use.
    AllowedDomains []string `json:"allowedDomains"`


    // ClientType - The type of display widget for which this Deployment is configured, which controls the administrator settings shown.
    ClientType string `json:"clientType"`


    // ClientConfig - The client configuration options that should be made available to the clients of this Deployment.
    ClientConfig Widgetclientconfig `json:"clientConfig"`


    

}

// String returns a JSON representation of the model
func (o *Widgetdeployment) String() string {
    
    
    
    
    
     o.AllowedDomains = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Widgetdeployment) MarshalJSON() ([]byte, error) {
    type Alias Widgetdeployment

    if WidgetdeploymentMarshalled {
        return []byte("{}"), nil
    }
    WidgetdeploymentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        AuthenticationRequired bool `json:"authenticationRequired"`
        
        Disabled bool `json:"disabled"`
        
        Flow Domainentityref `json:"flow"`
        
        AllowedDomains []string `json:"allowedDomains"`
        
        ClientType string `json:"clientType"`
        
        ClientConfig Widgetclientconfig `json:"clientConfig"`
        *Alias
    }{

        


        


        


        


        


        


        
        AllowedDomains: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

