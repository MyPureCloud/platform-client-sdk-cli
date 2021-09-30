package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhonecapabilitiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhonecapabilitiesDud struct { 
    


    


    


    


    


    


    


    


    

}

// Phonecapabilities
type Phonecapabilities struct { 
    // Provisions
    Provisions bool `json:"provisions"`


    // Registers
    Registers bool `json:"registers"`


    // DualRegisters
    DualRegisters bool `json:"dualRegisters"`


    // HardwareIdType
    HardwareIdType string `json:"hardwareIdType"`


    // AllowReboot
    AllowReboot bool `json:"allowReboot"`


    // NoRebalance
    NoRebalance bool `json:"noRebalance"`


    // NoCloudProvisioning
    NoCloudProvisioning bool `json:"noCloudProvisioning"`


    // MediaCodecs
    MediaCodecs []string `json:"mediaCodecs"`


    // Cdm
    Cdm bool `json:"cdm"`

}

// String returns a JSON representation of the model
func (o *Phonecapabilities) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.MediaCodecs = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phonecapabilities) MarshalJSON() ([]byte, error) {
    type Alias Phonecapabilities

    if PhonecapabilitiesMarshalled {
        return []byte("{}"), nil
    }
    PhonecapabilitiesMarshalled = true

    return json.Marshal(&struct { 
        Provisions bool `json:"provisions"`
        
        Registers bool `json:"registers"`
        
        DualRegisters bool `json:"dualRegisters"`
        
        HardwareIdType string `json:"hardwareIdType"`
        
        AllowReboot bool `json:"allowReboot"`
        
        NoRebalance bool `json:"noRebalance"`
        
        NoCloudProvisioning bool `json:"noCloudProvisioning"`
        
        MediaCodecs []string `json:"mediaCodecs"`
        
        Cdm bool `json:"cdm"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        MediaCodecs: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

