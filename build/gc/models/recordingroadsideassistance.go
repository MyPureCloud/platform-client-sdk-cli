package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingroadsideassistanceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingroadsideassistanceDud struct { 
    


    


    


    


    

}

// Recordingroadsideassistance
type Recordingroadsideassistance struct { 
    // Text - The Roadside Assistance message text.
    Text string `json:"text"`


    // PhoneNumber - Phone number the user provided.
    PhoneNumber string `json:"phoneNumber"`


    // IsDevicePhoneNumber - If the user provided their own phone number.
    IsDevicePhoneNumber bool `json:"isDevicePhoneNumber"`


    // Location - User Location object.
    Location Recordinglocation `json:"location"`


    // MessageNumber - The counter of the message.
    MessageNumber int `json:"messageNumber"`

}

// String returns a JSON representation of the model
func (o *Recordingroadsideassistance) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingroadsideassistance) MarshalJSON() ([]byte, error) {
    type Alias Recordingroadsideassistance

    if RecordingroadsideassistanceMarshalled {
        return []byte("{}"), nil
    }
    RecordingroadsideassistanceMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        IsDevicePhoneNumber bool `json:"isDevicePhoneNumber"`
        
        Location Recordinglocation `json:"location"`
        
        MessageNumber int `json:"messageNumber"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

