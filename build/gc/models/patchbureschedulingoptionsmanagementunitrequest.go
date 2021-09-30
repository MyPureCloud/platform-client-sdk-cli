package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchbureschedulingoptionsmanagementunitrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchbureschedulingoptionsmanagementunitrequestDud struct { 
    


    

}

// Patchbureschedulingoptionsmanagementunitrequest
type Patchbureschedulingoptionsmanagementunitrequest struct { 
    // ManagementUnitId - The management unit portion of the rescheduling run to update
    ManagementUnitId string `json:"managementUnitId"`


    // Applied - Whether to mark the run as applied.  Only applies to reschedule runs.  Once applied, a run cannot be un-marked as applied
    Applied bool `json:"applied"`

}

// String returns a JSON representation of the model
func (o *Patchbureschedulingoptionsmanagementunitrequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchbureschedulingoptionsmanagementunitrequest) MarshalJSON() ([]byte, error) {
    type Alias Patchbureschedulingoptionsmanagementunitrequest

    if PatchbureschedulingoptionsmanagementunitrequestMarshalled {
        return []byte("{}"), nil
    }
    PatchbureschedulingoptionsmanagementunitrequestMarshalled = true

    return json.Marshal(&struct { 
        ManagementUnitId string `json:"managementUnitId"`
        
        Applied bool `json:"applied"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

