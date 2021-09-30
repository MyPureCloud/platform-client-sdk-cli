package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DevelopmentactivityaggregatequeryresponsegroupeddataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DevelopmentactivityaggregatequeryresponsegroupeddataDud struct { 
    


    

}

// Developmentactivityaggregatequeryresponsegroupeddata
type Developmentactivityaggregatequeryresponsegroupeddata struct { 
    // Group - The group values for this data
    Group map[string]string `json:"group"`


    // Data - The metrics in this group
    Data []Developmentactivityaggregatequeryresponsedata `json:"data"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsegroupeddata) String() string {
    
    
     o.Group = map[string]string{"": ""} 
    
    
    
     o.Data = []Developmentactivityaggregatequeryresponsedata{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Developmentactivityaggregatequeryresponsegroupeddata) MarshalJSON() ([]byte, error) {
    type Alias Developmentactivityaggregatequeryresponsegroupeddata

    if DevelopmentactivityaggregatequeryresponsegroupeddataMarshalled {
        return []byte("{}"), nil
    }
    DevelopmentactivityaggregatequeryresponsegroupeddataMarshalled = true

    return json.Marshal(&struct { 
        Group map[string]string `json:"group"`
        
        Data []Developmentactivityaggregatequeryresponsedata `json:"data"`
        
        *Alias
    }{
        

        
        Group: map[string]string{"": ""},
        

        

        
        Data: []Developmentactivityaggregatequeryresponsedata{{}},
        

        
        Alias: (*Alias)(u),
    })
}

