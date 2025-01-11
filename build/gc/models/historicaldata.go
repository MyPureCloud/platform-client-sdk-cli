package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricaldataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricaldataDud struct { 
    


    

}

// Historicaldata
type Historicaldata struct { 
    // Daily - Daily time series for historical data
    Daily Daily `json:"daily"`


    // QuarterHour - Quarter hour time series for historical data
    QuarterHour Quarterhourly `json:"quarterHour"`

}

// String returns a JSON representation of the model
func (o *Historicaldata) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicaldata) MarshalJSON() ([]byte, error) {
    type Alias Historicaldata

    if HistoricaldataMarshalled {
        return []byte("{}"), nil
    }
    HistoricaldataMarshalled = true

    return json.Marshal(&struct {
        
        Daily Daily `json:"daily"`
        
        QuarterHour Quarterhourly `json:"quarterHour"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

