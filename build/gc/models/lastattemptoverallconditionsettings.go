package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LastattemptoverallconditionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LastattemptoverallconditionsettingsDud struct { 
    


    


    

}

// Lastattemptoverallconditionsettings
type Lastattemptoverallconditionsettings struct { 
    // MediaTypes - A list of media types to evaluate.
    MediaTypes []string `json:"mediaTypes"`


    // Operator - The operator to use when comparing values.
    Operator string `json:"operator"`


    // Value - The period value to compare against the contact's data.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Lastattemptoverallconditionsettings) String() string {
     o.MediaTypes = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lastattemptoverallconditionsettings) MarshalJSON() ([]byte, error) {
    type Alias Lastattemptoverallconditionsettings

    if LastattemptoverallconditionsettingsMarshalled {
        return []byte("{}"), nil
    }
    LastattemptoverallconditionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        MediaTypes []string `json:"mediaTypes"`
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        *Alias
    }{

        
        MediaTypes: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

