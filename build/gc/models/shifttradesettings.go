package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradesettingsDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Shifttradesettings
type Shifttradesettings struct { 
    // Enabled - Whether shift trading is enabled for this management unit
    Enabled bool `json:"enabled"`


    // AutoReview - Whether automatic shift trade review is enabled according to the rules defined in for this management unit
    AutoReview bool `json:"autoReview"`


    // AllowDirectTrades - Whether direct shift trades between agents are allowed
    AllowDirectTrades bool `json:"allowDirectTrades"`


    // MinHoursInFuture - The minimum number of hours in the future shift trades are allowed
    MinHoursInFuture int `json:"minHoursInFuture"`


    // UnequalPaid - How to handle shift trades which involve unequal paid times
    UnequalPaid string `json:"unequalPaid"`


    // OneSided - How to handle one-sided shift trades
    OneSided string `json:"oneSided"`


    // WeeklyMinPaidViolations - How to handle shift trades which result in violations of weekly minimum paid time constraint
    WeeklyMinPaidViolations string `json:"weeklyMinPaidViolations"`


    // WeeklyMaxPaidViolations - How to handle shift trades which result in violations of weekly maximum paid time constraint
    WeeklyMaxPaidViolations string `json:"weeklyMaxPaidViolations"`


    // RequiresMatchingQueues - Whether to constrain shift trades to agents with matching queues
    RequiresMatchingQueues bool `json:"requiresMatchingQueues"`


    // RequiresMatchingLanguages - Whether to constrain shift trades to agents with matching languages
    RequiresMatchingLanguages bool `json:"requiresMatchingLanguages"`


    // RequiresMatchingSkills - Whether to constrain shift trades to agents with matching skills
    RequiresMatchingSkills bool `json:"requiresMatchingSkills"`


    // RequiresMatchingPlanningGroups - Whether to constrain shift trades to agents with matching planning groups
    RequiresMatchingPlanningGroups bool `json:"requiresMatchingPlanningGroups"`


    // ActivityCategoryRules - Rules that specify what to do with activity categories that are part of a shift defined in a trade
    ActivityCategoryRules []Shifttradeactivityrule `json:"activityCategoryRules"`


    // MaxTradeSpanWeeks - The maximum number of weeks a shift trade can span
    MaxTradeSpanWeeks int `json:"maxTradeSpanWeeks"`


    // MaxTradesPerAgentPerWeek - The maximum number of shift trades an agent can submit per week
    MaxTradesPerAgentPerWeek int `json:"maxTradesPerAgentPerWeek"`


    // MinMinutesBetweenShifts - The minimum number of minutes between shifts
    MinMinutesBetweenShifts int `json:"minMinutesBetweenShifts"`


    // PlanningPeriodMinPaidViolations - How to handle shift trades which result in violations of planning period minimum paid time constraint
    PlanningPeriodMinPaidViolations string `json:"planningPeriodMinPaidViolations"`


    // PlanningPeriodMaxPaidViolations - How to handle shift trades which result in violations of planning period maximum paid time constraint
    PlanningPeriodMaxPaidViolations string `json:"planningPeriodMaxPaidViolations"`


    // MinMinutesBetweenShiftsViolations - How to handle shift trades which result in violations of minimum number of minutes between shifts constraint
    MinMinutesBetweenShiftsViolations string `json:"minMinutesBetweenShiftsViolations"`

}

// String returns a JSON representation of the model
func (o *Shifttradesettings) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.ActivityCategoryRules = []Shifttradeactivityrule{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradesettings) MarshalJSON() ([]byte, error) {
    type Alias Shifttradesettings

    if ShifttradesettingsMarshalled {
        return []byte("{}"), nil
    }
    ShifttradesettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        AutoReview bool `json:"autoReview"`
        
        AllowDirectTrades bool `json:"allowDirectTrades"`
        
        MinHoursInFuture int `json:"minHoursInFuture"`
        
        UnequalPaid string `json:"unequalPaid"`
        
        OneSided string `json:"oneSided"`
        
        WeeklyMinPaidViolations string `json:"weeklyMinPaidViolations"`
        
        WeeklyMaxPaidViolations string `json:"weeklyMaxPaidViolations"`
        
        RequiresMatchingQueues bool `json:"requiresMatchingQueues"`
        
        RequiresMatchingLanguages bool `json:"requiresMatchingLanguages"`
        
        RequiresMatchingSkills bool `json:"requiresMatchingSkills"`
        
        RequiresMatchingPlanningGroups bool `json:"requiresMatchingPlanningGroups"`
        
        ActivityCategoryRules []Shifttradeactivityrule `json:"activityCategoryRules"`
        
        MaxTradeSpanWeeks int `json:"maxTradeSpanWeeks"`
        
        MaxTradesPerAgentPerWeek int `json:"maxTradesPerAgentPerWeek"`
        
        MinMinutesBetweenShifts int `json:"minMinutesBetweenShifts"`
        
        PlanningPeriodMinPaidViolations string `json:"planningPeriodMinPaidViolations"`
        
        PlanningPeriodMaxPaidViolations string `json:"planningPeriodMaxPaidViolations"`
        
        MinMinutesBetweenShiftsViolations string `json:"minMinutesBetweenShiftsViolations"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        ActivityCategoryRules: []Shifttradeactivityrule{{}},
        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

