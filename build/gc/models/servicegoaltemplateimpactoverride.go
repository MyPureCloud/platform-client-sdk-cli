package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ServicegoaltemplateimpactoverrideMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ServicegoaltemplateimpactoverrideDud struct { 
    


    

}

// Servicegoaltemplateimpactoverride
type Servicegoaltemplateimpactoverride struct { 
    // Enabled - Whether service goal overrides are enabled for this service goal template
    Enabled bool `json:"enabled"`


    // Impact - Settings controlling max percent increase and decrease of service goals for this service goal template
    Impact Wfmservicegoalimpactsettings `json:"impact"`

}

// String returns a JSON representation of the model
func (o *Servicegoaltemplateimpactoverride) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Servicegoaltemplateimpactoverride) MarshalJSON() ([]byte, error) {
    type Alias Servicegoaltemplateimpactoverride

    if ServicegoaltemplateimpactoverrideMarshalled {
        return []byte("{}"), nil
    }
    ServicegoaltemplateimpactoverrideMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        Impact Wfmservicegoalimpactsettings `json:"impact"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

