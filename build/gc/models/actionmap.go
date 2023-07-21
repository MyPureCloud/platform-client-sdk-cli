package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionmapMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionmapDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`


    


    


    


    

}

// Actionmap
type Actionmap struct { 
    


    // Version - The version of the action map.
    Version int `json:"version"`


    // IsActive - Whether the action map is active.
    IsActive bool `json:"isActive"`


    // DisplayName - Display name of the action map.
    DisplayName string `json:"displayName"`


    // TriggerWithSegments - Trigger action map if any segment in the list is assigned to a given customer.
    TriggerWithSegments []string `json:"triggerWithSegments"`


    // TriggerWithEventConditions - List of event conditions that must be satisfied to trigger the action map.
    TriggerWithEventConditions []Eventcondition `json:"triggerWithEventConditions"`


    // TriggerWithOutcomeProbabilityConditions - Probability conditions for outcomes that must be satisfied to trigger the action map.
    TriggerWithOutcomeProbabilityConditions []Outcomeprobabilitycondition `json:"triggerWithOutcomeProbabilityConditions"`


    // TriggerWithOutcomePercentileConditions - (deprecated - use triggerWithOutcomeQuantileConditions instead) Percentile conditions for outcomes that must be satisfied to trigger the action map.
    TriggerWithOutcomePercentileConditions []Outcomepercentilecondition `json:"triggerWithOutcomePercentileConditions"`


    // TriggerWithOutcomeQuantileConditions - Quantile conditions for outcomes that must be satisfied to trigger the action map.
    TriggerWithOutcomeQuantileConditions []Outcomequantilecondition `json:"triggerWithOutcomeQuantileConditions"`


    // PageUrlConditions - URL conditions that a page must match for web actions to be displayable.
    PageUrlConditions []Urlcondition `json:"pageUrlConditions"`


    // Activation - Type of activation.
    Activation Activation `json:"activation"`


    // Weight - Weight of the action map with higher number denoting higher weight.
    Weight int `json:"weight"`


    // Action - The action that will be executed if this action map is triggered.
    Action Actionmapaction `json:"action"`


    // ActionMapScheduleGroups - The action map's associated schedule groups.
    ActionMapScheduleGroups Actionmapschedulegroups `json:"actionMapScheduleGroups"`


    // IgnoreFrequencyCap - Override organization-level frequency cap and always offer web engagements from this action map.
    IgnoreFrequencyCap bool `json:"ignoreFrequencyCap"`


    


    // CreatedDate - Timestamp indicating when the action map was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // ModifiedDate - Timestamp indicating when the action map was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    // StartDate - Timestamp at which the action map is scheduled to start firing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - Timestamp at which the action map is scheduled to stop firing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`

}

// String returns a JSON representation of the model
func (o *Actionmap) String() string {
    
    
    
     o.TriggerWithSegments = []string{""} 
     o.TriggerWithEventConditions = []Eventcondition{{}} 
     o.TriggerWithOutcomeProbabilityConditions = []Outcomeprobabilitycondition{{}} 
     o.TriggerWithOutcomePercentileConditions = []Outcomepercentilecondition{{}} 
     o.TriggerWithOutcomeQuantileConditions = []Outcomequantilecondition{{}} 
     o.PageUrlConditions = []Urlcondition{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionmap) MarshalJSON() ([]byte, error) {
    type Alias Actionmap

    if ActionmapMarshalled {
        return []byte("{}"), nil
    }
    ActionmapMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        
        IsActive bool `json:"isActive"`
        
        DisplayName string `json:"displayName"`
        
        TriggerWithSegments []string `json:"triggerWithSegments"`
        
        TriggerWithEventConditions []Eventcondition `json:"triggerWithEventConditions"`
        
        TriggerWithOutcomeProbabilityConditions []Outcomeprobabilitycondition `json:"triggerWithOutcomeProbabilityConditions"`
        
        TriggerWithOutcomePercentileConditions []Outcomepercentilecondition `json:"triggerWithOutcomePercentileConditions"`
        
        TriggerWithOutcomeQuantileConditions []Outcomequantilecondition `json:"triggerWithOutcomeQuantileConditions"`
        
        PageUrlConditions []Urlcondition `json:"pageUrlConditions"`
        
        Activation Activation `json:"activation"`
        
        Weight int `json:"weight"`
        
        Action Actionmapaction `json:"action"`
        
        ActionMapScheduleGroups Actionmapschedulegroups `json:"actionMapScheduleGroups"`
        
        IgnoreFrequencyCap bool `json:"ignoreFrequencyCap"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        *Alias
    }{

        


        


        


        


        
        TriggerWithSegments: []string{""},
        


        
        TriggerWithEventConditions: []Eventcondition{{}},
        


        
        TriggerWithOutcomeProbabilityConditions: []Outcomeprobabilitycondition{{}},
        


        
        TriggerWithOutcomePercentileConditions: []Outcomepercentilecondition{{}},
        


        
        TriggerWithOutcomeQuantileConditions: []Outcomequantilecondition{{}},
        


        
        PageUrlConditions: []Urlcondition{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

