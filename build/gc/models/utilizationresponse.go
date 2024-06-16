package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UtilizationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UtilizationresponseDud struct { 
    


    

}

// Utilizationresponse
type Utilizationresponse struct { 
    // Utilization - Map of media type to utilization settings.
    Utilization map[string]Mediautilization `json:"utilization"`


    // LabelUtilizations - Map of label ids to utilization settings.
    LabelUtilizations map[string]Labelutilizationresponse `json:"labelUtilizations"`

}

// String returns a JSON representation of the model
func (o *Utilizationresponse) String() string {
     o.Utilization = map[string]Mediautilization{"": {}} 
     o.LabelUtilizations = map[string]Labelutilizationresponse{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Utilizationresponse) MarshalJSON() ([]byte, error) {
    type Alias Utilizationresponse

    if UtilizationresponseMarshalled {
        return []byte("{}"), nil
    }
    UtilizationresponseMarshalled = true

    return json.Marshal(&struct {
        
        Utilization map[string]Mediautilization `json:"utilization"`
        
        LabelUtilizations map[string]Labelutilizationresponse `json:"labelUtilizations"`
        *Alias
    }{

        
        Utilization: map[string]Mediautilization{"": {}},
        


        
        LabelUtilizations: map[string]Labelutilizationresponse{"": {}},
        

        Alias: (*Alias)(u),
    })
}

