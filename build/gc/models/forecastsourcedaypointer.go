package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ForecastsourcedaypointerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ForecastsourcedaypointerDud struct { 
    


    


    


    


    

}

// Forecastsourcedaypointer - Pointer to look up source data for a short term forecast
type Forecastsourcedaypointer struct { 
    // DayOfWeek - The forecast day of week for this source data
    DayOfWeek string `json:"dayOfWeek"`


    // Weight - The relative weight to apply to this source data item for weighted averages
    Weight int `json:"weight"`


    // Date - The date this source data represents, in yyyy-MM-dd format
    Date string `json:"date"`


    // FileName - The name of the source file this data came from if it originated from a data import
    FileName string `json:"fileName"`


    // DataKey - The key to look up the forecast source data for this source day
    DataKey string `json:"dataKey"`

}

// String returns a JSON representation of the model
func (o *Forecastsourcedaypointer) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Forecastsourcedaypointer) MarshalJSON() ([]byte, error) {
    type Alias Forecastsourcedaypointer

    if ForecastsourcedaypointerMarshalled {
        return []byte("{}"), nil
    }
    ForecastsourcedaypointerMarshalled = true

    return json.Marshal(&struct { 
        DayOfWeek string `json:"dayOfWeek"`
        
        Weight int `json:"weight"`
        
        Date string `json:"date"`
        
        FileName string `json:"fileName"`
        
        DataKey string `json:"dataKey"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

