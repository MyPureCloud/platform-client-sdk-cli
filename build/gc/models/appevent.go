package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppeventDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Appevent
type Appevent struct { 
    // EventName - Represents the action the customer performed. A good event name is typically an object followed by the action performed in past tense (e.g. screen_viewed, order_completed, user_registered).
    EventName string `json:"eventName"`


    // ScreenName - The name of the screen in the app that the event took place.
    ScreenName string `json:"screenName"`


    // App - Application that the customer is interacting with.
    App Journeyapp `json:"app"`


    // Device - Customer's device.
    Device Device `json:"device"`


    // IpAddress - Customer's IP address. May be null if the business configures the tracker to not collect IP addresses.
    IpAddress string `json:"ipAddress"`


    // IpOrganization - Customer's IP-based organization or ISP name.
    IpOrganization string `json:"ipOrganization"`


    // Geolocation - Customer's geolocation.
    Geolocation Journeygeolocation `json:"geolocation"`


    // SdkLibrary - SDK library used to generate the event.
    SdkLibrary Sdklibrary `json:"sdkLibrary"`


    // NetworkConnectivity - Information relating to the device's network connectivity.
    NetworkConnectivity Networkconnectivity `json:"networkConnectivity"`


    // MktCampaign - Marketing / traffic source information.
    MktCampaign Journeycampaign `json:"mktCampaign"`


    // SearchQuery - Represents the keywords in a customer search query.
    SearchQuery string `json:"searchQuery"`


    // Attributes - User-defined attributes associated with a particular event.
    Attributes map[string]Customeventattribute `json:"attributes"`


    // Traits - Traits are attributes intrinsic to the customer that may be sent in selected events. Examples are email, name, phone.
    Traits map[string]Customeventattribute `json:"traits"`

}

// String returns a JSON representation of the model
func (o *Appevent) String() string {
    
    
    
    
    
    
    
    
    
    
    
     o.Attributes = map[string]Customeventattribute{"": {}} 
     o.Traits = map[string]Customeventattribute{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appevent) MarshalJSON() ([]byte, error) {
    type Alias Appevent

    if AppeventMarshalled {
        return []byte("{}"), nil
    }
    AppeventMarshalled = true

    return json.Marshal(&struct {
        
        EventName string `json:"eventName"`
        
        ScreenName string `json:"screenName"`
        
        App Journeyapp `json:"app"`
        
        Device Device `json:"device"`
        
        IpAddress string `json:"ipAddress"`
        
        IpOrganization string `json:"ipOrganization"`
        
        Geolocation Journeygeolocation `json:"geolocation"`
        
        SdkLibrary Sdklibrary `json:"sdkLibrary"`
        
        NetworkConnectivity Networkconnectivity `json:"networkConnectivity"`
        
        MktCampaign Journeycampaign `json:"mktCampaign"`
        
        SearchQuery string `json:"searchQuery"`
        
        Attributes map[string]Customeventattribute `json:"attributes"`
        
        Traits map[string]Customeventattribute `json:"traits"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        
        Attributes: map[string]Customeventattribute{"": {}},
        


        
        Traits: map[string]Customeventattribute{"": {}},
        

        Alias: (*Alias)(u),
    })
}

