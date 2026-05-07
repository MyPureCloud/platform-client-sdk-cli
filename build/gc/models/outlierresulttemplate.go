package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutlierresulttemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutlierresulttemplateDud struct { 
    

}

// Outlierresulttemplate
type Outlierresulttemplate struct { 
    // Entities - Outliers result bodies for planning groups
    Entities []Outlierresultbody `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Outlierresulttemplate) String() string {
     o.Entities = []Outlierresultbody{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outlierresulttemplate) MarshalJSON() ([]byte, error) {
    type Alias Outlierresulttemplate

    if OutlierresulttemplateMarshalled {
        return []byte("{}"), nil
    }
    OutlierresulttemplateMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Outlierresultbody `json:"entities"`
        *Alias
    }{

        
        Entities: []Outlierresultbody{{}},
        

        Alias: (*Alias)(u),
    })
}

