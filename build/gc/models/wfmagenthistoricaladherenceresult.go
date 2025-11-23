package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmagenthistoricaladherenceresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmagenthistoricaladherenceresultDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Wfmagenthistoricaladherenceresult
type Wfmagenthistoricaladherenceresult struct { 
    // User - The user who submitted the agent historical adherence query
    User Userreference `json:"user"`


    // StartDate - Beginning of the date range that was queried, in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - End of the date range that was queried, in ISO-8601 format. If it was not set, end date will be set to the queried time
    EndDate time.Time `json:"endDate"`


    // CalculationsCompletedDate - Completed date of calculations that was queried, in ISO-8601 format.
    CalculationsCompletedDate time.Time `json:"calculationsCompletedDate"`


    // TargetAdherencePercentage - Target percentage for this user, in the scale of 0 - 100
    TargetAdherencePercentage float64 `json:"targetAdherencePercentage"`


    // AdherencePercentage - Adherence percentage for this user, in the scale of 0 - 100
    AdherencePercentage float64 `json:"adherencePercentage"`


    // ConformancePercentage - Conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on queue time is greater than the scheduled on queue time for the same period.
    ConformancePercentage float64 `json:"conformancePercentage"`


    // Impact - The impact of the current adherence state for this user
    Impact string `json:"impact"`


    // ExceptionInfo - List of adherence exceptions for this user
    ExceptionInfo []Historicaladherenceexceptioninfo `json:"exceptionInfo"`


    // DayMetrics - Adherence and conformance metrics for days in query range
    DayMetrics []Agentadherencedaymetrics `json:"dayMetrics"`


    // Actuals - List of actual activity with offset for this user
    Actuals []Historicaladherenceactuals `json:"actuals"`


    // ScheduledActivities - List of scheduled activities for this user
    ScheduledActivities []Agentadherencescheduledactivity `json:"scheduledActivities"`


    // SecondaryPresenceLookupItems - List of secondary presence lookup ID to corresponding secondary presence ID item
    SecondaryPresenceLookupItems []Secondarypresencelookupitem `json:"secondaryPresenceLookupItems"`

}

// String returns a JSON representation of the model
func (o *Wfmagenthistoricaladherenceresult) String() string {
    
    
    
    
    
    
    
    
     o.ExceptionInfo = []Historicaladherenceexceptioninfo{{}} 
     o.DayMetrics = []Agentadherencedaymetrics{{}} 
     o.Actuals = []Historicaladherenceactuals{{}} 
     o.ScheduledActivities = []Agentadherencescheduledactivity{{}} 
     o.SecondaryPresenceLookupItems = []Secondarypresencelookupitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmagenthistoricaladherenceresult) MarshalJSON() ([]byte, error) {
    type Alias Wfmagenthistoricaladherenceresult

    if WfmagenthistoricaladherenceresultMarshalled {
        return []byte("{}"), nil
    }
    WfmagenthistoricaladherenceresultMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        CalculationsCompletedDate time.Time `json:"calculationsCompletedDate"`
        
        TargetAdherencePercentage float64 `json:"targetAdherencePercentage"`
        
        AdherencePercentage float64 `json:"adherencePercentage"`
        
        ConformancePercentage float64 `json:"conformancePercentage"`
        
        Impact string `json:"impact"`
        
        ExceptionInfo []Historicaladherenceexceptioninfo `json:"exceptionInfo"`
        
        DayMetrics []Agentadherencedaymetrics `json:"dayMetrics"`
        
        Actuals []Historicaladherenceactuals `json:"actuals"`
        
        ScheduledActivities []Agentadherencescheduledactivity `json:"scheduledActivities"`
        
        SecondaryPresenceLookupItems []Secondarypresencelookupitem `json:"secondaryPresenceLookupItems"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        ExceptionInfo: []Historicaladherenceexceptioninfo{{}},
        


        
        DayMetrics: []Agentadherencedaymetrics{{}},
        


        
        Actuals: []Historicaladherenceactuals{{}},
        


        
        ScheduledActivities: []Agentadherencescheduledactivity{{}},
        


        
        SecondaryPresenceLookupItems: []Secondarypresencelookupitem{{}},
        

        Alias: (*Alias)(u),
    })
}

