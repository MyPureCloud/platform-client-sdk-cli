package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuupdateagentscheduleuploadschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuupdateagentscheduleuploadschemaDud struct { 
    


    


    


    


    


    


    

}

// Buupdateagentscheduleuploadschema
type Buupdateagentscheduleuploadschema struct { 
    // UserId - The ID of the user to whom this agent schedule applies
    UserId string `json:"userId"`


    // WorkPlanId - The ID of the work plan for this user.  Mutually exclusive with workPlanIdsPerWeek
    WorkPlanId Valuewrapperstring `json:"workPlanId"`


    // WorkPlanIdsPerWeek - The IDs of the work plans per week for this user.  Mutually exclusive with workPlanId
    WorkPlanIdsPerWeek Listwrapperstring `json:"workPlanIdsPerWeek"`


    // Shifts - The shift definitions for this agent schedule
    Shifts []Buagentscheduleshift `json:"shifts"`


    // FullDayTimeOffMarkers - Any full day time off markers that apply to this agent schedule
    FullDayTimeOffMarkers []Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers"`


    // Metadata - Version metadata for this agent schedule. Required if updating or deleting an existing agent schedule, otherwise should be omitted
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    // Delete - Whether to delete this agent's schedule. Defaults to false if not set
    Delete bool `json:"delete"`

}

// String returns a JSON representation of the model
func (o *Buupdateagentscheduleuploadschema) String() string {
    
    
    
     o.Shifts = []Buagentscheduleshift{{}} 
     o.FullDayTimeOffMarkers = []Bufulldaytimeoffmarker{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buupdateagentscheduleuploadschema) MarshalJSON() ([]byte, error) {
    type Alias Buupdateagentscheduleuploadschema

    if BuupdateagentscheduleuploadschemaMarshalled {
        return []byte("{}"), nil
    }
    BuupdateagentscheduleuploadschemaMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        WorkPlanId Valuewrapperstring `json:"workPlanId"`
        
        WorkPlanIdsPerWeek Listwrapperstring `json:"workPlanIdsPerWeek"`
        
        Shifts []Buagentscheduleshift `json:"shifts"`
        
        FullDayTimeOffMarkers []Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        Delete bool `json:"delete"`
        *Alias
    }{

        


        


        


        
        Shifts: []Buagentscheduleshift{{}},
        


        
        FullDayTimeOffMarkers: []Bufulldaytimeoffmarker{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

