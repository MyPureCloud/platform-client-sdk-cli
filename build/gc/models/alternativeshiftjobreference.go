package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshiftjobreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshiftjobreferenceDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Alternativeshiftjobreference
type Alternativeshiftjobreference struct { 
    


    // Status - The status of the alternative shift job
    Status string `json:"status"`


    // VarType - The type of alternative shift asynchronous job
    VarType string `json:"type"`


    

}

// String returns a JSON representation of the model
func (o *Alternativeshiftjobreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshiftjobreference) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshiftjobreference

    if AlternativeshiftjobreferenceMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshiftjobreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

