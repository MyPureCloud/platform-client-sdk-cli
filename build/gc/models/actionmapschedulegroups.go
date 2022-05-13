package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionmapschedulegroupsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionmapschedulegroupsDud struct { 
    


    

}

// Actionmapschedulegroups
type Actionmapschedulegroups struct { 
    // ActionMapScheduleGroup - The actions map's associated schedule group.
    ActionMapScheduleGroup Actionmapschedulegroup `json:"actionMapScheduleGroup"`


    // EmergencyActionMapScheduleGroup - The action map's associated emergency schedule group.
    EmergencyActionMapScheduleGroup Actionmapschedulegroup `json:"emergencyActionMapScheduleGroup"`

}

// String returns a JSON representation of the model
func (o *Actionmapschedulegroups) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionmapschedulegroups) MarshalJSON() ([]byte, error) {
    type Alias Actionmapschedulegroups

    if ActionmapschedulegroupsMarshalled {
        return []byte("{}"), nil
    }
    ActionmapschedulegroupsMarshalled = true

    return json.Marshal(&struct {
        
        ActionMapScheduleGroup Actionmapschedulegroup `json:"actionMapScheduleGroup"`
        
        EmergencyActionMapScheduleGroup Actionmapschedulegroup `json:"emergencyActionMapScheduleGroup"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

