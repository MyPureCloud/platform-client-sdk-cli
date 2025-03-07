package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutboundsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutboundsettingsDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    MaxConfigurableCallsPerAgent int `json:"maxConfigurableCallsPerAgent"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Outboundsettings
type Outboundsettings struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // MaxCallsPerAgent - The maximum number of calls that can be placed per agent on any campaign
    MaxCallsPerAgent int `json:"maxCallsPerAgent"`


    // MaxCallsPerAgentDecimal - The maximum number of calls that can be placed per agent on any campaign with decimal precision
    MaxCallsPerAgentDecimal float64 `json:"maxCallsPerAgentDecimal"`


    


    // MaxLineUtilization - The maximum percentage of lines that should be used for Outbound, expressed as a decimal in the range [0.0, 1.0]
    MaxLineUtilization float64 `json:"maxLineUtilization"`


    // AbandonSeconds - The number of seconds used to determine if a call is abandoned
    AbandonSeconds float64 `json:"abandonSeconds"`


    // ComplianceAbandonRateDenominator - The denominator to be used in determining the compliance abandon rate
    ComplianceAbandonRateDenominator string `json:"complianceAbandonRateDenominator"`


    // AutomaticTimeZoneMapping - The settings for automatic time zone mapping. Note that changing these settings will change them for both voice and messaging campaigns.
    AutomaticTimeZoneMapping Automatictimezonemappingsettings `json:"automaticTimeZoneMapping"`


    // RescheduleTimeZoneSkippedContacts - Whether or not to reschedule time-zone blocked contacts
    RescheduleTimeZoneSkippedContacts bool `json:"rescheduleTimeZoneSkippedContacts"`


    

}

// String returns a JSON representation of the model
func (o *Outboundsettings) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outboundsettings) MarshalJSON() ([]byte, error) {
    type Alias Outboundsettings

    if OutboundsettingsMarshalled {
        return []byte("{}"), nil
    }
    OutboundsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        MaxCallsPerAgent int `json:"maxCallsPerAgent"`
        
        MaxCallsPerAgentDecimal float64 `json:"maxCallsPerAgentDecimal"`
        
        MaxLineUtilization float64 `json:"maxLineUtilization"`
        
        AbandonSeconds float64 `json:"abandonSeconds"`
        
        ComplianceAbandonRateDenominator string `json:"complianceAbandonRateDenominator"`
        
        AutomaticTimeZoneMapping Automatictimezonemappingsettings `json:"automaticTimeZoneMapping"`
        
        RescheduleTimeZoneSkippedContacts bool `json:"rescheduleTimeZoneSkippedContacts"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

