package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherencebulkresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherencebulkresultDud struct { 
    


    


    


    


    

}

// Wfmhistoricaladherencebulkresult
type Wfmhistoricaladherencebulkresult struct { 
    // StartDate - Beginning of the date range for this result in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - End of the date range for this result in ISO-8601 format
    EndDate time.Time `json:"endDate"`


    // ManagementUnitId - The ID of the management unit for this result
    ManagementUnitId string `json:"managementUnitId"`


    // UserResults - The individual results for each user
    UserResults []Wfmhistoricaladherencebulkuserresult `json:"userResults"`


    // LookupIdToSecondaryPresenceId - Map of secondary presence lookup ID to corresponding secondary presence ID
    LookupIdToSecondaryPresenceId map[string]string `json:"lookupIdToSecondaryPresenceId"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkresult) String() string {
    
    
    
     o.UserResults = []Wfmhistoricaladherencebulkuserresult{{}} 
     o.LookupIdToSecondaryPresenceId = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherencebulkresult) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherencebulkresult

    if WfmhistoricaladherencebulkresultMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherencebulkresultMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        ManagementUnitId string `json:"managementUnitId"`
        
        UserResults []Wfmhistoricaladherencebulkuserresult `json:"userResults"`
        
        LookupIdToSecondaryPresenceId map[string]string `json:"lookupIdToSecondaryPresenceId"`
        *Alias
    }{

        


        


        


        
        UserResults: []Wfmhistoricaladherencebulkuserresult{{}},
        


        
        LookupIdToSecondaryPresenceId: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

