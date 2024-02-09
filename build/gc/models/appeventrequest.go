package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppeventrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppeventrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    

}

// Appeventrequest
type Appeventrequest struct { 
    // EventName - Represents the action the customer performed. Event types are created for each unique event name and can be faceted on in segment and outcome conditions. A valid event name must only contain alphanumeric characters and underscores. A good event name is typically an object followed by the action performed in past tense, e.g. screen_viewed, search_performed, user_registered.
    EventName string `json:"eventName"`


    // ScreenName - The name of the screen, view, or fragment in the app where the event took place.
    ScreenName string `json:"screenName"`


    // App - Application that the customer is interacting with.
    App Journeyapp `json:"app"`


    // Device - Customer's device.
    Device Requestdevice `json:"device"`


    // SdkLibrary - SDK library used to generate the event.
    SdkLibrary Sdklibrary `json:"sdkLibrary"`


    // NetworkConnectivity - Information relating to the device's network connectivity.
    NetworkConnectivity Networkconnectivity `json:"networkConnectivity"`


    // ReferrerUrl - The referrer URL of the first event in the app session.
    ReferrerUrl string `json:"referrerUrl"`


    // SearchQuery - Represents the keywords in a customer search query.
    SearchQuery string `json:"searchQuery"`


    // Attributes - User-defined attributes associated with a particular event. These attributes provide additional context about the event. For example, items_in_cart or subscription_level.
    Attributes map[string]Customeventattribute `json:"attributes"`


    // Traits - Traits are attributes intrinsic to the customer that may be sent in selected events, (e.g. email, lastName, cellPhone). Traits are used to collect information for identity resolution. For example, the same person might be using an application on different devices which might create two sessions with different customerIds. Additional information can be provided as traits to help link those two sessions and customers to a single external contact through common identifiers that were submitted via a form fill, message, or other input in both sessions.
    Traits map[string]Customeventattribute `json:"traits"`


    // CustomerCookieId - Cookie ID of the customer associated with the app event. This is expected to be set per application install or device and can be used to identify a single customer across multiple sessions. This identifier, along with others passed as traits, is used for identity resolution.
    CustomerCookieId string `json:"customerCookieId"`


    // CreatedDate - UTC timestamp indicating when the event actually took place, events older than an hour will be rejected. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`

}

// String returns a JSON representation of the model
func (o *Appeventrequest) String() string {
    
    
    
    
    
    
    
    
     o.Attributes = map[string]Customeventattribute{"": {}} 
     o.Traits = map[string]Customeventattribute{"": {}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appeventrequest) MarshalJSON() ([]byte, error) {
    type Alias Appeventrequest

    if AppeventrequestMarshalled {
        return []byte("{}"), nil
    }
    AppeventrequestMarshalled = true

    return json.Marshal(&struct {
        
        EventName string `json:"eventName"`
        
        ScreenName string `json:"screenName"`
        
        App Journeyapp `json:"app"`
        
        Device Requestdevice `json:"device"`
        
        SdkLibrary Sdklibrary `json:"sdkLibrary"`
        
        NetworkConnectivity Networkconnectivity `json:"networkConnectivity"`
        
        ReferrerUrl string `json:"referrerUrl"`
        
        SearchQuery string `json:"searchQuery"`
        
        Attributes map[string]Customeventattribute `json:"attributes"`
        
        Traits map[string]Customeventattribute `json:"traits"`
        
        CustomerCookieId string `json:"customerCookieId"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        Attributes: map[string]Customeventattribute{"": {}},
        


        
        Traits: map[string]Customeventattribute{"": {}},
        


        


        

        Alias: (*Alias)(u),
    })
}

