package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentaggregatequeryresponsegroupeddataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentaggregatequeryresponsegroupeddataDud struct { 
    


    

}

// Learningassignmentaggregatequeryresponsegroupeddata
type Learningassignmentaggregatequeryresponsegroupeddata struct { 
    // Group - The group values for this data
    Group map[string]string `json:"group"`


    // Data - The metrics in this group
    Data []Learningassignmentaggregatequeryresponsedata `json:"data"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryresponsegroupeddata) String() string {
    
    
     o.Group = map[string]string{"": ""} 
    
    
    
     o.Data = []Learningassignmentaggregatequeryresponsedata{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentaggregatequeryresponsegroupeddata) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentaggregatequeryresponsegroupeddata

    if LearningassignmentaggregatequeryresponsegroupeddataMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentaggregatequeryresponsegroupeddataMarshalled = true

    return json.Marshal(&struct { 
        Group map[string]string `json:"group"`
        
        Data []Learningassignmentaggregatequeryresponsedata `json:"data"`
        
        *Alias
    }{
        

        
        Group: map[string]string{"": ""},
        

        

        
        Data: []Learningassignmentaggregatequeryresponsedata{{}},
        

        
        Alias: (*Alias)(u),
    })
}

