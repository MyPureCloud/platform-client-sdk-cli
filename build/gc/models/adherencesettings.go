package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherencesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherencesettingsDud struct { 
    


    


    


    


    


    

}

// Adherencesettings
type Adherencesettings struct { 
    // SevereAlertThresholdMinutes - The threshold in minutes where an alert will be triggered when an agent is considered severely out of adherence
    SevereAlertThresholdMinutes int `json:"severeAlertThresholdMinutes"`


    // AdherenceTargetPercent - Target adherence percentage
    AdherenceTargetPercent int `json:"adherenceTargetPercent"`


    // AdherenceExceptionThresholdSeconds - The threshold in seconds for which agents should not be penalized for being momentarily out of adherence
    AdherenceExceptionThresholdSeconds int `json:"adherenceExceptionThresholdSeconds"`


    // NonOnQueueActivitiesEquivalent - Whether to treat all non-on-queue activities as equivalent for adherence purposes
    NonOnQueueActivitiesEquivalent bool `json:"nonOnQueueActivitiesEquivalent"`


    // TrackOnQueueActivity - Whether to track on-queue activities
    TrackOnQueueActivity bool `json:"trackOnQueueActivity"`


    // IgnoredActivityCategories - Activity categories that should be ignored for adherence purposes
    IgnoredActivityCategories Ignoredactivitycategories `json:"ignoredActivityCategories"`

}

// String returns a JSON representation of the model
func (o *Adherencesettings) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherencesettings) MarshalJSON() ([]byte, error) {
    type Alias Adherencesettings

    if AdherencesettingsMarshalled {
        return []byte("{}"), nil
    }
    AdherencesettingsMarshalled = true

    return json.Marshal(&struct {
        
        SevereAlertThresholdMinutes int `json:"severeAlertThresholdMinutes"`
        
        AdherenceTargetPercent int `json:"adherenceTargetPercent"`
        
        AdherenceExceptionThresholdSeconds int `json:"adherenceExceptionThresholdSeconds"`
        
        NonOnQueueActivitiesEquivalent bool `json:"nonOnQueueActivitiesEquivalent"`
        
        TrackOnQueueActivity bool `json:"trackOnQueueActivity"`
        
        IgnoredActivityCategories Ignoredactivitycategories `json:"ignoredActivityCategories"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

