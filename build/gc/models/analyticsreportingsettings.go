package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsreportingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsreportingsettingsDud struct { 
    


    

}

// Analyticsreportingsettings
type Analyticsreportingsettings struct { 
    // PiiMaskingEnabled - Indication of whether or not personal data is masked in data export and the Analytics/Reporting UI
    PiiMaskingEnabled bool `json:"piiMaskingEnabled"`


    // QueueAgentAccessObfuscation - Indication of whether or not to obfuscate export data from the Queue Agent Details view based on User ACL
    QueueAgentAccessObfuscation bool `json:"queueAgentAccessObfuscation"`

}

// String returns a JSON representation of the model
func (o *Analyticsreportingsettings) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsreportingsettings) MarshalJSON() ([]byte, error) {
    type Alias Analyticsreportingsettings

    if AnalyticsreportingsettingsMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsreportingsettingsMarshalled = true

    return json.Marshal(&struct { 
        PiiMaskingEnabled bool `json:"piiMaskingEnabled"`
        
        QueueAgentAccessObfuscation bool `json:"queueAgentAccessObfuscation"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

