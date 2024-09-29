package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillcombinationinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillcombinationinfoDud struct { 
    


    TotalCount int `json:"totalCount"`


    RemainingCount int `json:"remainingCount"`

}

// Skillcombinationinfo
type Skillcombinationinfo struct { 
    // SkillCombination - A skill combination in the contact queue
    SkillCombination []string `json:"skillCombination"`


    


    

}

// String returns a JSON representation of the model
func (o *Skillcombinationinfo) String() string {
     o.SkillCombination = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillcombinationinfo) MarshalJSON() ([]byte, error) {
    type Alias Skillcombinationinfo

    if SkillcombinationinfoMarshalled {
        return []byte("{}"), nil
    }
    SkillcombinationinfoMarshalled = true

    return json.Marshal(&struct {
        
        SkillCombination []string `json:"skillCombination"`
        *Alias
    }{

        
        SkillCombination: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

