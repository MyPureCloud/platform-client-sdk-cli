package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceexplanationjobreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceexplanationjobreferenceDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Adherenceexplanationjobreference
type Adherenceexplanationjobreference struct { 
    


    // VarType - The type of the adherence explanation job
    VarType string `json:"type"`


    // Status - The status of the adherence explanation job
    Status string `json:"status"`


    

}

// String returns a JSON representation of the model
func (o *Adherenceexplanationjobreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceexplanationjobreference) MarshalJSON() ([]byte, error) {
    type Alias Adherenceexplanationjobreference

    if AdherenceexplanationjobreferenceMarshalled {
        return []byte("{}"), nil
    }
    AdherenceexplanationjobreferenceMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

