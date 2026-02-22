package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingintegrationDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    


    


    Status string `json:"status"`


    MessengerType string `json:"messengerType"`


    OpenExtensionType string `json:"openExtensionType"`


    Recipient Domainentityref `json:"recipient"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    CreatedBy Domainentityref `json:"createdBy"`


    ModifiedBy Domainentityref `json:"modifiedBy"`


    Version int `json:"version"`


    SelfUri string `json:"selfUri"`

}

// Messagingintegration
type Messagingintegration struct { 
    


    


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting
    MessagingSetting Messagingsettingreference `json:"messagingSetting"`


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Messagingintegration) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingintegration) MarshalJSON() ([]byte, error) {
    type Alias Messagingintegration

    if MessagingintegrationMarshalled {
        return []byte("{}"), nil
    }
    MessagingintegrationMarshalled = true

    return json.Marshal(&struct {
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingreference `json:"messagingSetting"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

