package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalydetectionqueueanalysisresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalydetectionqueueanalysisresponseDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Anomalydetectionqueueanalysisresponse
type Anomalydetectionqueueanalysisresponse struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Analyses - Anomaly analysis for each signal
    Analyses []Anomalydetectionqueueanalysis `json:"analyses"`


    

}

// String returns a JSON representation of the model
func (o *Anomalydetectionqueueanalysisresponse) String() string {
    
     o.Analyses = []Anomalydetectionqueueanalysis{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalydetectionqueueanalysisresponse) MarshalJSON() ([]byte, error) {
    type Alias Anomalydetectionqueueanalysisresponse

    if AnomalydetectionqueueanalysisresponseMarshalled {
        return []byte("{}"), nil
    }
    AnomalydetectionqueueanalysisresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Analyses []Anomalydetectionqueueanalysis `json:"analyses"`
        *Alias
    }{

        


        
        Analyses: []Anomalydetectionqueueanalysis{{}},
        


        

        Alias: (*Alias)(u),
    })
}

