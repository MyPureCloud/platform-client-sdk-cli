package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserdeviceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserdeviceDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Userdevice
type Userdevice struct { 
    


    // Name
    Name string `json:"name"`


    // DeviceToken - device token sent by mobile clients.
    DeviceToken string `json:"deviceToken"`


    // NotificationId - notification id of the device.
    NotificationId string `json:"notificationId"`


    // Make - make of the device.
    Make string `json:"make"`


    // Model - Device model
    Model string `json:"model"`


    // AcceptNotifications - if the device accepts notifications
    AcceptNotifications bool `json:"acceptNotifications"`


    // VarType - type of the device; ios or android
    VarType string `json:"type"`


    // SessionHash
    SessionHash string `json:"sessionHash"`


    

}

// String returns a JSON representation of the model
func (o *Userdevice) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userdevice) MarshalJSON() ([]byte, error) {
    type Alias Userdevice

    if UserdeviceMarshalled {
        return []byte("{}"), nil
    }
    UserdeviceMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        DeviceToken string `json:"deviceToken"`
        
        NotificationId string `json:"notificationId"`
        
        Make string `json:"make"`
        
        Model string `json:"model"`
        
        AcceptNotifications bool `json:"acceptNotifications"`
        
        VarType string `json:"type"`
        
        SessionHash string `json:"sessionHash"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

