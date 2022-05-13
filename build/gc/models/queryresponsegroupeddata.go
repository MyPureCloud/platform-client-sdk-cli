package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryresponsegroupeddataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryresponsegroupeddataDud struct { 
    


    

}

// Queryresponsegroupeddata
type Queryresponsegroupeddata struct { 
    // Group - The group values for this data
    Group map[string]string `json:"group"`


    // Data - The metrics in this group
    Data []Queryresponsedata `json:"data"`

}

// String returns a JSON representation of the model
func (o *Queryresponsegroupeddata) String() string {
     o.Group = map[string]string{"": ""} 
     o.Data = []Queryresponsedata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryresponsegroupeddata) MarshalJSON() ([]byte, error) {
    type Alias Queryresponsegroupeddata

    if QueryresponsegroupeddataMarshalled {
        return []byte("{}"), nil
    }
    QueryresponsegroupeddataMarshalled = true

    return json.Marshal(&struct {
        
        Group map[string]string `json:"group"`
        
        Data []Queryresponsedata `json:"data"`
        *Alias
    }{

        
        Group: map[string]string{"": ""},
        


        
        Data: []Queryresponsedata{{}},
        

        Alias: (*Alias)(u),
    })
}

