package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PerformancepredictionuploadschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PerformancepredictionuploadschemaDud struct { 
    


    

}

// Performancepredictionuploadschema
type Performancepredictionuploadschema struct { 
    // CalculationStartDate - Date as an ISO-8601 string, corresponding to the beginning of the performance prediction results
    CalculationStartDate time.Time `json:"calculationStartDate"`


    // OnQueueTimes - List of agent on-queue times by management unit
    OnQueueTimes []Muagentqueuetimerequest `json:"onQueueTimes"`

}

// String returns a JSON representation of the model
func (o *Performancepredictionuploadschema) String() string {
    
     o.OnQueueTimes = []Muagentqueuetimerequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Performancepredictionuploadschema) MarshalJSON() ([]byte, error) {
    type Alias Performancepredictionuploadschema

    if PerformancepredictionuploadschemaMarshalled {
        return []byte("{}"), nil
    }
    PerformancepredictionuploadschemaMarshalled = true

    return json.Marshal(&struct {
        
        CalculationStartDate time.Time `json:"calculationStartDate"`
        
        OnQueueTimes []Muagentqueuetimerequest `json:"onQueueTimes"`
        *Alias
    }{

        


        
        OnQueueTimes: []Muagentqueuetimerequest{{}},
        

        Alias: (*Alias)(u),
    })
}

