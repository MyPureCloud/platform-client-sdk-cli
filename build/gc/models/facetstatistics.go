package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacetstatisticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacetstatisticsDud struct { 
    


    


    


    


    


    


    

}

// Facetstatistics
type Facetstatistics struct { 
    // Count
    Count int `json:"count"`


    // Min
    Min float64 `json:"min"`


    // Max
    Max float64 `json:"max"`


    // Mean
    Mean float64 `json:"mean"`


    // StdDeviation
    StdDeviation float64 `json:"stdDeviation"`


    // DateMin - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateMin time.Time `json:"dateMin"`


    // DateMax - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateMax time.Time `json:"dateMax"`

}

// String returns a JSON representation of the model
func (o *Facetstatistics) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facetstatistics) MarshalJSON() ([]byte, error) {
    type Alias Facetstatistics

    if FacetstatisticsMarshalled {
        return []byte("{}"), nil
    }
    FacetstatisticsMarshalled = true

    return json.Marshal(&struct {
        
        Count int `json:"count"`
        
        Min float64 `json:"min"`
        
        Max float64 `json:"max"`
        
        Mean float64 `json:"mean"`
        
        StdDeviation float64 `json:"stdDeviation"`
        
        DateMin time.Time `json:"dateMin"`
        
        DateMax time.Time `json:"dateMax"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

