package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ComplianceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ComplianceDud struct { 
    


    


    

}

// Compliance
type Compliance struct { 
    // StopSettings - List of configurations for 'StopSettings' compliance
    StopSettings []Stopsettings `json:"stopSettings"`


    // OptInSettings - List of configurations for 'OptInSettings' compliance
    OptInSettings []Optinsettings `json:"optInSettings"`


    // HelpSettings - List of configurations for 'HelpSettings' compliance
    HelpSettings []Helpsettings `json:"helpSettings"`

}

// String returns a JSON representation of the model
func (o *Compliance) String() string {
     o.StopSettings = []Stopsettings{{}} 
     o.OptInSettings = []Optinsettings{{}} 
     o.HelpSettings = []Helpsettings{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Compliance) MarshalJSON() ([]byte, error) {
    type Alias Compliance

    if ComplianceMarshalled {
        return []byte("{}"), nil
    }
    ComplianceMarshalled = true

    return json.Marshal(&struct {
        
        StopSettings []Stopsettings `json:"stopSettings"`
        
        OptInSettings []Optinsettings `json:"optInSettings"`
        
        HelpSettings []Helpsettings `json:"helpSettings"`
        *Alias
    }{

        
        StopSettings: []Stopsettings{{}},
        


        
        OptInSettings: []Optinsettings{{}},
        


        
        HelpSettings: []Helpsettings{{}},
        

        Alias: (*Alias)(u),
    })
}

