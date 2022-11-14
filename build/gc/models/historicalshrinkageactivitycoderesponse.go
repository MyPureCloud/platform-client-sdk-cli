package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricalshrinkageactivitycoderesponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricalshrinkageactivitycoderesponseDud struct { 
    


    

}

// Historicalshrinkageactivitycoderesponse
type Historicalshrinkageactivitycoderesponse struct { 
    // ActivityCodeId - The ID of the activity code for which shrinkage data is provided
    ActivityCodeId string `json:"activityCodeId"`


    // ShrinkageForActivityCode - Aggregated shrinkage data for the activity code
    ShrinkageForActivityCode Historicalshrinkageaggregateresponse `json:"shrinkageForActivityCode"`

}

// String returns a JSON representation of the model
func (o *Historicalshrinkageactivitycoderesponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicalshrinkageactivitycoderesponse) MarshalJSON() ([]byte, error) {
    type Alias Historicalshrinkageactivitycoderesponse

    if HistoricalshrinkageactivitycoderesponseMarshalled {
        return []byte("{}"), nil
    }
    HistoricalshrinkageactivitycoderesponseMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCodeId string `json:"activityCodeId"`
        
        ShrinkageForActivityCode Historicalshrinkageaggregateresponse `json:"shrinkageForActivityCode"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

