package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButimeofflimitresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButimeofflimitresponseDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Butimeofflimitresponse
type Butimeofflimitresponse struct { 
    


    // StaffingGroup - The staffing group to which this time-off limit is associated. If managementUnit is set, then the staffing group belongs to that management unit.Otherwise, if managementUnit is not set, it is a business unit level staffing group.At least one of managementUnit and staffingGroup must be set
    StaffingGroup Staffinggroupreference `json:"staffingGroup"`


    // ManagementUnit - The management unit to which this time-off limit is associated. If staffingGroup is set, then the limit is associated with that staffing group, which belongs to this management unit.At least one of managementUnit and staffingGroup must be set
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // Metadata - Version metadata for the time-off limit
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Butimeofflimitresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Butimeofflimitresponse) MarshalJSON() ([]byte, error) {
    type Alias Butimeofflimitresponse

    if ButimeofflimitresponseMarshalled {
        return []byte("{}"), nil
    }
    ButimeofflimitresponseMarshalled = true

    return json.Marshal(&struct {
        
        StaffingGroup Staffinggroupreference `json:"staffingGroup"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

