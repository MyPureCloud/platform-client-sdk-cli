package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GooglebusinessprofileopenintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GooglebusinessprofileopenintegrationDud struct { 
    Id string `json:"id"`


    


    


    


    


    Recipient Domainentityref `json:"recipient"`


    


    


    


    


    CreateStatus string `json:"createStatus"`


    CreateError Errorbody `json:"createError"`


    


    


    SelfUri string `json:"selfUri"`

}

// Googlebusinessprofileopenintegration
type Googlebusinessprofileopenintegration struct { 
    


    // Name - The name of the Google Business Profile Open Integration.
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting
    MessagingSetting Messagingsettingreference `json:"messagingSetting"`


    // Status - The status of the Google Business Profile Open Integration
    Status string `json:"status"`


    


    // DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date this Integration was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // CreatedBy - User reference that created this Integration
    CreatedBy Domainentityref `json:"createdBy"`


    // ModifiedBy - User reference that last modified this Integration
    ModifiedBy Domainentityref `json:"modifiedBy"`


    


    


    // Token - Google OAuth 2 access token reference. The actual token cannot be accessed via Genesys API, only referenced by this property by its ID. When the token is not referenced by any integration, it is deleted after 24 hours.
    Token Googleauthtokenreference `json:"token"`


    // Account - Google Business Profile account accessible with the authorization token
    Account Googlebusinessprofileaccountreference `json:"account"`


    

}

// String returns a JSON representation of the model
func (o *Googlebusinessprofileopenintegration) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Googlebusinessprofileopenintegration) MarshalJSON() ([]byte, error) {
    type Alias Googlebusinessprofileopenintegration

    if GooglebusinessprofileopenintegrationMarshalled {
        return []byte("{}"), nil
    }
    GooglebusinessprofileopenintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingreference `json:"messagingSetting"`
        
        Status string `json:"status"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        CreatedBy Domainentityref `json:"createdBy"`
        
        ModifiedBy Domainentityref `json:"modifiedBy"`
        
        Token Googleauthtokenreference `json:"token"`
        
        Account Googlebusinessprofileaccountreference `json:"account"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

