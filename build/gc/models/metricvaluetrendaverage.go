package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetricvaluetrendaverageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetricvaluetrendaverageDud struct { 
    DateStartWorkday time.Time `json:"dateStartWorkday"`


    DateEndWorkday time.Time `json:"dateEndWorkday"`


    DateReferenceWorkday time.Time `json:"dateReferenceWorkday"`


    Division Division `json:"division"`


    User Userreference `json:"user"`


    Timezone string `json:"timezone"`


    Result Workdayvaluesmetricitem `json:"result"`


    PerformanceProfile Addressableentityref `json:"performanceProfile"`


    Metric Addressableentityref `json:"metric"`

}

// Metricvaluetrendaverage
type Metricvaluetrendaverage struct { 
    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Metricvaluetrendaverage) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metricvaluetrendaverage) MarshalJSON() ([]byte, error) {
    type Alias Metricvaluetrendaverage

    if MetricvaluetrendaverageMarshalled {
        return []byte("{}"), nil
    }
    MetricvaluetrendaverageMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

