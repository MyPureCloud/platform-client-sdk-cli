package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulksourceintentsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulksourceintentsresponseDud struct { 
    


    


    

}

// Bulksourceintentsresponse
type Bulksourceintentsresponse struct { 
    // Results - Results of operation
    Results []Bulkresults `json:"results"`


    // ErrorCount - Count of errors
    ErrorCount int `json:"errorCount"`


    // ErrorResultIds - List of ids failed to be added or removed
    ErrorResultIds []string `json:"errorResultIds"`

}

// String returns a JSON representation of the model
func (o *Bulksourceintentsresponse) String() string {
     o.Results = []Bulkresults{{}} 
    
     o.ErrorResultIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulksourceintentsresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulksourceintentsresponse

    if BulksourceintentsresponseMarshalled {
        return []byte("{}"), nil
    }
    BulksourceintentsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkresults `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        
        ErrorResultIds []string `json:"errorResultIds"`
        *Alias
    }{

        
        Results: []Bulkresults{{}},
        


        


        
        ErrorResultIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

