package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkdaymetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkdaymetricDud struct { 
    Metric Metric `json:"metric"`


    Objective Objective `json:"objective"`


    Points int `json:"points"`


    MaxPoints int `json:"maxPoints"`


    Value float64 `json:"value"`


    PunctualityEvents []Punctualityevent `json:"punctualityEvents"`

}

// Workdaymetric
type Workdaymetric struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Workdaymetric) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workdaymetric) MarshalJSON() ([]byte, error) {
    type Alias Workdaymetric

    if WorkdaymetricMarshalled {
        return []byte("{}"), nil
    }
    WorkdaymetricMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

