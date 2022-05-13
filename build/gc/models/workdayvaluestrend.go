package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkdayvaluestrendMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkdayvaluestrendDud struct { 
    DateStartWorkday time.Time `json:"dateStartWorkday"`


    DateEndWorkday time.Time `json:"dateEndWorkday"`


    DateReferenceWorkday time.Time `json:"dateReferenceWorkday"`


    Division Division `json:"division"`


    User Userreference `json:"user"`


    Timezone string `json:"timezone"`


    Results []Workdayvaluesmetricitem `json:"results"`


    PerformanceProfile Addressableentityref `json:"performanceProfile"`


    Metric Addressableentityref `json:"metric"`

}

// Workdayvaluestrend
type Workdayvaluestrend struct { 
    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Workdayvaluestrend) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workdayvaluestrend) MarshalJSON() ([]byte, error) {
    type Alias Workdayvaluestrend

    if WorkdayvaluestrendMarshalled {
        return []byte("{}"), nil
    }
    WorkdayvaluestrendMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

