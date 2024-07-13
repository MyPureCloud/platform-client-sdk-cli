package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateactivityplanrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateactivityplanrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Createactivityplanrequest
type Createactivityplanrequest struct { 
    // Name - The name of the activity plan
    Name string `json:"name"`


    // ManagementUnitIds - The management units to which this activity plan applies. Empty list or null means this activity plan applies to the entire business unit
    ManagementUnitIds []string `json:"managementUnitIds"`


    // Description - The description of the activity plan
    Description string `json:"description"`


    // ActivityCodeId - The activity code associated with the activity plan
    ActivityCodeId string `json:"activityCodeId"`


    // VarType - The type of the activity plan
    VarType string `json:"type"`


    // LengthMinutes - The length in minutes of the activity plan
    LengthMinutes int `json:"lengthMinutes"`


    // InitialSchedulePeriod - The initial scheduling period for the activity plan
    InitialSchedulePeriod Schedulingperiod `json:"initialSchedulePeriod"`


    // GroupSettings - Group settings for the activity plan
    GroupSettings Groupsettings `json:"groupSettings"`


    // RecurrenceSettings - Settings controlling recurrence for the activity plan. If not set the activity plan will only occur once
    RecurrenceSettings Recurrencesettings `json:"recurrenceSettings"`


    // AttendeesSearchRule - Attendee search rule for this activity plan
    AttendeesSearchRule Usersearchrule `json:"attendeesSearchRule"`


    // Facilitated - Whether the sessions created by this activity plan should be facilitated
    Facilitated bool `json:"facilitated"`


    // FacilitatorsSearchRule - Facilitator search rule for this activity plan
    FacilitatorsSearchRule Usersearchrule `json:"facilitatorsSearchRule"`


    // TransitionTimeMinutes - Transition time in minutes between facilitated sessions
    TransitionTimeMinutes int `json:"transitionTimeMinutes"`


    // ServiceGoalImpactOverrides - Allowable service goal impact override settings for this activity plan. If not set the business unit setting will be used
    ServiceGoalImpactOverrides Activityplanservicegoalimpactoverrides `json:"serviceGoalImpactOverrides"`


    // OptimizationObjective - The optimization objective of this activity plan
    OptimizationObjective string `json:"optimizationObjective"`


    // State - The state of this activity plan
    State string `json:"state"`


    // CountsAsPaidTime - Whether the activity should count as paid time
    CountsAsPaidTime bool `json:"countsAsPaidTime"`


    // FixedAvailability - Fixed availability configuration for the activity plan
    FixedAvailability []Fixedavailability `json:"fixedAvailability"`

}

// String returns a JSON representation of the model
func (o *Createactivityplanrequest) String() string {
    
     o.ManagementUnitIds = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.FixedAvailability = []Fixedavailability{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createactivityplanrequest) MarshalJSON() ([]byte, error) {
    type Alias Createactivityplanrequest

    if CreateactivityplanrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateactivityplanrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ManagementUnitIds []string `json:"managementUnitIds"`
        
        Description string `json:"description"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        VarType string `json:"type"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        InitialSchedulePeriod Schedulingperiod `json:"initialSchedulePeriod"`
        
        GroupSettings Groupsettings `json:"groupSettings"`
        
        RecurrenceSettings Recurrencesettings `json:"recurrenceSettings"`
        
        AttendeesSearchRule Usersearchrule `json:"attendeesSearchRule"`
        
        Facilitated bool `json:"facilitated"`
        
        FacilitatorsSearchRule Usersearchrule `json:"facilitatorsSearchRule"`
        
        TransitionTimeMinutes int `json:"transitionTimeMinutes"`
        
        ServiceGoalImpactOverrides Activityplanservicegoalimpactoverrides `json:"serviceGoalImpactOverrides"`
        
        OptimizationObjective string `json:"optimizationObjective"`
        
        State string `json:"state"`
        
        CountsAsPaidTime bool `json:"countsAsPaidTime"`
        
        FixedAvailability []Fixedavailability `json:"fixedAvailability"`
        *Alias
    }{

        


        
        ManagementUnitIds: []string{""},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        FixedAvailability: []Fixedavailability{{}},
        

        Alias: (*Alias)(u),
    })
}

