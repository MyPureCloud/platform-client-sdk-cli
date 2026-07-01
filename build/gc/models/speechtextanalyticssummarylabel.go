package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SpeechtextanalyticssummarylabelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SpeechtextanalyticssummarylabelDud struct { 
    


    


    

}

// Speechtextanalyticssummarylabel
type Speechtextanalyticssummarylabel struct { 
    // Name - The label name
    Name string `json:"name"`


    // Description - The label description
    Description string `json:"description"`


    // VarType - The type of the label
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Speechtextanalyticssummarylabel) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Speechtextanalyticssummarylabel) MarshalJSON() ([]byte, error) {
    type Alias Speechtextanalyticssummarylabel

    if SpeechtextanalyticssummarylabelMarshalled {
        return []byte("{}"), nil
    }
    SpeechtextanalyticssummarylabelMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

