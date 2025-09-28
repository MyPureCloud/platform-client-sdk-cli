package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppleintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppleintegrationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    Recipient Domainentityref `json:"recipient"`


    


    


    


    


    CreateStatus string `json:"createStatus"`


    CreateError Errorbody `json:"createError"`


    IdentityResolution Identityresolutionconfig `json:"identityResolution"`


    SelfUri string `json:"selfUri"`

}

// Appleintegration
type Appleintegration struct { 
    


    // Name - The name of the Apple messaging integration.
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting
    MessagingSetting Messagingsettingreference `json:"messagingSetting"`


    // MessagesForBusinessId - The Apple Messages for Business Id for the Apple messaging integration.
    MessagesForBusinessId string `json:"messagesForBusinessId"`


    // BusinessName - The name of the business.
    BusinessName string `json:"businessName"`


    // LogoUrl - The url of the businesses logo.
    LogoUrl string `json:"logoUrl"`


    // Status - The status of the Apple Integration
    Status string `json:"status"`


    


    // DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date this Integration was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // CreatedBy - User reference that created this Integration
    CreatedBy Domainentityref `json:"createdBy"`


    // ModifiedBy - User reference that last modified this Integration
    ModifiedBy Domainentityref `json:"modifiedBy"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Appleintegration) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appleintegration) MarshalJSON() ([]byte, error) {
    type Alias Appleintegration

    if AppleintegrationMarshalled {
        return []byte("{}"), nil
    }
    AppleintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingreference `json:"messagingSetting"`
        
        MessagesForBusinessId string `json:"messagesForBusinessId"`
        
        BusinessName string `json:"businessName"`
        
        LogoUrl string `json:"logoUrl"`
        
        Status string `json:"status"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        CreatedBy Domainentityref `json:"createdBy"`
        
        ModifiedBy Domainentityref `json:"modifiedBy"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

