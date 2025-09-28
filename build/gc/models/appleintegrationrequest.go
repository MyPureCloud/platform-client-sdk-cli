package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppleintegrationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppleintegrationrequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Appleintegrationrequest
type Appleintegrationrequest struct { 
    


    // Name - The name of the Apple messaging integration.
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting - Defines the message settings to be applied for this integration
    MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`


    // MessagesForBusinessId - The Apple Messages for Business Id for the Apple messaging integration.
    MessagesForBusinessId string `json:"messagesForBusinessId"`


    // BusinessName - The name of the business.
    BusinessName string `json:"businessName"`


    // LogoUrl - The url of the businesses logo.
    LogoUrl string `json:"logoUrl"`


    

}

// String returns a JSON representation of the model
func (o *Appleintegrationrequest) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appleintegrationrequest) MarshalJSON() ([]byte, error) {
    type Alias Appleintegrationrequest

    if AppleintegrationrequestMarshalled {
        return []byte("{}"), nil
    }
    AppleintegrationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`
        
        MessagesForBusinessId string `json:"messagesForBusinessId"`
        
        BusinessName string `json:"businessName"`
        
        LogoUrl string `json:"logoUrl"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

