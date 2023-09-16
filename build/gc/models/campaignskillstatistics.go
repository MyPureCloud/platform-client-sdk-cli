package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignskillstatisticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignskillstatisticsDud struct { 
    SkillCombinations int `json:"skillCombinations"`


    EligibleSkilledAgents int `json:"eligibleSkilledAgents"`

}

// Campaignskillstatistics
type Campaignskillstatistics struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Campaignskillstatistics) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignskillstatistics) MarshalJSON() ([]byte, error) {
    type Alias Campaignskillstatistics

    if CampaignskillstatisticsMarshalled {
        return []byte("{}"), nil
    }
    CampaignskillstatisticsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

