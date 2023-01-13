package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherencebulkitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherencebulkitemDud struct { 
    


    


    


    


    

}

// Wfmhistoricaladherencebulkitem
type Wfmhistoricaladherencebulkitem struct { 
    // ManagementUnitId - The ID of the management unit to query
    ManagementUnitId string `json:"managementUnitId"`


    // StartDate - Beginning of the date range to query in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - End of the date range to query in ISO-8601 format
    EndDate time.Time `json:"endDate"`


    // UserIds - The IDs of the users to query. If not included, will query every user in the management unit
    UserIds []string `json:"userIds"`


    // IncludeExceptions - Whether user exceptions should be returned as part of the results. If not included, will default to false
    IncludeExceptions bool `json:"includeExceptions"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkitem) String() string {
    
    
    
     o.UserIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherencebulkitem) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherencebulkitem

    if WfmhistoricaladherencebulkitemMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherencebulkitemMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnitId string `json:"managementUnitId"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        UserIds []string `json:"userIds"`
        
        IncludeExceptions bool `json:"includeExceptions"`
        *Alias
    }{

        


        


        


        
        UserIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

