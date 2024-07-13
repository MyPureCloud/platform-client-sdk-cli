package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebeventdeviceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebeventdeviceDud struct { 
    


    


    


    


    


    


    


    


    

}

// Webeventdevice
type Webeventdevice struct { 
    // Category - Device category.
    Category string `json:"category"`


    // VarType - Device type (e.g. iPad, iPhone, Other).
    VarType string `json:"type"`


    // IsMobile - Flag that is true for mobile devices.
    IsMobile bool `json:"isMobile"`


    // ScreenHeight - Device's screen height.
    ScreenHeight int `json:"screenHeight"`


    // ScreenWidth - Device's screen width.
    ScreenWidth int `json:"screenWidth"`


    // ScreenDensity - Device's screen density, measured as a scale factor where a value of 1 represents a baseline 1:1 ratio of pixels to logical (device-independent) pixels.
    ScreenDensity int `json:"screenDensity"`


    // OsFamily - Operating system family.
    OsFamily string `json:"osFamily"`


    // OsVersion - Operating system version.
    OsVersion string `json:"osVersion"`


    // Manufacturer - Manufacturer of the device.
    Manufacturer string `json:"manufacturer"`

}

// String returns a JSON representation of the model
func (o *Webeventdevice) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webeventdevice) MarshalJSON() ([]byte, error) {
    type Alias Webeventdevice

    if WebeventdeviceMarshalled {
        return []byte("{}"), nil
    }
    WebeventdeviceMarshalled = true

    return json.Marshal(&struct {
        
        Category string `json:"category"`
        
        VarType string `json:"type"`
        
        IsMobile bool `json:"isMobile"`
        
        ScreenHeight int `json:"screenHeight"`
        
        ScreenWidth int `json:"screenWidth"`
        
        ScreenDensity int `json:"screenDensity"`
        
        OsFamily string `json:"osFamily"`
        
        OsVersion string `json:"osVersion"`
        
        Manufacturer string `json:"manufacturer"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

