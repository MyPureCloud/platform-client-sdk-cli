package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuimporttimeofflimitvaluesschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuimporttimeofflimitvaluesschemaDud struct { 
    


    

}

// Buimporttimeofflimitvaluesschema
type Buimporttimeofflimitvaluesschema struct { 
    // LimitValues - Time-off limit values to import. The list is collection of date and time interval for which allocated limit in minutes is imported.For a time-off limit with daily granularity, the only time interval that should be set for a given date is '00:00'For a time-off limit with fifteen minutes granularity, minimum of one time interval must be specified
    LimitValues []Buimporttimeofflimitvalue `json:"limitValues"`


    // Metadata - Version metadata for the time-off limit
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Buimporttimeofflimitvaluesschema) String() string {
     o.LimitValues = []Buimporttimeofflimitvalue{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buimporttimeofflimitvaluesschema) MarshalJSON() ([]byte, error) {
    type Alias Buimporttimeofflimitvaluesschema

    if BuimporttimeofflimitvaluesschemaMarshalled {
        return []byte("{}"), nil
    }
    BuimporttimeofflimitvaluesschemaMarshalled = true

    return json.Marshal(&struct {
        
        LimitValues []Buimporttimeofflimitvalue `json:"limitValues"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        
        LimitValues: []Buimporttimeofflimitvalue{{}},
        


        

        Alias: (*Alias)(u),
    })
}

