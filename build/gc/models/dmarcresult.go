package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DmarcresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DmarcresultDud struct { 
    Status string `json:"status"`


    DetectedPolicy string `json:"detectedPolicy"`


    DateChecked time.Time `json:"dateChecked"`


    Records []Record `json:"records"`

}

// Dmarcresult - Represents the DMARC verification result for an email domain
type Dmarcresult struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Dmarcresult) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dmarcresult) MarshalJSON() ([]byte, error) {
    type Alias Dmarcresult

    if DmarcresultMarshalled {
        return []byte("{}"), nil
    }
    DmarcresultMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

