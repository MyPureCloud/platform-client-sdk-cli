package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymentDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    Snippet string `json:"snippet"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    LastModifiedUser Addressableentityref `json:"lastModifiedUser"`


    


    


    SelfUri string `json:"selfUri"`

}

// Webdeployment - Details about a Web Deployment
type Webdeployment struct { 
    


    // Name - The deployment name
    Name string `json:"name"`


    // Description - The description of the config
    Description string `json:"description"`


    // Configuration - The config version this deployment uses
    Configuration Webdeploymentconfigurationversionentityref `json:"configuration"`


    // AllowAllDomains - Property indicates whether all domains are allowed or not. allowedDomains must be empty when this is set as true.
    AllowAllDomains bool `json:"allowAllDomains"`


    // AllowedDomains - The list of domains that are approved to use this deployment; the list will be added to CORS headers for ease of web use.
    AllowedDomains []string `json:"allowedDomains"`


    


    


    


    


    // Flow - A reference to the inboundshortmessage flow used by this deployment
    Flow Domainentityref `json:"flow"`


    // Status - The current status of the deployment
    Status string `json:"status"`


    

}

// String returns a JSON representation of the model
func (o *Webdeployment) String() string {
    
    
    
    
     o.AllowedDomains = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeployment) MarshalJSON() ([]byte, error) {
    type Alias Webdeployment

    if WebdeploymentMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Configuration Webdeploymentconfigurationversionentityref `json:"configuration"`
        
        AllowAllDomains bool `json:"allowAllDomains"`
        
        AllowedDomains []string `json:"allowedDomains"`
        
        Flow Domainentityref `json:"flow"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        


        


        


        
        AllowedDomains: []string{""},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

