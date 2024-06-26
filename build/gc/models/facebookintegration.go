package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacebookintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacebookintegrationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    PageName string `json:"pageName"`


    PageProfileImageUrl string `json:"pageProfileImageUrl"`


    


    Recipient Domainentityref `json:"recipient"`


    


    


    


    


    


    CreateStatus string `json:"createStatus"`


    CreateError Errorbody `json:"createError"`


    SelfUri string `json:"selfUri"`

}

// Facebookintegration
type Facebookintegration struct { 
    


    // Name - The name of the Facebook Integration
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting
    MessagingSetting Messagingsettingreference `json:"messagingSetting"`


    // AppId - The App Id from Facebook messenger
    AppId string `json:"appId"`


    // PageId - The Page Id from Facebook messenger
    PageId string `json:"pageId"`


    


    


    // Status - The status of the Facebook Integration
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
func (o *Facebookintegration) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facebookintegration) MarshalJSON() ([]byte, error) {
    type Alias Facebookintegration

    if FacebookintegrationMarshalled {
        return []byte("{}"), nil
    }
    FacebookintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingreference `json:"messagingSetting"`
        
        AppId string `json:"appId"`
        
        PageId string `json:"pageId"`
        
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

