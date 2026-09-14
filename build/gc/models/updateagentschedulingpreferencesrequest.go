package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateagentschedulingpreferencesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateagentschedulingpreferencesrequestDud struct { 
    


    

}

// Updateagentschedulingpreferencesrequest
type Updateagentschedulingpreferencesrequest struct { 
    // PreferencesToAdd - Scheduling preferences to add
    PreferencesToAdd []Agentschedulingpreference `json:"preferencesToAdd"`


    // PreferenceIdsToRemove - IDs of scheduling preferences to remove
    PreferenceIdsToRemove []string `json:"preferenceIdsToRemove"`

}

// String returns a JSON representation of the model
func (o *Updateagentschedulingpreferencesrequest) String() string {
     o.PreferencesToAdd = []Agentschedulingpreference{{}} 
     o.PreferenceIdsToRemove = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateagentschedulingpreferencesrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateagentschedulingpreferencesrequest

    if UpdateagentschedulingpreferencesrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateagentschedulingpreferencesrequestMarshalled = true

    return json.Marshal(&struct {
        
        PreferencesToAdd []Agentschedulingpreference `json:"preferencesToAdd"`
        
        PreferenceIdsToRemove []string `json:"preferenceIdsToRemove"`
        *Alias
    }{

        
        PreferencesToAdd: []Agentschedulingpreference{{}},
        


        
        PreferenceIdsToRemove: []string{""},
        

        Alias: (*Alias)(u),
    })
}

