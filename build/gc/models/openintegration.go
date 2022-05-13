package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenintegrationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    Recipient Domainentityref `json:"recipient"`


    


    


    


    


    CreateStatus string `json:"createStatus"`


    CreateError Errorbody `json:"createError"`


    SelfUri string `json:"selfUri"`

}

// Openintegration
type Openintegration struct { 
    


    // Name - The name of the Open messaging integration.
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting
    MessagingSetting Messagingsettingreference `json:"messagingSetting"`


    // OutboundNotificationWebhookUrl - The outbound notification webhook URL for the Open messaging integration.
    OutboundNotificationWebhookUrl string `json:"outboundNotificationWebhookUrl"`


    // OutboundNotificationWebhookSignatureSecretToken - The outbound notification webhook signature secret token.
    OutboundNotificationWebhookSignatureSecretToken string `json:"outboundNotificationWebhookSignatureSecretToken"`


    // WebhookHeaders - The user specified headers for the Open messaging integration.
    WebhookHeaders map[string]string `json:"webhookHeaders"`


    // Status - The status of the Open Integration
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
func (o *Openintegration) String() string {
    
    
    
    
    
     o.WebhookHeaders = map[string]string{"": ""} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openintegration) MarshalJSON() ([]byte, error) {
    type Alias Openintegration

    if OpenintegrationMarshalled {
        return []byte("{}"), nil
    }
    OpenintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingreference `json:"messagingSetting"`
        
        OutboundNotificationWebhookUrl string `json:"outboundNotificationWebhookUrl"`
        
        OutboundNotificationWebhookSignatureSecretToken string `json:"outboundNotificationWebhookSignatureSecretToken"`
        
        WebhookHeaders map[string]string `json:"webhookHeaders"`
        
        Status string `json:"status"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        CreatedBy Domainentityref `json:"createdBy"`
        
        ModifiedBy Domainentityref `json:"modifiedBy"`
        *Alias
    }{

        


        


        


        


        


        


        
        WebhookHeaders: map[string]string{"": ""},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

