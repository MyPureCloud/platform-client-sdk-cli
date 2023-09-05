package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulingsettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulingsettingsresponseDud struct { 
    


    


    

}

// Buschedulingsettingsresponse
type Buschedulingsettingsresponse struct { 
    // MessageSeverities - Schedule generation message severity configuration
    MessageSeverities []Schedulermessagetypeseverity `json:"messageSeverities"`


    // SyncTimeOffProperties - Synchronize set of time off properties from scheduled activities to time off requests when the schedule is published.
    SyncTimeOffProperties []string `json:"syncTimeOffProperties"`


    // ServiceGoalImpact - Configures the max percent increase and decrease of service goals for this business unit
    ServiceGoalImpact Wfmservicegoalimpactsettings `json:"serviceGoalImpact"`

}

// String returns a JSON representation of the model
func (o *Buschedulingsettingsresponse) String() string {
     o.MessageSeverities = []Schedulermessagetypeseverity{{}} 
     o.SyncTimeOffProperties = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulingsettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Buschedulingsettingsresponse

    if BuschedulingsettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    BuschedulingsettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        MessageSeverities []Schedulermessagetypeseverity `json:"messageSeverities"`
        
        SyncTimeOffProperties []string `json:"syncTimeOffProperties"`
        
        ServiceGoalImpact Wfmservicegoalimpactsettings `json:"serviceGoalImpact"`
        *Alias
    }{

        
        MessageSeverities: []Schedulermessagetypeseverity{{}},
        


        
        SyncTimeOffProperties: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

