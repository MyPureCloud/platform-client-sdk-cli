package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuintradaydatagroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuintradaydatagroupDud struct { 
    


    


    


    


    


    


    

}

// Buintradaydatagroup
type Buintradaydatagroup struct { 
    // MediaType - The media type associated with this intraday group
    MediaType string `json:"mediaType"`


    // ForecastDataSummary - Forecast data summary for this date range
    ForecastDataSummary Buintradayforecastdata `json:"forecastDataSummary"`


    // ForecastDataPerInterval - Forecast data per interval for this date range
    ForecastDataPerInterval []Buintradayforecastdata `json:"forecastDataPerInterval"`


    // ScheduleDataSummary - Schedule data summary for this date range
    ScheduleDataSummary Buintradayscheduledata `json:"scheduleDataSummary"`


    // ScheduleDataPerInterval - Schedule data per interval for this date range
    ScheduleDataPerInterval []Buintradayscheduledata `json:"scheduleDataPerInterval"`


    // PerformancePredictionDataSummary - Performance prediction data summary for this date range
    PerformancePredictionDataSummary Intradayperformancepredictiondata `json:"performancePredictionDataSummary"`


    // PerformancePredictionDataPerInterval - Performance prediction data per interval for this date range
    PerformancePredictionDataPerInterval []Intradayperformancepredictiondata `json:"performancePredictionDataPerInterval"`

}

// String returns a JSON representation of the model
func (o *Buintradaydatagroup) String() string {
    
    
     o.ForecastDataPerInterval = []Buintradayforecastdata{{}} 
    
     o.ScheduleDataPerInterval = []Buintradayscheduledata{{}} 
    
     o.PerformancePredictionDataPerInterval = []Intradayperformancepredictiondata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buintradaydatagroup) MarshalJSON() ([]byte, error) {
    type Alias Buintradaydatagroup

    if BuintradaydatagroupMarshalled {
        return []byte("{}"), nil
    }
    BuintradaydatagroupMarshalled = true

    return json.Marshal(&struct {
        
        MediaType string `json:"mediaType"`
        
        ForecastDataSummary Buintradayforecastdata `json:"forecastDataSummary"`
        
        ForecastDataPerInterval []Buintradayforecastdata `json:"forecastDataPerInterval"`
        
        ScheduleDataSummary Buintradayscheduledata `json:"scheduleDataSummary"`
        
        ScheduleDataPerInterval []Buintradayscheduledata `json:"scheduleDataPerInterval"`
        
        PerformancePredictionDataSummary Intradayperformancepredictiondata `json:"performancePredictionDataSummary"`
        
        PerformancePredictionDataPerInterval []Intradayperformancepredictiondata `json:"performancePredictionDataPerInterval"`
        *Alias
    }{

        


        


        
        ForecastDataPerInterval: []Buintradayforecastdata{{}},
        


        


        
        ScheduleDataPerInterval: []Buintradayscheduledata{{}},
        


        


        
        PerformancePredictionDataPerInterval: []Intradayperformancepredictiondata{{}},
        

        Alias: (*Alias)(u),
    })
}

