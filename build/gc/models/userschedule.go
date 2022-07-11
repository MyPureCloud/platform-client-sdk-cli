package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserscheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserscheduleDud struct { 
    


    


    


    


    WorkPlanId string `json:"workPlanId"`

}

// Userschedule
type Userschedule struct { 
    // Shifts - The shifts that belong to this schedule
    Shifts []Userscheduleshift `json:"shifts"`


    // FullDayTimeOffMarkers - Markers to indicate a full day time off request, relative to the management unit time zone
    FullDayTimeOffMarkers []Userschedulefulldaytimeoffmarker `json:"fullDayTimeOffMarkers"`


    // Delete - If marked true for updating an existing user schedule, it will be deleted
    Delete bool `json:"delete"`


    // Metadata - Version metadata for this schedule
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Userschedule) String() string {
     o.Shifts = []Userscheduleshift{{}} 
     o.FullDayTimeOffMarkers = []Userschedulefulldaytimeoffmarker{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userschedule) MarshalJSON() ([]byte, error) {
    type Alias Userschedule

    if UserscheduleMarshalled {
        return []byte("{}"), nil
    }
    UserscheduleMarshalled = true

    return json.Marshal(&struct {
        
        Shifts []Userscheduleshift `json:"shifts"`
        
        FullDayTimeOffMarkers []Userschedulefulldaytimeoffmarker `json:"fullDayTimeOffMarkers"`
        
        Delete bool `json:"delete"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        
        Shifts: []Userscheduleshift{{}},
        


        
        FullDayTimeOffMarkers: []Userschedulefulldaytimeoffmarker{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

