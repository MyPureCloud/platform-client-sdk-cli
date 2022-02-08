package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappintegrationDud struct { 
    Id string `json:"id"`


    


    


    


    


    Recipient Domainentityref `json:"recipient"`


    


    


    


    


    


    ActivationStatusCode string `json:"activationStatusCode"`


    ActivationErrorInfo Errorbody `json:"activationErrorInfo"`


    CreateStatus string `json:"createStatus"`


    CreateError Errorbody `json:"createError"`


    SelfUri string `json:"selfUri"`

}

// Whatsappintegration
type Whatsappintegration struct { 
    


    // Name - The name of the WhatsApp integration.
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // PhoneNumber - The phone number associated to the whatsApp integration.
    PhoneNumber string `json:"phoneNumber"`


    // Status - The status of the WhatsApp Integration
    Status string `json:"status"`


    


    // DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date this Integration was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // CreatedBy - User reference that created this Integration
    CreatedBy Domainentityref `json:"createdBy"`


    // ModifiedBy - User reference that last modified this Integration
    ModifiedBy Domainentityref `json:"modifiedBy"`


    // Version - Version number required for updates.
    Version int `json:"version"`


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Whatsappintegration) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappintegration) MarshalJSON() ([]byte, error) {
    type Alias Whatsappintegration

    if WhatsappintegrationMarshalled {
        return []byte("{}"), nil
    }
    WhatsappintegrationMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        PhoneNumber string `json:"phoneNumber"`
        
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

