package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherencebulkuserresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherencebulkuserresultDud struct { 
    


    


    


    


    


    


    

}

// Wfmhistoricaladherencebulkuserresult
type Wfmhistoricaladherencebulkuserresult struct { 
    // UserId - The ID of the user for whom the adherence is queried
    UserId string `json:"userId"`


    // AdherencePercentage - Adherence percentage for this user, in the scale of 0 - 100
    AdherencePercentage float64 `json:"adherencePercentage"`


    // ConformancePercentage - Conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on queue time is greater than the scheduled on queue time for the same period.
    ConformancePercentage float64 `json:"conformancePercentage"`


    // Impact - The impact of the current adherence state for this user
    Impact string `json:"impact"`


    // ExceptionInfo - List of adherence exceptions for this user
    ExceptionInfo []Historicaladherenceexceptioninfo `json:"exceptionInfo"`


    // Actuals - List of adherence actuals for this user
    Actuals []Historicaladherenceactuals `json:"actuals"`


    // DayMetrics - Adherence and conformance metrics for days in query range
    DayMetrics []Wfmhistoricaladherencebulkuserdaymetrics `json:"dayMetrics"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkuserresult) String() string {
    
    
    
    
     o.ExceptionInfo = []Historicaladherenceexceptioninfo{{}} 
     o.Actuals = []Historicaladherenceactuals{{}} 
     o.DayMetrics = []Wfmhistoricaladherencebulkuserdaymetrics{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherencebulkuserresult) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherencebulkuserresult

    if WfmhistoricaladherencebulkuserresultMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherencebulkuserresultMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        AdherencePercentage float64 `json:"adherencePercentage"`
        
        ConformancePercentage float64 `json:"conformancePercentage"`
        
        Impact string `json:"impact"`
        
        ExceptionInfo []Historicaladherenceexceptioninfo `json:"exceptionInfo"`
        
        Actuals []Historicaladherenceactuals `json:"actuals"`
        
        DayMetrics []Wfmhistoricaladherencebulkuserdaymetrics `json:"dayMetrics"`
        *Alias
    }{

        


        


        


        


        
        ExceptionInfo: []Historicaladherenceexceptioninfo{{}},
        


        
        Actuals: []Historicaladherenceactuals{{}},
        


        
        DayMetrics: []Wfmhistoricaladherencebulkuserdaymetrics{{}},
        

        Alias: (*Alias)(u),
    })
}

