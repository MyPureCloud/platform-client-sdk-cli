package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LineintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LineintegrationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    Recipient Domainentityref `json:"recipient"`


    


    


    


    


    


    CreateStatus string `json:"createStatus"`


    CreateError Errorbody `json:"createError"`


    SelfUri string `json:"selfUri"`

}

// Lineintegration
type Lineintegration struct { 
    


    // Name - The name of the LINE Integration
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting
    MessagingSetting Messagingsettingreference `json:"messagingSetting"`


    // ChannelId - The Channel Id from LINE messenger
    ChannelId string `json:"channelId"`


    // WebhookUri - The Webhook URI to be updated in LINE platform
    WebhookUri string `json:"webhookUri"`


    // Status - The status of the LINE Integration
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
func (o *Lineintegration) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lineintegration) MarshalJSON() ([]byte, error) {
    type Alias Lineintegration

    if LineintegrationMarshalled {
        return []byte("{}"), nil
    }
    LineintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingreference `json:"messagingSetting"`
        
        ChannelId string `json:"channelId"`
        
        WebhookUri string `json:"webhookUri"`
        
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

