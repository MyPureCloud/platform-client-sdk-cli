package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitterintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitterintegrationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    Recipient Domainentityref `json:"recipient"`


    


    


    


    


    


    CreateStatus string `json:"createStatus"`


    CreateError Errorbody `json:"createError"`


    SelfUri string `json:"selfUri"`

}

// Twitterintegration
type Twitterintegration struct { 
    


    // Name - The name of the Twitter Integration
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting
    MessagingSetting Messagingsettingreference `json:"messagingSetting"`


    // AccessTokenKey - The Access Token Key from Twitter messenger
    AccessTokenKey string `json:"accessTokenKey"`


    // ConsumerKey - The Consumer Key from Twitter messenger
    ConsumerKey string `json:"consumerKey"`


    // Username - The Username from Twitter
    Username string `json:"username"`


    // UserId - The UserId from Twitter
    UserId string `json:"userId"`


    // Status - The status of the Twitter Integration
    Status string `json:"status"`


    // Tier - The type of twitter account to be used for the integration
    Tier string `json:"tier"`


    // EnvName - The Twitter environment name, e.g.: env-beta (required for premium tier)
    EnvName string `json:"envName"`


    


    // DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date this Integration was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // CreatedBy - User reference that created this Integration
    CreatedBy Domainentityref `json:"createdBy"`


    // ModifiedBy - User reference that last modified this Integration
    ModifiedBy Domainentityref `json:"modifiedBy"`


    // Version - Version number required for updates.
    Version int `json:"version"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Twitterintegration) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitterintegration) MarshalJSON() ([]byte, error) {
    type Alias Twitterintegration

    if TwitterintegrationMarshalled {
        return []byte("{}"), nil
    }
    TwitterintegrationMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingreference `json:"messagingSetting"`
        
        AccessTokenKey string `json:"accessTokenKey"`
        
        ConsumerKey string `json:"consumerKey"`
        
        Username string `json:"username"`
        
        UserId string `json:"userId"`
        
        Status string `json:"status"`
        
        Tier string `json:"tier"`
        
        EnvName string `json:"envName"`
        
        
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        CreatedBy Domainentityref `json:"createdBy"`
        
        ModifiedBy Domainentityref `json:"modifiedBy"`
        
        Version int `json:"version"`
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

