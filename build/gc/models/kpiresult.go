package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KpiresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KpiresultDud struct { 
    KpiTotalOn int `json:"kpiTotalOn"`


    KpiTotalOff int `json:"kpiTotalOff"`


    InteractionCountOn int `json:"interactionCountOn"`


    InteractionCountOff int `json:"interactionCountOff"`


    MediaType string `json:"mediaType"`

}

// Kpiresult
type Kpiresult struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Kpiresult) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Kpiresult) MarshalJSON() ([]byte, error) {
    type Alias Kpiresult

    if KpiresultMarshalled {
        return []byte("{}"), nil
    }
    KpiresultMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

