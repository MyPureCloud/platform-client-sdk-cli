package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PrograminsightssettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PrograminsightssettingsDud struct { 
    


    

}

// Programinsightssettings
type Programinsightssettings struct { 
    // Program - The ID of the program
    Program Baseprogramentity `json:"program"`


    // Enabled - The program AI Insights settings
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Programinsightssettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Programinsightssettings) MarshalJSON() ([]byte, error) {
    type Alias Programinsightssettings

    if PrograminsightssettingsMarshalled {
        return []byte("{}"), nil
    }
    PrograminsightssettingsMarshalled = true

    return json.Marshal(&struct {
        
        Program Baseprogramentity `json:"program"`
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

