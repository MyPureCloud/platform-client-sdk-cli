package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LineintegrationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LineintegrationrequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Lineintegrationrequest
type Lineintegrationrequest struct { 
    


    // Name - The name of the LINE Integration
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // ChannelId - The Channel Id from LINE messenger. New Official LINE account: To create a new official account, LINE requires a Webhook URL. It can be created without specifying Channel Id & Channel Secret. Once the Official account is created by LINE, use the update LINE Integration API to update Channel Id and Channel Secret.  All other accounts: Channel Id is mandatory. (NOTE: ChannelId can only be updated if the integration is set to inactive)
    ChannelId string `json:"channelId"`


    // ChannelSecret - The Channel Secret from LINE messenger. New Official LINE account: To create a new official account, LINE requires a Webhook URL. It can be created without specifying Channel Id & Channel Secret. Once the Official account is created by LINE, use the update LINE Integration API to update Channel Id and Channel Secret.  All other accounts: Channel Secret is mandatory. (NOTE: ChannelSecret can only be updated if the integration is set to inactive)
    ChannelSecret string `json:"channelSecret"`


    // SwitcherSecret - The Switcher Secret from LINE messenger. Some line official accounts are switcher functionality enabled. If the LINE account used for this integration is switcher enabled, then switcher secret is a required field. This secret can be found in your create documentation provided by LINE
    SwitcherSecret string `json:"switcherSecret"`


    // ServiceCode - The Service Code from LINE messenger. Only applicable to LINE Enterprise accounts. This service code can be found in your create documentation provided by LINE
    ServiceCode string `json:"serviceCode"`


    

}

// String returns a JSON representation of the model
func (o *Lineintegrationrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lineintegrationrequest) MarshalJSON() ([]byte, error) {
    type Alias Lineintegrationrequest

    if LineintegrationrequestMarshalled {
        return []byte("{}"), nil
    }
    LineintegrationrequestMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        ChannelId string `json:"channelId"`
        
        ChannelSecret string `json:"channelSecret"`
        
        SwitcherSecret string `json:"switcherSecret"`
        
        ServiceCode string `json:"serviceCode"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

