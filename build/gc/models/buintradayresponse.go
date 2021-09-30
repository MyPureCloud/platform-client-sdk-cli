package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuintradayresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuintradayresponseDud struct { 
    


    


    


    


    


    


    


    

}

// Buintradayresponse
type Buintradayresponse struct { 
    // StartDate - The start of the date range for which this data applies.  This is also the start reference point for the intervals represented in the various arrays. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - The end of the date range for which this data applies. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`


    // IntervalLengthMinutes - The aggregation period in minutes, which determines the interval duration of the returned data
    IntervalLengthMinutes int `json:"intervalLengthMinutes"`


    // NoDataReason - If not null, the reason there was no data for the request
    NoDataReason string `json:"noDataReason"`


    // Categories - The categories to which this data corresponds
    Categories []string `json:"categories"`


    // ShortTermForecast - Short term forecast reference
    ShortTermForecast Bushorttermforecastreference `json:"shortTermForecast"`


    // Schedule - Schedule reference
    Schedule Buschedulereference `json:"schedule"`


    // IntradayDataGroupings - Intraday data grouped by a single media type and set of planning group IDs
    IntradayDataGroupings []Buintradaydatagroup `json:"intradayDataGroupings"`

}

// String returns a JSON representation of the model
func (o *Buintradayresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Categories = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
     o.IntradayDataGroupings = []Buintradaydatagroup{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buintradayresponse) MarshalJSON() ([]byte, error) {
    type Alias Buintradayresponse

    if BuintradayresponseMarshalled {
        return []byte("{}"), nil
    }
    BuintradayresponseMarshalled = true

    return json.Marshal(&struct { 
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        IntervalLengthMinutes int `json:"intervalLengthMinutes"`
        
        NoDataReason string `json:"noDataReason"`
        
        Categories []string `json:"categories"`
        
        ShortTermForecast Bushorttermforecastreference `json:"shortTermForecast"`
        
        Schedule Buschedulereference `json:"schedule"`
        
        IntradayDataGroupings []Buintradaydatagroup `json:"intradayDataGroupings"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Categories: []string{""},
        

        

        

        

        

        

        
        IntradayDataGroupings: []Buintradaydatagroup{{}},
        

        
        Alias: (*Alias)(u),
    })
}

