package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlterationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlterationDud struct { 
    


    


    

}

// Alteration
type Alteration struct { 
    // VarType - Range type (Exclusion: used to exclude a specific time within the recurrence. Inclusion: used to include a specific time within the recurrence which will execute in addition to the normal recurrence. If both an exclusion and inclusion are specified, the inclusion will take precedence over the exclusion.)
    VarType string `json:"type"`


    // Start - The start date of an alteration range as an ISO-8601 string in UTC time, e.g: 2023-12-21T16:30:25.000Z
    Start string `json:"start"`


    // End - The end date of an alteration range as an ISO-8601 string in UTC time, e.g: 2023-12-21T18:30:25.000Z
    End string `json:"end"`

}

// String returns a JSON representation of the model
func (o *Alteration) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alteration) MarshalJSON() ([]byte, error) {
    type Alias Alteration

    if AlterationMarshalled {
        return []byte("{}"), nil
    }
    AlterationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Start string `json:"start"`
        
        End string `json:"end"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

