package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricaladherencedaymetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricaladherencedaymetricsDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Historicaladherencedaymetrics
type Historicaladherencedaymetrics struct { 
    // DayStartOffsetSecs - Start of day offset in seconds relative to query start time
    DayStartOffsetSecs int `json:"dayStartOffsetSecs"`


    // AdherenceScheduleSecs - Duration of schedule in seconds included for adherence percentage calculation
    AdherenceScheduleSecs int `json:"adherenceScheduleSecs"`


    // ConformanceScheduleSecs - Total scheduled duration in seconds for OnQueue activities
    ConformanceScheduleSecs int `json:"conformanceScheduleSecs"`


    // ConformanceActualSecs - Total actually worked duration in seconds for OnQueue activities
    ConformanceActualSecs int `json:"conformanceActualSecs"`


    // ExceptionCount - Total number of adherence exceptions for this user
    ExceptionCount int `json:"exceptionCount"`


    // ExceptionDurationSecs - Total duration in seconds of adherence exceptions for this user
    ExceptionDurationSecs int `json:"exceptionDurationSecs"`


    // ImpactSeconds - The impact duration in seconds of current adherence state for this user
    ImpactSeconds int `json:"impactSeconds"`


    // ScheduleLengthSecs - Total duration in seconds for all scheduled activities
    ScheduleLengthSecs int `json:"scheduleLengthSecs"`


    // ActualLengthSecs - Total duration in seconds for all actually worked activities
    ActualLengthSecs int `json:"actualLengthSecs"`


    // AdherencePercentage - Total adherence percentage for this user, in the scale of 0 - 100
    AdherencePercentage float64 `json:"adherencePercentage"`


    // ConformancePercentage - Total conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on queue time is greater than the scheduled on queue time for the same period.
    ConformancePercentage float64 `json:"conformancePercentage"`

}

// String returns a JSON representation of the model
func (o *Historicaladherencedaymetrics) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicaladherencedaymetrics) MarshalJSON() ([]byte, error) {
    type Alias Historicaladherencedaymetrics

    if HistoricaladherencedaymetricsMarshalled {
        return []byte("{}"), nil
    }
    HistoricaladherencedaymetricsMarshalled = true

    return json.Marshal(&struct {
        
        DayStartOffsetSecs int `json:"dayStartOffsetSecs"`
        
        AdherenceScheduleSecs int `json:"adherenceScheduleSecs"`
        
        ConformanceScheduleSecs int `json:"conformanceScheduleSecs"`
        
        ConformanceActualSecs int `json:"conformanceActualSecs"`
        
        ExceptionCount int `json:"exceptionCount"`
        
        ExceptionDurationSecs int `json:"exceptionDurationSecs"`
        
        ImpactSeconds int `json:"impactSeconds"`
        
        ScheduleLengthSecs int `json:"scheduleLengthSecs"`
        
        ActualLengthSecs int `json:"actualLengthSecs"`
        
        AdherencePercentage float64 `json:"adherencePercentage"`
        
        ConformancePercentage float64 `json:"conformancePercentage"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

