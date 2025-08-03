package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SnapshotinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SnapshotinfoDud struct { 
    


    


    


    


    

}

// Snapshotinfo
type Snapshotinfo struct { 
    // Version - Version of the snapshot
    Version int `json:"version"`


    // SnapshotId - Snapshot Id of the continuous forecast session
    SnapshotId string `json:"snapshotId"`


    // DateSnapshot - Date of snapshot generation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateSnapshot time.Time `json:"dateSnapshot"`


    // BusinessUnitId - Business unit ID of the continuous forecast session
    BusinessUnitId string `json:"businessUnitId"`


    // PlanningGroupsVersion - Version of the planning groups
    PlanningGroupsVersion int `json:"planningGroupsVersion"`

}

// String returns a JSON representation of the model
func (o *Snapshotinfo) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Snapshotinfo) MarshalJSON() ([]byte, error) {
    type Alias Snapshotinfo

    if SnapshotinfoMarshalled {
        return []byte("{}"), nil
    }
    SnapshotinfoMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        
        SnapshotId string `json:"snapshotId"`
        
        DateSnapshot time.Time `json:"dateSnapshot"`
        
        BusinessUnitId string `json:"businessUnitId"`
        
        PlanningGroupsVersion int `json:"planningGroupsVersion"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

