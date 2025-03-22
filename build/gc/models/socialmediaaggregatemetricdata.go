package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediaaggregatemetricdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediaaggregatemetricdataDud struct { 
    


    


    

}

// Socialmediaaggregatemetricdata
type Socialmediaaggregatemetricdata struct { 
    // Metric
    Metric string `json:"metric"`


    // Qualifier
    Qualifier string `json:"qualifier"`


    // Stats
    Stats Socialmediastatisticalsummary `json:"stats"`

}

// String returns a JSON representation of the model
func (o *Socialmediaaggregatemetricdata) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediaaggregatemetricdata) MarshalJSON() ([]byte, error) {
    type Alias Socialmediaaggregatemetricdata

    if SocialmediaaggregatemetricdataMarshalled {
        return []byte("{}"), nil
    }
    SocialmediaaggregatemetricdataMarshalled = true

    return json.Marshal(&struct {
        
        Metric string `json:"metric"`
        
        Qualifier string `json:"qualifier"`
        
        Stats Socialmediastatisticalsummary `json:"stats"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

