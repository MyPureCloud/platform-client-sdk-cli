package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutlierinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutlierinfoDud struct { 
    


    

}

// Outlierinfo
type Outlierinfo struct { 
    // Outlier - Boolean to identify if an outlier or not.
    Outlier bool `json:"outlier"`


    // Score - Outlier score for this utterance. The score is always 0 or greater and higher the score, the more outlier.
    Score float32 `json:"score"`

}

// String returns a JSON representation of the model
func (o *Outlierinfo) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outlierinfo) MarshalJSON() ([]byte, error) {
    type Alias Outlierinfo

    if OutlierinfoMarshalled {
        return []byte("{}"), nil
    }
    OutlierinfoMarshalled = true

    return json.Marshal(&struct {
        
        Outlier bool `json:"outlier"`
        
        Score float32 `json:"score"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

