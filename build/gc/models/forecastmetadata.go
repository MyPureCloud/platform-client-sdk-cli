package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ForecastmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ForecastmetadataDud struct { 
    


    


    


    


    


    

}

// Forecastmetadata
type Forecastmetadata struct { 
    // DateCreated - Forecast creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // ForecastId - Forecast ID
    ForecastId string `json:"forecastId"`


    // IntervalLengthInMinutes - Interval length
    IntervalLengthInMinutes string `json:"intervalLengthInMinutes"`


    // Source - Forecast source
    Source string `json:"source"`


    // DateStart - Forecast start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // TimeZone - Timezone of the business unit
    TimeZone string `json:"timeZone"`

}

// String returns a JSON representation of the model
func (o *Forecastmetadata) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Forecastmetadata) MarshalJSON() ([]byte, error) {
    type Alias Forecastmetadata

    if ForecastmetadataMarshalled {
        return []byte("{}"), nil
    }
    ForecastmetadataMarshalled = true

    return json.Marshal(&struct {
        
        DateCreated time.Time `json:"dateCreated"`
        
        ForecastId string `json:"forecastId"`
        
        IntervalLengthInMinutes string `json:"intervalLengthInMinutes"`
        
        Source string `json:"source"`
        
        DateStart time.Time `json:"dateStart"`
        
        TimeZone string `json:"timeZone"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

