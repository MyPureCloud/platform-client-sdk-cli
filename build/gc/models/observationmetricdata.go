package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ObservationmetricdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ObservationmetricdataDud struct { 
    


    


    


    


    

}

// Observationmetricdata
type Observationmetricdata struct { 
    // Metric
    Metric string `json:"metric"`


    // Qualifier
    Qualifier string `json:"qualifier"`


    // Stats
    Stats Statisticalsummary `json:"stats"`


    // Truncated - Flag for a truncated list of observations. If truncated, the first half of the list of observations will contain the oldest observations and the second half the newest observations.
    Truncated bool `json:"truncated"`


    // Observations - List of observations sorted by timestamp in ascending order. This list may be truncated.
    Observations []Observationvalue `json:"observations"`

}

// String returns a JSON representation of the model
func (o *Observationmetricdata) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Observations = []Observationvalue{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Observationmetricdata) MarshalJSON() ([]byte, error) {
    type Alias Observationmetricdata

    if ObservationmetricdataMarshalled {
        return []byte("{}"), nil
    }
    ObservationmetricdataMarshalled = true

    return json.Marshal(&struct { 
        Metric string `json:"metric"`
        
        Qualifier string `json:"qualifier"`
        
        Stats Statisticalsummary `json:"stats"`
        
        Truncated bool `json:"truncated"`
        
        Observations []Observationvalue `json:"observations"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Observations: []Observationvalue{{}},
        

        
        Alias: (*Alias)(u),
    })
}

