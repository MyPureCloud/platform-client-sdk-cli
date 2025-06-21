package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherencebulkuserdaymetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherencebulkuserdaymetricsDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Wfmhistoricaladherencebulkuserdaymetrics
type Wfmhistoricaladherencebulkuserdaymetrics struct { 
    // DayStartOffsetSeconds - Start of day offset in seconds relative to query start time
    DayStartOffsetSeconds int `json:"dayStartOffsetSeconds"`


    // AdherenceScheduleSeconds - Duration of schedule in seconds included for adherence percentage calculation
    AdherenceScheduleSeconds int `json:"adherenceScheduleSeconds"`


    // ConformanceScheduleSeconds - Total scheduled duration in seconds for OnQueue activities
    ConformanceScheduleSeconds int `json:"conformanceScheduleSeconds"`


    // ConformanceActualSeconds - Total actually worked duration in seconds for OnQueue activities
    ConformanceActualSeconds int `json:"conformanceActualSeconds"`


    // ExceptionCount - Total number of adherence exceptions for this user
    ExceptionCount int `json:"exceptionCount"`


    // ExceptionDurationSeconds - Total duration in seconds of adherence exceptions for this user
    ExceptionDurationSeconds int `json:"exceptionDurationSeconds"`


    // ImpactSeconds - The impact duration in seconds of current adherence state for this user
    ImpactSeconds int `json:"impactSeconds"`


    // ScheduleLengthSeconds - Total duration in seconds for all scheduled activities
    ScheduleLengthSeconds int `json:"scheduleLengthSeconds"`


    // ActualLengthSeconds - Total duration in seconds for all actually worked activities
    ActualLengthSeconds int `json:"actualLengthSeconds"`


    // AdherencePercentage - Total adherence percentage for this user, in the scale of 0 - 100
    AdherencePercentage float64 `json:"adherencePercentage"`


    // ConformancePercentage - Total conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on-queue time is greater than the scheduled on-queue time for the same period.
    ConformancePercentage float64 `json:"conformancePercentage"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkuserdaymetrics) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherencebulkuserdaymetrics) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherencebulkuserdaymetrics

    if WfmhistoricaladherencebulkuserdaymetricsMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherencebulkuserdaymetricsMarshalled = true

    return json.Marshal(&struct {
        
        DayStartOffsetSeconds int `json:"dayStartOffsetSeconds"`
        
        AdherenceScheduleSeconds int `json:"adherenceScheduleSeconds"`
        
        ConformanceScheduleSeconds int `json:"conformanceScheduleSeconds"`
        
        ConformanceActualSeconds int `json:"conformanceActualSeconds"`
        
        ExceptionCount int `json:"exceptionCount"`
        
        ExceptionDurationSeconds int `json:"exceptionDurationSeconds"`
        
        ImpactSeconds int `json:"impactSeconds"`
        
        ScheduleLengthSeconds int `json:"scheduleLengthSeconds"`
        
        ActualLengthSeconds int `json:"actualLengthSeconds"`
        
        AdherencePercentage float64 `json:"adherencePercentage"`
        
        ConformancePercentage float64 `json:"conformancePercentage"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

