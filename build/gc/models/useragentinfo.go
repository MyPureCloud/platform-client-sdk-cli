package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseragentinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseragentinfoDud struct { 
    


    


    

}

// Useragentinfo
type Useragentinfo struct { 
    // FirmwareVersion - The firmware version of the phone.
    FirmwareVersion string `json:"firmwareVersion"`


    // Manufacturer - The manufacturer of the phone.
    Manufacturer string `json:"manufacturer"`


    // Model - The model of the phone.
    Model string `json:"model"`

}

// String returns a JSON representation of the model
func (o *Useragentinfo) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useragentinfo) MarshalJSON() ([]byte, error) {
    type Alias Useragentinfo

    if UseragentinfoMarshalled {
        return []byte("{}"), nil
    }
    UseragentinfoMarshalled = true

    return json.Marshal(&struct { 
        FirmwareVersion string `json:"firmwareVersion"`
        
        Manufacturer string `json:"manufacturer"`
        
        Model string `json:"model"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

