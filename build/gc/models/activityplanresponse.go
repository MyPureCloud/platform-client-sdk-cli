package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Activityplanresponse
type Activityplanresponse struct { 
    


    // Name - The name of the activity plan
    Name string `json:"name"`


    // ManagementUnits - The management units to which this activity plan applies. Empty list or null means this activity plan applies to the entire business unit
    ManagementUnits []Managementunitreference `json:"managementUnits"`


    // Description - The description of this activity plan
    Description string `json:"description"`


    // ActivityCode - The activity code associated with this activity plan. It is recommended to load and cache the entire list of activity codes rather than look up individual codes
    ActivityCode Activitycodereference `json:"activityCode"`


    // VarType - The type of the activity plan
    VarType string `json:"type"`


    // InitialSchedulePeriod - The initial schedule period of the activity plan
    InitialSchedulePeriod Schedulingperiod `json:"initialSchedulePeriod"`


    // LengthMinutes - The length of the activity in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // GroupSettings - Group settings for this activity plan
    GroupSettings Groupsettings `json:"groupSettings"`


    // RecurrenceSettings - Recurrence settings for this activity plan
    RecurrenceSettings Recurrencesettings `json:"recurrenceSettings"`


    // AttendeesSearchRule - Attendee search rule for this activity plan
    AttendeesSearchRule Usersearchrule `json:"attendeesSearchRule"`


    // Facilitated - Whether the sessions created by this activity plan should be facilitated
    Facilitated bool `json:"facilitated"`


    // FacilitatorsSearchRule - Facilitator search rule for this activity plan
    FacilitatorsSearchRule Usersearchrule `json:"facilitatorsSearchRule"`


    // TransitionTimeMinutes - Transition time in minutes between facilitated sessions
    TransitionTimeMinutes int `json:"transitionTimeMinutes"`


    // ServiceGoalImpactOverrides - Allowable service goal impact override settings for this activity plan
    ServiceGoalImpactOverrides Activityplanservicegoalimpactoverrides `json:"serviceGoalImpactOverrides"`


    // OptimizationObjective - The optimization objective of this activity plan
    OptimizationObjective string `json:"optimizationObjective"`


    // FixedAvailability - Fixed availability configuration for this activity plan
    FixedAvailability []Fixedavailability `json:"fixedAvailability"`


    // State - The state of this activity plan
    State string `json:"state"`


    // CountsAsPaidTime - Whether the activity should count as paid time
    CountsAsPaidTime bool `json:"countsAsPaidTime"`


    // CreatedDate - The date the activity plan was created, in ISO-8601 format
    CreatedDate time.Time `json:"createdDate"`


    // CreatedBy - The user who created this activity plan
    CreatedBy Userreference `json:"createdBy"`


    // ModifiedDate - The date the activity plan was modified, in ISO-8601 format
    ModifiedDate time.Time `json:"modifiedDate"`


    // ModifiedBy - The last user to modify this activity plan. The id may be 'System' if it was an automated process
    ModifiedBy Userreference `json:"modifiedBy"`


    // LastRunDate - The date on which the activity plan was last manually run, in ISO-8601 format
    LastRunDate time.Time `json:"lastRunDate"`


    // LastRunBy - The last user to run this activity plan
    LastRunBy Userreference `json:"lastRunBy"`


    

}

// String returns a JSON representation of the model
func (o *Activityplanresponse) String() string {
    
     o.ManagementUnits = []Managementunitreference{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.FixedAvailability = []Fixedavailability{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanresponse) MarshalJSON() ([]byte, error) {
    type Alias Activityplanresponse

    if ActivityplanresponseMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ManagementUnits []Managementunitreference `json:"managementUnits"`
        
        Description string `json:"description"`
        
        ActivityCode Activitycodereference `json:"activityCode"`
        
        VarType string `json:"type"`
        
        InitialSchedulePeriod Schedulingperiod `json:"initialSchedulePeriod"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        GroupSettings Groupsettings `json:"groupSettings"`
        
        RecurrenceSettings Recurrencesettings `json:"recurrenceSettings"`
        
        AttendeesSearchRule Usersearchrule `json:"attendeesSearchRule"`
        
        Facilitated bool `json:"facilitated"`
        
        FacilitatorsSearchRule Usersearchrule `json:"facilitatorsSearchRule"`
        
        TransitionTimeMinutes int `json:"transitionTimeMinutes"`
        
        ServiceGoalImpactOverrides Activityplanservicegoalimpactoverrides `json:"serviceGoalImpactOverrides"`
        
        OptimizationObjective string `json:"optimizationObjective"`
        
        FixedAvailability []Fixedavailability `json:"fixedAvailability"`
        
        State string `json:"state"`
        
        CountsAsPaidTime bool `json:"countsAsPaidTime"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        LastRunDate time.Time `json:"lastRunDate"`
        
        LastRunBy Userreference `json:"lastRunBy"`
        *Alias
    }{

        


        


        
        ManagementUnits: []Managementunitreference{{}},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        FixedAvailability: []Fixedavailability{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

