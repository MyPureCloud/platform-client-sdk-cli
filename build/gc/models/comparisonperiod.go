package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ComparisonperiodMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ComparisonperiodDud struct { 
    Id string `json:"id"`


    Kpi string `json:"kpi"`


    DateStarted time.Time `json:"dateStarted"`


    DateEnded time.Time `json:"dateEnded"`


    KpiResults []Kpiresult `json:"kpiResults"`


    SelfUri string `json:"selfUri"`

}

// Comparisonperiod
type Comparisonperiod struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Comparisonperiod) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Comparisonperiod) MarshalJSON() ([]byte, error) {
    type Alias Comparisonperiod

    if ComparisonperiodMarshalled {
        return []byte("{}"), nil
    }
    ComparisonperiodMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

