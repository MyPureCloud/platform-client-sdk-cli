package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummarygenerationconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummarygenerationconfigDud struct { 
    


    


    


    


    

}

// Summarygenerationconfig
type Summarygenerationconfig struct { 
    // Enabled - Copilot generated summary is enabled.
    Enabled bool `json:"enabled"`


    // SummarySetting - Configured summary setting object.
    SummarySetting Summarysettingentity `json:"summarySetting"`


    // RetentionSeconds - Summary retention time in seconds. Can only be modified on the parent assistant.
    RetentionSeconds int `json:"retentionSeconds"`


    // OnDemandSummaryConfig - On-demand summary configuration.
    OnDemandSummaryConfig Ondemandsummaryconfig `json:"onDemandSummaryConfig"`


    // ModelConfig - Model configuration for summarization.
    ModelConfig Modelconfig `json:"modelConfig"`

}

// String returns a JSON representation of the model
func (o *Summarygenerationconfig) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summarygenerationconfig) MarshalJSON() ([]byte, error) {
    type Alias Summarygenerationconfig

    if SummarygenerationconfigMarshalled {
        return []byte("{}"), nil
    }
    SummarygenerationconfigMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        SummarySetting Summarysettingentity `json:"summarySetting"`
        
        RetentionSeconds int `json:"retentionSeconds"`
        
        OnDemandSummaryConfig Ondemandsummaryconfig `json:"onDemandSummaryConfig"`
        
        ModelConfig Modelconfig `json:"modelConfig"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

