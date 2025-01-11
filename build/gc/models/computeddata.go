package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ComputeddataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ComputeddataDud struct { 
    


    

}

// Computeddata
type Computeddata struct { 
    // Weekly - Weekly time series for forecast data
    Weekly Weekly `json:"weekly"`


    // QuarterHour - Quarter hour time series for forecast data
    QuarterHour Quarterhourly `json:"quarterHour"`

}

// String returns a JSON representation of the model
func (o *Computeddata) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Computeddata) MarshalJSON() ([]byte, error) {
    type Alias Computeddata

    if ComputeddataMarshalled {
        return []byte("{}"), nil
    }
    ComputeddataMarshalled = true

    return json.Marshal(&struct {
        
        Weekly Weekly `json:"weekly"`
        
        QuarterHour Quarterhourly `json:"quarterHour"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

