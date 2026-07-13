package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateactivityplanrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateactivityplanrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Updateactivityplanrequest
type Updateactivityplanrequest struct { 
    // Name - The name of the activity plan
    Name string `json:"name"`


    // Description - The description of the activity plan
    Description string `json:"description"`


    // GroupSettings - Group settings for the activity plan
    GroupSettings Valuewrappergroupsettings `json:"groupSettings"`


    // AttendeesSearchRule - Attendee search rule for this activity plan
    AttendeesSearchRule Valuewrapperusersearchrule `json:"attendeesSearchRule"`


    // FacilitatorsSearchRule - Facilitator search rule for this activity plan
    FacilitatorsSearchRule Valuewrapperusersearchrule `json:"facilitatorsSearchRule"`


    // TransitionTimeMinutes - Transition time in minutes between facilitated sessions
    TransitionTimeMinutes int `json:"transitionTimeMinutes"`


    // ServiceGoalImpactOverrides - Allowable service goal impact override settings for this activity plan
    ServiceGoalImpactOverrides Valuewrapperactivityplanservicegoalimpactoverrides `json:"serviceGoalImpactOverrides"`


    // OptimizationObjective - The optimization objective of this activity plan
    OptimizationObjective string `json:"optimizationObjective"`


    // State - The state of this activity plan
    State string `json:"state"`


    // FixedAvailability - Fixed availability configuration for the activity plan
    FixedAvailability Listwrapperfixedavailability `json:"fixedAvailability"`


    // ActivityCodeId - The activity code associated with the activity plan
    ActivityCodeId string `json:"activityCodeId"`


    // VarType - The type of the activity plan
    VarType string `json:"type"`


    // LengthMinutes - The length in minutes of the activity plan
    LengthMinutes int `json:"lengthMinutes"`


    // CountsAsPaidTime - Whether the activity should count as paid time
    CountsAsPaidTime bool `json:"countsAsPaidTime"`


    // RecurrenceSettings - Settings controlling recurrence for the activity plan. If not set the activity plan will only occur once
    RecurrenceSettings Recurrencesettingsbase `json:"recurrenceSettings"`

}

// String returns a JSON representation of the model
func (o *Updateactivityplanrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateactivityplanrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateactivityplanrequest

    if UpdateactivityplanrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateactivityplanrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        GroupSettings Valuewrappergroupsettings `json:"groupSettings"`
        
        AttendeesSearchRule Valuewrapperusersearchrule `json:"attendeesSearchRule"`
        
        FacilitatorsSearchRule Valuewrapperusersearchrule `json:"facilitatorsSearchRule"`
        
        TransitionTimeMinutes int `json:"transitionTimeMinutes"`
        
        ServiceGoalImpactOverrides Valuewrapperactivityplanservicegoalimpactoverrides `json:"serviceGoalImpactOverrides"`
        
        OptimizationObjective string `json:"optimizationObjective"`
        
        State string `json:"state"`
        
        FixedAvailability Listwrapperfixedavailability `json:"fixedAvailability"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        VarType string `json:"type"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        CountsAsPaidTime bool `json:"countsAsPaidTime"`
        
        RecurrenceSettings Recurrencesettingsbase `json:"recurrenceSettings"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

