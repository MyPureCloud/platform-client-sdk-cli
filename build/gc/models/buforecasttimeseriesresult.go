package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuforecasttimeseriesresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuforecasttimeseriesresultDud struct { 
    


    


    

}

// Buforecasttimeseriesresult
type Buforecasttimeseriesresult struct { 
    // Metric - The metric this result applies to
    Metric string `json:"metric"`


    // ForecastingMethod - The forecasting method that was used for this metric
    ForecastingMethod string `json:"forecastingMethod"`


    // ForecastType - The forecasting type in this forecast result
    ForecastType string `json:"forecastType"`

}

// String returns a JSON representation of the model
func (o *Buforecasttimeseriesresult) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buforecasttimeseriesresult) MarshalJSON() ([]byte, error) {
    type Alias Buforecasttimeseriesresult

    if BuforecasttimeseriesresultMarshalled {
        return []byte("{}"), nil
    }
    BuforecasttimeseriesresultMarshalled = true

    return json.Marshal(&struct { 
        Metric string `json:"metric"`
        
        ForecastingMethod string `json:"forecastingMethod"`
        
        ForecastType string `json:"forecastType"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

