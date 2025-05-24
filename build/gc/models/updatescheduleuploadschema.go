package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatescheduleuploadschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatescheduleuploadschemaDud struct { 
    


    


    


    


    


    

}

// Updatescheduleuploadschema
type Updatescheduleuploadschema struct { 
    // Description - The description to set for the schedule
    Description string `json:"description"`


    // Published - Whether to publish the schedule. Note: a schedule cannot be un-published unless another schedule is published over it
    Published bool `json:"published"`


    // ShortTermForecast - The short term forecast to associate with the schedule
    ShortTermForecast Bushorttermforecastreference `json:"shortTermForecast"`


    // HeadcountForecast - The headcount forecast to associate with the schedule
    HeadcountForecast Buheadcountforecastbuplanninggroupheadcountforecastuploadschema `json:"headcountForecast"`


    // AgentSchedules - Individual agent schedules
    AgentSchedules []Buupdateagentscheduleuploadschema `json:"agentSchedules"`


    // Metadata - Version metadata for this schedule
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updatescheduleuploadschema) String() string {
    
    
    
    
     o.AgentSchedules = []Buupdateagentscheduleuploadschema{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatescheduleuploadschema) MarshalJSON() ([]byte, error) {
    type Alias Updatescheduleuploadschema

    if UpdatescheduleuploadschemaMarshalled {
        return []byte("{}"), nil
    }
    UpdatescheduleuploadschemaMarshalled = true

    return json.Marshal(&struct {
        
        Description string `json:"description"`
        
        Published bool `json:"published"`
        
        ShortTermForecast Bushorttermforecastreference `json:"shortTermForecast"`
        
        HeadcountForecast Buheadcountforecastbuplanninggroupheadcountforecastuploadschema `json:"headcountForecast"`
        
        AgentSchedules []Buupdateagentscheduleuploadschema `json:"agentSchedules"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        
        AgentSchedules: []Buupdateagentscheduleuploadschema{{}},
        


        

        Alias: (*Alias)(u),
    })
}

