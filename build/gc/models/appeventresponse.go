package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppeventresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppeventresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Appeventresponse
type Appeventresponse struct { 
    // Id - System-generated UUID for the event.
    Id string `json:"id"`


    // CustomerId - Identifier of the customer in the source of the event.
    CustomerId string `json:"customerId"`


    // CustomerIdType - Type of identifier for the customer ID (cookie, email etc.).
    CustomerIdType string `json:"customerIdType"`


    // EventName - Represents the action the customer performed. A good event name is typically an object followed by the action performed in past tense (e.g. screen_viewed, order_completed, user_registered).
    EventName string `json:"eventName"`


    // ScreenName - The name of the screen in the app that the event took place.
    ScreenName string `json:"screenName"`


    // App - Application that the customer is interacting with.
    App Journeyapp `json:"app"`


    // Device - Customer's device.
    Device Device `json:"device"`


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


    // Session - The app session the event belongs to.
    Session Appeventresponsesession `json:"session"`


    // SearchQuery - Represents the keywords in a customer search query.
    SearchQuery string `json:"searchQuery"`


    // Attributes - User-defined attributes associated with a particular event.
    Attributes map[string]Customeventattribute `json:"attributes"`


    // Traits - Traits are attributes intrinsic to the customer that may be sent in selected events (e.g. email, givenName, cellPhone).
    Traits map[string]Customeventattribute `json:"traits"`


    // CreatedDate - UTC timestamp indicating when the event actually took place. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`

}

// String returns a JSON representation of the model
func (o *Appeventresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Attributes = map[string]Customeventattribute{"": {}} 
     o.Traits = map[string]Customeventattribute{"": {}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appeventresponse) MarshalJSON() ([]byte, error) {
    type Alias Appeventresponse

    if AppeventresponseMarshalled {
        return []byte("{}"), nil
    }
    AppeventresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        CustomerId string `json:"customerId"`
        
        CustomerIdType string `json:"customerIdType"`
        
        EventName string `json:"eventName"`
        
        ScreenName string `json:"screenName"`
        
        App Journeyapp `json:"app"`
        
        Device Device `json:"device"`
        
        IpOrganization string `json:"ipOrganization"`
        
        Geolocation Journeygeolocation `json:"geolocation"`
        
        SdkLibrary Sdklibrary `json:"sdkLibrary"`
        
        NetworkConnectivity Networkconnectivity `json:"networkConnectivity"`
        
        MktCampaign Journeycampaign `json:"mktCampaign"`
        
        Session Appeventresponsesession `json:"session"`
        
        SearchQuery string `json:"searchQuery"`
        
        Attributes map[string]Customeventattribute `json:"attributes"`
        
        Traits map[string]Customeventattribute `json:"traits"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Attributes: map[string]Customeventattribute{"": {}},
        


        
        Traits: map[string]Customeventattribute{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

