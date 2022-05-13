package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchactionmapschedulegroupsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchactionmapschedulegroupsDud struct { 
    


    

}

// Patchactionmapschedulegroups
type Patchactionmapschedulegroups struct { 
    // ActionMapScheduleGroup - The actions map's associated schedule group.
    ActionMapScheduleGroup Actionmapschedulegroup `json:"actionMapScheduleGroup"`


    // EmergencyActionMapScheduleGroup - The action map's associated emergency schedule group.
    EmergencyActionMapScheduleGroup Actionmapschedulegroup `json:"emergencyActionMapScheduleGroup"`

}

// String returns a JSON representation of the model
func (o *Patchactionmapschedulegroups) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchactionmapschedulegroups) MarshalJSON() ([]byte, error) {
    type Alias Patchactionmapschedulegroups

    if PatchactionmapschedulegroupsMarshalled {
        return []byte("{}"), nil
    }
    PatchactionmapschedulegroupsMarshalled = true

    return json.Marshal(&struct {
        
        ActionMapScheduleGroup Actionmapschedulegroup `json:"actionMapScheduleGroup"`
        
        EmergencyActionMapScheduleGroup Actionmapschedulegroup `json:"emergencyActionMapScheduleGroup"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

