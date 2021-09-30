package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ForecastplanninggroupsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ForecastplanninggroupsresponseDud struct { 
    

}

// Forecastplanninggroupsresponse
type Forecastplanninggroupsresponse struct { 
    // Entities
    Entities []Forecastplanninggroupresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Forecastplanninggroupsresponse) String() string {
    
    
     o.Entities = []Forecastplanninggroupresponse{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Forecastplanninggroupsresponse) MarshalJSON() ([]byte, error) {
    type Alias Forecastplanninggroupsresponse

    if ForecastplanninggroupsresponseMarshalled {
        return []byte("{}"), nil
    }
    ForecastplanninggroupsresponseMarshalled = true

    return json.Marshal(&struct { 
        Entities []Forecastplanninggroupresponse `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Forecastplanninggroupresponse{{}},
        

        
        Alias: (*Alias)(u),
    })
}

