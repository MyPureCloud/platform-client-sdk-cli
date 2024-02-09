package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActioneventactionmapMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActioneventactionmapDud struct { 
    


    


    


    

}

// Actioneventactionmap
type Actioneventactionmap struct { 
    // Id - The ID of the action map.
    Id string `json:"id"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // Version - The version of the action map.
    Version int `json:"version"`


    // DisplayName - Display name of the action map.
    DisplayName string `json:"displayName"`

}

// String returns a JSON representation of the model
func (o *Actioneventactionmap) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actioneventactionmap) MarshalJSON() ([]byte, error) {
    type Alias Actioneventactionmap

    if ActioneventactionmapMarshalled {
        return []byte("{}"), nil
    }
    ActioneventactionmapMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SelfUri string `json:"selfUri"`
        
        Version int `json:"version"`
        
        DisplayName string `json:"displayName"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

