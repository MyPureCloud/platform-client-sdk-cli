package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesnlusupportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesnlusupportDud struct { 
    


    


    


    

}

// Capabilitiesnlusupport
type Capabilitiesnlusupport struct { 
    // SupportedLanguages
    SupportedLanguages []Capabilitiesnlusupportedlanguage `json:"supportedLanguages"`


    // Limits
    Limits map[string]int `json:"limits"`


    // NluEngines
    NluEngines map[string]Capabilitiesnlusettings `json:"nluEngines"`


    // FlowTypesToEngines
    FlowTypesToEngines map[string]interface{} `json:"flowTypesToEngines"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesnlusupport) String() string {
     o.SupportedLanguages = []Capabilitiesnlusupportedlanguage{{}} 
     o.Limits = map[string]int{"": 0} 
     o.NluEngines = map[string]Capabilitiesnlusettings{"": {}} 
     o.FlowTypesToEngines = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesnlusupport) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesnlusupport

    if CapabilitiesnlusupportMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesnlusupportMarshalled = true

    return json.Marshal(&struct {
        
        SupportedLanguages []Capabilitiesnlusupportedlanguage `json:"supportedLanguages"`
        
        Limits map[string]int `json:"limits"`
        
        NluEngines map[string]Capabilitiesnlusettings `json:"nluEngines"`
        
        FlowTypesToEngines map[string]interface{} `json:"flowTypesToEngines"`
        *Alias
    }{

        
        SupportedLanguages: []Capabilitiesnlusupportedlanguage{{}},
        


        
        Limits: map[string]int{"": 0},
        


        
        NluEngines: map[string]Capabilitiesnlusettings{"": {}},
        


        
        FlowTypesToEngines: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

