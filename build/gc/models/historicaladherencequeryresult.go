package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricaladherencequeryresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricaladherencequeryresultDud struct { 
    


    


    


    


    


    


    


    


    

}

// Historicaladherencequeryresult
type Historicaladherencequeryresult struct { 
    // UserId - The ID of the user for whom the adherence is queried
    UserId string `json:"userId"`


    // StartDate - Beginning of the date range that was queried, in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - End of the date range that was queried, in ISO-8601 format. If it was not set, end date will be set to the queried time
    EndDate time.Time `json:"endDate"`


    // AdherencePercentage - Adherence percentage for this user, in the scale of 0 - 100
    AdherencePercentage float64 `json:"adherencePercentage"`


    // ConformancePercentage - Conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on queue time is greater than the scheduled on queue time for the same period.
    ConformancePercentage float64 `json:"conformancePercentage"`


    // Impact - The impact of the current adherence state for this user
    Impact string `json:"impact"`


    // ExceptionInfo - List of adherence exceptions for this user
    ExceptionInfo []Historicaladherenceexceptioninfo `json:"exceptionInfo"`


    // DayMetrics - Adherence and conformance metrics for days in query range
    DayMetrics []Historicaladherencedaymetrics `json:"dayMetrics"`


    // Actuals - List of actual activity with offset for this user
    Actuals []Historicaladherenceactuals `json:"actuals"`

}

// String returns a JSON representation of the model
func (o *Historicaladherencequeryresult) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.ExceptionInfo = []Historicaladherenceexceptioninfo{{}} 
    
    
    
     o.DayMetrics = []Historicaladherencedaymetrics{{}} 
    
    
    
     o.Actuals = []Historicaladherenceactuals{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicaladherencequeryresult) MarshalJSON() ([]byte, error) {
    type Alias Historicaladherencequeryresult

    if HistoricaladherencequeryresultMarshalled {
        return []byte("{}"), nil
    }
    HistoricaladherencequeryresultMarshalled = true

    return json.Marshal(&struct { 
        UserId string `json:"userId"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        AdherencePercentage float64 `json:"adherencePercentage"`
        
        ConformancePercentage float64 `json:"conformancePercentage"`
        
        Impact string `json:"impact"`
        
        ExceptionInfo []Historicaladherenceexceptioninfo `json:"exceptionInfo"`
        
        DayMetrics []Historicaladherencedaymetrics `json:"dayMetrics"`
        
        Actuals []Historicaladherenceactuals `json:"actuals"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        ExceptionInfo: []Historicaladherenceexceptioninfo{{}},
        

        

        
        DayMetrics: []Historicaladherencedaymetrics{{}},
        

        

        
        Actuals: []Historicaladherenceactuals{{}},
        

        
        Alias: (*Alias)(u),
    })
}

