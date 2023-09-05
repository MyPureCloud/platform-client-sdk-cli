package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExpandablewebdeploymentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExpandablewebdeploymentDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    Snippet string `json:"snippet"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    LastModifiedUser Addressableentityref `json:"lastModifiedUser"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Expandablewebdeployment - Details about a Web Deployment
type Expandablewebdeployment struct { 
    


    // Name - The deployment name
    Name string `json:"name"`


    // Description - The description of the config
    Description string `json:"description"`


    // AllowAllDomains - Property indicates whether all domains are allowed or not. allowedDomains must be empty when this is set as true.
    AllowAllDomains bool `json:"allowAllDomains"`


    // AllowedDomains - The list of domains that are approved to use this deployment; the list will be added to CORS headers for ease of web use.
    AllowedDomains []string `json:"allowedDomains"`


    // SupportedContent - The supported content profile for a deployment
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    


    


    


    


    // Flow - A reference to the inboundshortmessage flow used by this deployment
    Flow Domainentityref `json:"flow"`


    // Status - The current status of the deployment
    Status string `json:"status"`


    // Configuration - The config version this deployment uses
    Configuration Webdeploymentconfigurationversionresponse `json:"configuration"`


    

}

// String returns a JSON representation of the model
func (o *Expandablewebdeployment) String() string {
    
    
    
     o.AllowedDomains = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Expandablewebdeployment) MarshalJSON() ([]byte, error) {
    type Alias Expandablewebdeployment

    if ExpandablewebdeploymentMarshalled {
        return []byte("{}"), nil
    }
    ExpandablewebdeploymentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        AllowAllDomains bool `json:"allowAllDomains"`
        
        AllowedDomains []string `json:"allowedDomains"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        Flow Domainentityref `json:"flow"`
        
        Status string `json:"status"`
        
        Configuration Webdeploymentconfigurationversionresponse `json:"configuration"`
        *Alias
    }{

        


        


        


        


        
        AllowedDomains: []string{""},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

