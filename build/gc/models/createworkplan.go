package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateworkplanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateworkplanDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Createworkplan - Work plan information
type Createworkplan struct { 
    // Name - Name of this work plan
    Name string `json:"name"`


    // Enabled - Whether the work plan is enabled for scheduling
    Enabled bool `json:"enabled"`


    // ConstrainWeeklyPaidTime - Whether the weekly paid time constraint is enabled for this work plan
    ConstrainWeeklyPaidTime bool `json:"constrainWeeklyPaidTime"`


    // FlexibleWeeklyPaidTime - Whether the weekly paid time constraint is flexible for this work plan
    FlexibleWeeklyPaidTime bool `json:"flexibleWeeklyPaidTime"`


    // WeeklyExactPaidMinutes - Exact weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == false
    WeeklyExactPaidMinutes int `json:"weeklyExactPaidMinutes"`


    // WeeklyMinimumPaidMinutes - Minimum weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == true
    WeeklyMinimumPaidMinutes int `json:"weeklyMinimumPaidMinutes"`


    // WeeklyMaximumPaidMinutes - Maximum weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == true
    WeeklyMaximumPaidMinutes int `json:"weeklyMaximumPaidMinutes"`


    // ConstrainPaidTimeGranularity - Whether paid time granularity should be constrained for this workplan
    ConstrainPaidTimeGranularity bool `json:"constrainPaidTimeGranularity"`


    // PaidTimeGranularityMinutes - Granularity in minutes allowed for shift paid time in this work plan. Used if constrainPaidTimeGranularity == true
    PaidTimeGranularityMinutes int `json:"paidTimeGranularityMinutes"`


    // ConstrainMinimumTimeBetweenShifts - Whether the minimum time between shifts constraint is enabled for this work plan
    ConstrainMinimumTimeBetweenShifts bool `json:"constrainMinimumTimeBetweenShifts"`


    // MinimumTimeBetweenShiftsMinutes - Minimum time between shifts in minutes defined in this work plan. Used if constrainMinimumTimeBetweenShifts == true
    MinimumTimeBetweenShiftsMinutes int `json:"minimumTimeBetweenShiftsMinutes"`


    // MaximumDays - Maximum number days in a week allowed to be scheduled for this work plan
    MaximumDays int `json:"maximumDays"`


    // MinimumConsecutiveNonWorkingMinutesPerWeek - Minimum amount of consecutive non working minutes per week that agents who are assigned this work plan are allowed to have off
    MinimumConsecutiveNonWorkingMinutesPerWeek int `json:"minimumConsecutiveNonWorkingMinutesPerWeek"`


    // ConstrainMaximumConsecutiveWorkingWeekends - Whether to constrain the maximum consecutive working weekends
    ConstrainMaximumConsecutiveWorkingWeekends bool `json:"constrainMaximumConsecutiveWorkingWeekends"`


    // MaximumConsecutiveWorkingWeekends - The maximum number of consecutive weekends that agents who are assigned to this work plan are allowed to work
    MaximumConsecutiveWorkingWeekends int `json:"maximumConsecutiveWorkingWeekends"`


    // MinimumWorkingDaysPerWeek - The minimum number of days that agents assigned to a work plan must work per week
    MinimumWorkingDaysPerWeek int `json:"minimumWorkingDaysPerWeek"`


    // ConstrainMaximumConsecutiveWorkingDays - Whether to constrain the maximum consecutive working days
    ConstrainMaximumConsecutiveWorkingDays bool `json:"constrainMaximumConsecutiveWorkingDays"`


    // MaximumConsecutiveWorkingDays - The maximum number of consecutive days that agents assigned to this work plan are allowed to work. Used if constrainMaximumConsecutiveWorkingDays == true
    MaximumConsecutiveWorkingDays int `json:"maximumConsecutiveWorkingDays"`


    // MinimumShiftStartDistanceMinutes - The time period in minutes for the duration between the start times of two consecutive working days
    MinimumShiftStartDistanceMinutes int `json:"minimumShiftStartDistanceMinutes"`


    // MinimumDaysOffPerPlanningPeriod - Minimum days off in the planning period
    MinimumDaysOffPerPlanningPeriod int `json:"minimumDaysOffPerPlanningPeriod"`


    // MaximumDaysOffPerPlanningPeriod - Maximum days off in the planning period
    MaximumDaysOffPerPlanningPeriod int `json:"maximumDaysOffPerPlanningPeriod"`


    // MinimumPaidMinutesPerPlanningPeriod - Minimum paid minutes in the planning period
    MinimumPaidMinutesPerPlanningPeriod int `json:"minimumPaidMinutesPerPlanningPeriod"`


    // MaximumPaidMinutesPerPlanningPeriod - Maximum paid minutes in the planning period
    MaximumPaidMinutesPerPlanningPeriod int `json:"maximumPaidMinutesPerPlanningPeriod"`


    // OptionalDays - Optional days to schedule for this work plan
    OptionalDays Setwrapperdayofweek `json:"optionalDays"`


    // ShiftStartVarianceType - This constraint ensures that an agent starts each workday within a user-defined time threshold
    ShiftStartVarianceType string `json:"shiftStartVarianceType"`


    // ShiftStartVariances - Variance in minutes among start times of shifts in this work plan
    ShiftStartVariances Listwrappershiftstartvariance `json:"shiftStartVariances"`


    // Shifts - Shifts in this work plan
    Shifts []Createworkplanshift `json:"shifts"`


    // Agents - Agents in this work plan
    Agents []Userreference `json:"agents"`

}

// String returns a JSON representation of the model
func (o *Createworkplan) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Shifts = []Createworkplanshift{{}} 
    
    
    
     o.Agents = []Userreference{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createworkplan) MarshalJSON() ([]byte, error) {
    type Alias Createworkplan

    if CreateworkplanMarshalled {
        return []byte("{}"), nil
    }
    CreateworkplanMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Enabled bool `json:"enabled"`
        
        ConstrainWeeklyPaidTime bool `json:"constrainWeeklyPaidTime"`
        
        FlexibleWeeklyPaidTime bool `json:"flexibleWeeklyPaidTime"`
        
        WeeklyExactPaidMinutes int `json:"weeklyExactPaidMinutes"`
        
        WeeklyMinimumPaidMinutes int `json:"weeklyMinimumPaidMinutes"`
        
        WeeklyMaximumPaidMinutes int `json:"weeklyMaximumPaidMinutes"`
        
        ConstrainPaidTimeGranularity bool `json:"constrainPaidTimeGranularity"`
        
        PaidTimeGranularityMinutes int `json:"paidTimeGranularityMinutes"`
        
        ConstrainMinimumTimeBetweenShifts bool `json:"constrainMinimumTimeBetweenShifts"`
        
        MinimumTimeBetweenShiftsMinutes int `json:"minimumTimeBetweenShiftsMinutes"`
        
        MaximumDays int `json:"maximumDays"`
        
        MinimumConsecutiveNonWorkingMinutesPerWeek int `json:"minimumConsecutiveNonWorkingMinutesPerWeek"`
        
        ConstrainMaximumConsecutiveWorkingWeekends bool `json:"constrainMaximumConsecutiveWorkingWeekends"`
        
        MaximumConsecutiveWorkingWeekends int `json:"maximumConsecutiveWorkingWeekends"`
        
        MinimumWorkingDaysPerWeek int `json:"minimumWorkingDaysPerWeek"`
        
        ConstrainMaximumConsecutiveWorkingDays bool `json:"constrainMaximumConsecutiveWorkingDays"`
        
        MaximumConsecutiveWorkingDays int `json:"maximumConsecutiveWorkingDays"`
        
        MinimumShiftStartDistanceMinutes int `json:"minimumShiftStartDistanceMinutes"`
        
        MinimumDaysOffPerPlanningPeriod int `json:"minimumDaysOffPerPlanningPeriod"`
        
        MaximumDaysOffPerPlanningPeriod int `json:"maximumDaysOffPerPlanningPeriod"`
        
        MinimumPaidMinutesPerPlanningPeriod int `json:"minimumPaidMinutesPerPlanningPeriod"`
        
        MaximumPaidMinutesPerPlanningPeriod int `json:"maximumPaidMinutesPerPlanningPeriod"`
        
        OptionalDays Setwrapperdayofweek `json:"optionalDays"`
        
        ShiftStartVarianceType string `json:"shiftStartVarianceType"`
        
        ShiftStartVariances Listwrappershiftstartvariance `json:"shiftStartVariances"`
        
        Shifts []Createworkplanshift `json:"shifts"`
        
        Agents []Userreference `json:"agents"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Shifts: []Createworkplanshift{{}},
        

        

        
        Agents: []Userreference{{}},
        

        
        Alias: (*Alias)(u),
    })
}

