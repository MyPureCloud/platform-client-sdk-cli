package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightstrendsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightstrendsDud struct { 
    


    

}

// Insightstrends
type Insightstrends struct { 
    // ComparativePeriod - List of trend data in the comparative period
    ComparativePeriod []Trenddata `json:"comparativePeriod"`


    // PrimaryPeriod - List of trend data in the primary period
    PrimaryPeriod []Trenddata `json:"primaryPeriod"`

}

// String returns a JSON representation of the model
func (o *Insightstrends) String() string {
     o.ComparativePeriod = []Trenddata{{}} 
     o.PrimaryPeriod = []Trenddata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightstrends) MarshalJSON() ([]byte, error) {
    type Alias Insightstrends

    if InsightstrendsMarshalled {
        return []byte("{}"), nil
    }
    InsightstrendsMarshalled = true

    return json.Marshal(&struct {
        
        ComparativePeriod []Trenddata `json:"comparativePeriod"`
        
        PrimaryPeriod []Trenddata `json:"primaryPeriod"`
        *Alias
    }{

        
        ComparativePeriod: []Trenddata{{}},
        


        
        PrimaryPeriod: []Trenddata{{}},
        

        Alias: (*Alias)(u),
    })
}

