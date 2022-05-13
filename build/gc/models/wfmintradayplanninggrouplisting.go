package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmintradayplanninggrouplistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmintradayplanninggrouplistingDud struct { 
    


    

}

// Wfmintradayplanninggrouplisting - A list of IntradayPlanningGroup objects
type Wfmintradayplanninggrouplisting struct { 
    // Entities
    Entities []Forecastplanninggroupresponse `json:"entities"`


    // NoDataReason - The reason there was no data for the request
    NoDataReason string `json:"noDataReason"`

}

// String returns a JSON representation of the model
func (o *Wfmintradayplanninggrouplisting) String() string {
     o.Entities = []Forecastplanninggroupresponse{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmintradayplanninggrouplisting) MarshalJSON() ([]byte, error) {
    type Alias Wfmintradayplanninggrouplisting

    if WfmintradayplanninggrouplistingMarshalled {
        return []byte("{}"), nil
    }
    WfmintradayplanninggrouplistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Forecastplanninggroupresponse `json:"entities"`
        
        NoDataReason string `json:"noDataReason"`
        *Alias
    }{

        
        Entities: []Forecastplanninggroupresponse{{}},
        


        

        Alias: (*Alias)(u),
    })
}

