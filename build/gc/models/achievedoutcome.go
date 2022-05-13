package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AchievedoutcomeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AchievedoutcomeDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Achievedoutcome
type Achievedoutcome struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Achievedoutcome) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Achievedoutcome) MarshalJSON() ([]byte, error) {
    type Alias Achievedoutcome

    if AchievedoutcomeMarshalled {
        return []byte("{}"), nil
    }
    AchievedoutcomeMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

