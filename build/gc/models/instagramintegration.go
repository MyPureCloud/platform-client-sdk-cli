package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InstagramintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InstagramintegrationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    Recipient Domainentityref `json:"recipient"`


    


    


    


    


    


    CreateStatus string `json:"createStatus"`


    CreateError Errorbody `json:"createError"`


    SelfUri string `json:"selfUri"`

}

// Instagramintegration
type Instagramintegration struct { 
    


    // Name - The name of the Instagram Integration
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting
    MessagingSetting Messagingsettingreference `json:"messagingSetting"`


    // AppId - The App ID from Facebook
    AppId string `json:"appId"`


    // PageId - The Page ID from Instagram messenger
    PageId string `json:"pageId"`


    // InstagramId - The ID from Instagram messenger
    InstagramId string `json:"instagramId"`


    // InstagramUsername - The Username from Instagram messenger
    InstagramUsername string `json:"instagramUsername"`


    // InstagramName - The name from Instagram messenger
    InstagramName string `json:"instagramName"`


    // InstagramProfileImageUrl - The url of the profile image from Instagram messenger
    InstagramProfileImageUrl string `json:"instagramProfileImageUrl"`


    // Status - The status of the Instagram Integration
    Status string `json:"status"`


    


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
func (o *Instagramintegration) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Instagramintegration) MarshalJSON() ([]byte, error) {
    type Alias Instagramintegration

    if InstagramintegrationMarshalled {
        return []byte("{}"), nil
    }
    InstagramintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingreference `json:"messagingSetting"`
        
        AppId string `json:"appId"`
        
        PageId string `json:"pageId"`
        
        InstagramId string `json:"instagramId"`
        
        InstagramUsername string `json:"instagramUsername"`
        
        InstagramName string `json:"instagramName"`
        
        InstagramProfileImageUrl string `json:"instagramProfileImageUrl"`
        
        Status string `json:"status"`
        
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

