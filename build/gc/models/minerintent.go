package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MinerintentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MinerintentDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Minerintent
type Minerintent struct { 
    


    // Name - Intent name.
    Name string `json:"name"`


    // Miner - The miner to which the intent belongs.
    Miner Miner `json:"miner"`


    // Utterances - The utterances that are extracted for an Intent.
    Utterances []Utterance `json:"utterances"`


    // AnalyticVolumePercent - Percentage of conversations belonging to the intent.
    AnalyticVolumePercent float64 `json:"analyticVolumePercent"`


    

}

// String returns a JSON representation of the model
func (o *Minerintent) String() string {
    
    
     o.Utterances = []Utterance{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Minerintent) MarshalJSON() ([]byte, error) {
    type Alias Minerintent

    if MinerintentMarshalled {
        return []byte("{}"), nil
    }
    MinerintentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Miner Miner `json:"miner"`
        
        Utterances []Utterance `json:"utterances"`
        
        AnalyticVolumePercent float64 `json:"analyticVolumePercent"`
        *Alias
    }{

        


        


        


        
        Utterances: []Utterance{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

