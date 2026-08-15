package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AggregatedexportjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AggregatedexportjobrequestDud struct { 
    


    


    


    


    


    

}

// Aggregatedexportjobrequest
type Aggregatedexportjobrequest struct { 
    // Granularity - Granularity of the exported data, defaults to day
    Granularity string `json:"granularity"`


    // TimeZone - The requested time zone of the exported data, in Olson format. Defaults to business unit time zone
    TimeZone string `json:"timeZone"`


    // Delimiter - The delimiter to use between fields in the export, defaults to comma
    Delimiter string `json:"delimiter"`


    // PlanningGroupIds - The IDs of the planning groups to include in the export, defaults to all planning groups in the business unit
    PlanningGroupIds []string `json:"planningGroupIds"`


    // DateStart - Start date-time of the export range in ISO-8601 format
    DateStart time.Time `json:"dateStart"`


    // DateEnd - End date-time of the export range in ISO-8601 format
    DateEnd time.Time `json:"dateEnd"`

}

// String returns a JSON representation of the model
func (o *Aggregatedexportjobrequest) String() string {
    
    
    
     o.PlanningGroupIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aggregatedexportjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Aggregatedexportjobrequest

    if AggregatedexportjobrequestMarshalled {
        return []byte("{}"), nil
    }
    AggregatedexportjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        Granularity string `json:"granularity"`
        
        TimeZone string `json:"timeZone"`
        
        Delimiter string `json:"delimiter"`
        
        PlanningGroupIds []string `json:"planningGroupIds"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        *Alias
    }{

        


        


        


        
        PlanningGroupIds: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

