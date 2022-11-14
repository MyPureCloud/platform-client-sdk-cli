package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricalshrinkageactivitycategoryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricalshrinkageactivitycategoryresponseDud struct { 
    


    


    

}

// Historicalshrinkageactivitycategoryresponse
type Historicalshrinkageactivitycategoryresponse struct { 
    // ActivityCategory - Activity category for which shrinkage data is provided
    ActivityCategory string `json:"activityCategory"`


    // ShrinkageForActivityCategory - Aggregated shrinkage data for the activity category
    ShrinkageForActivityCategory Historicalshrinkageaggregateresponse `json:"shrinkageForActivityCategory"`


    // ShrinkageForActivityCodes - Shrinkage for the activity codes under this activity category
    ShrinkageForActivityCodes []Historicalshrinkageactivitycoderesponse `json:"shrinkageForActivityCodes"`

}

// String returns a JSON representation of the model
func (o *Historicalshrinkageactivitycategoryresponse) String() string {
    
    
     o.ShrinkageForActivityCodes = []Historicalshrinkageactivitycoderesponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicalshrinkageactivitycategoryresponse) MarshalJSON() ([]byte, error) {
    type Alias Historicalshrinkageactivitycategoryresponse

    if HistoricalshrinkageactivitycategoryresponseMarshalled {
        return []byte("{}"), nil
    }
    HistoricalshrinkageactivitycategoryresponseMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCategory string `json:"activityCategory"`
        
        ShrinkageForActivityCategory Historicalshrinkageaggregateresponse `json:"shrinkageForActivityCategory"`
        
        ShrinkageForActivityCodes []Historicalshrinkageactivitycoderesponse `json:"shrinkageForActivityCodes"`
        *Alias
    }{

        


        


        
        ShrinkageForActivityCodes: []Historicalshrinkageactivitycoderesponse{{}},
        

        Alias: (*Alias)(u),
    })
}

