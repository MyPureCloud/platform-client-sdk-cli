package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentroadsideassistanceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentroadsideassistanceDud struct { 
    


    


    


    


    

}

// Conversationcontentroadsideassistance - RoadsideAssistance content object.
type Conversationcontentroadsideassistance struct { 
    // Text - The Roadside Assistance message text
    Text string `json:"text"`


    // PhoneNumber - Phone number the user provided
    PhoneNumber string `json:"phoneNumber"`


    // IsDevicePhoneNumber - If the user provided their own phone number
    IsDevicePhoneNumber bool `json:"isDevicePhoneNumber"`


    // Location - User Location object
    Location Conversationcontentlocation `json:"location"`


    // MessageNumber - The counter of the message
    MessageNumber int `json:"messageNumber"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentroadsideassistance) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentroadsideassistance) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentroadsideassistance

    if ConversationcontentroadsideassistanceMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentroadsideassistanceMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        IsDevicePhoneNumber bool `json:"isDevicePhoneNumber"`
        
        Location Conversationcontentlocation `json:"location"`
        
        MessageNumber int `json:"messageNumber"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

