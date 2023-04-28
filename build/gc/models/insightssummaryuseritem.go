package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightssummaryuseritemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightssummaryuseritemDud struct { 
    


    


    

}

// Insightssummaryuseritem
type Insightssummaryuseritem struct { 
    // User - Queried user
    User Userreference `json:"user"`


    // MetricData - The list of insights data for each metric of the user
    MetricData []Insightssummarymetricitem `json:"metricData"`


    // OverallData - Overall insights data of the user
    OverallData Insightssummaryoverallitem `json:"overallData"`

}

// String returns a JSON representation of the model
func (o *Insightssummaryuseritem) String() string {
    
     o.MetricData = []Insightssummarymetricitem{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightssummaryuseritem) MarshalJSON() ([]byte, error) {
    type Alias Insightssummaryuseritem

    if InsightssummaryuseritemMarshalled {
        return []byte("{}"), nil
    }
    InsightssummaryuseritemMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        MetricData []Insightssummarymetricitem `json:"metricData"`
        
        OverallData Insightssummaryoverallitem `json:"overallData"`
        *Alias
    }{

        


        
        MetricData: []Insightssummarymetricitem{{}},
        


        

        Alias: (*Alias)(u),
    })
}

