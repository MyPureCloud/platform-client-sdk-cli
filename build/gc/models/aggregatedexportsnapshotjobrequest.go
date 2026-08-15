package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AggregatedexportsnapshotjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AggregatedexportsnapshotjobrequestDud struct { 
    


    


    


    


    


    


    

}

// Aggregatedexportsnapshotjobrequest
type Aggregatedexportsnapshotjobrequest struct { 
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


    // SnapshotId - The ID of the snapshot to export
    SnapshotId string `json:"snapshotId"`

}

// String returns a JSON representation of the model
func (o *Aggregatedexportsnapshotjobrequest) String() string {
    
    
    
     o.PlanningGroupIds = []string{""} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aggregatedexportsnapshotjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Aggregatedexportsnapshotjobrequest

    if AggregatedexportsnapshotjobrequestMarshalled {
        return []byte("{}"), nil
    }
    AggregatedexportsnapshotjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        Granularity string `json:"granularity"`
        
        TimeZone string `json:"timeZone"`
        
        Delimiter string `json:"delimiter"`
        
        PlanningGroupIds []string `json:"planningGroupIds"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        SnapshotId string `json:"snapshotId"`
        *Alias
    }{

        


        


        


        
        PlanningGroupIds: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

