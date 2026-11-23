package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesttsvoiceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesttsvoiceDud struct { 
    


    


    


    


    


    


    

}

// Capabilitiesttsvoice
type Capabilitiesttsvoice struct { 
    // DisplayName - The display name of the TTS voice
    DisplayName string `json:"displayName"`


    // Gender - The gender of the TTS voice
    Gender string `json:"gender"`


    // Language - The language supported by the TTS voice, as a lowercase language-region tag, e.g. en-us
    Language string `json:"language"`


    // Name - The name of the TTS voice
    Name string `json:"name"`


    // Provider - The provider of the TTS voice. Values are defined by the upstream Speech Integration Service and driven by third-party vendor integrations, so new values may appear without a corresponding API change; treat as an opaque string
    Provider string `json:"provider"`


    // SupportedModels - The models supported by the TTS voice. Model identifiers are vendor-specific and are defined by the upstream Speech Integration Service, so new values may appear without a corresponding API change; treat as opaque strings
    SupportedModels []string `json:"supportedModels"`


    // VoiceType - The type of the TTS voice. Values are defined by the upstream Speech Integration Service, so new values may appear without a corresponding API change; treat as an opaque string
    VoiceType string `json:"voiceType"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesttsvoice) String() string {
    
    
    
    
    
     o.SupportedModels = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesttsvoice) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesttsvoice

    if CapabilitiesttsvoiceMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesttsvoiceMarshalled = true

    return json.Marshal(&struct {
        
        DisplayName string `json:"displayName"`
        
        Gender string `json:"gender"`
        
        Language string `json:"language"`
        
        Name string `json:"name"`
        
        Provider string `json:"provider"`
        
        SupportedModels []string `json:"supportedModels"`
        
        VoiceType string `json:"voiceType"`
        *Alias
    }{

        


        


        


        


        


        
        SupportedModels: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

