package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymentconfigurationversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymentconfigurationversionDud struct { 
    Id string `json:"id"`


    


    Version string `json:"version"`


    


    


    


    


    


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    DatePublished time.Time `json:"datePublished"`


    LastModifiedUser Addressableentityref `json:"lastModifiedUser"`


    CreatedUser Addressableentityref `json:"createdUser"`


    PublishedUser Addressableentityref `json:"publishedUser"`


    


    SelfUri string `json:"selfUri"`

}

// Webdeploymentconfigurationversion - Details about the configuration version of a Web Deployment
type Webdeploymentconfigurationversion struct { 
    


    // Name - The configuration version name
    Name string `json:"name"`


    


    // Description - The description of the configuration
    Description string `json:"description"`


    // Languages - A list of languages supported on the configuration
    Languages []string `json:"languages"`


    // DefaultLanguage - The default language to use for the configuration
    DefaultLanguage string `json:"defaultLanguage"`


    // Messenger - The settings for messenger
    Messenger Messengersettings `json:"messenger"`


    // Position - The settings for position
    Position Positionsettings `json:"position"`


    // SupportCenter - The settings for support center
    SupportCenter Supportcentersettings `json:"supportCenter"`


    // Cobrowse - The settings for cobrowse
    Cobrowse Cobrowsesettings `json:"cobrowse"`


    // JourneyEvents - The settings for journey events
    JourneyEvents Journeyeventssettings `json:"journeyEvents"`


    // AuthenticationSettings - The settings for authenticated deployments
    AuthenticationSettings Authenticationsettings `json:"authenticationSettings"`


    


    


    


    


    


    


    // Status - The current status of the configuration version
    Status string `json:"status"`


    

}

// String returns a JSON representation of the model
func (o *Webdeploymentconfigurationversion) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Languages = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeploymentconfigurationversion) MarshalJSON() ([]byte, error) {
    type Alias Webdeploymentconfigurationversion

    if WebdeploymentconfigurationversionMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymentconfigurationversionMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        Description string `json:"description"`
        
        Languages []string `json:"languages"`
        
        DefaultLanguage string `json:"defaultLanguage"`
        
        Messenger Messengersettings `json:"messenger"`
        
        Position Positionsettings `json:"position"`
        
        SupportCenter Supportcentersettings `json:"supportCenter"`
        
        Cobrowse Cobrowsesettings `json:"cobrowse"`
        
        JourneyEvents Journeyeventssettings `json:"journeyEvents"`
        
        AuthenticationSettings Authenticationsettings `json:"authenticationSettings"`
        
        
        
        
        
        
        
        
        
        
        
        
        
        Status string `json:"status"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Languages: []string{""},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

