package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuadherenceadjustmentssettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuadherenceadjustmentssettingsDud struct { 
    


    

}

// Buadherenceadjustmentssettings
type Buadherenceadjustmentssettings struct { 
    // SubmissionRangeConstraintDays - The maximum number of days in the past that an adherence adjustment can be submitted
    SubmissionRangeConstraintDays int `json:"submissionRangeConstraintDays"`


    // Metadata - Version info metadata for these adherence adjustments settings
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Buadherenceadjustmentssettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buadherenceadjustmentssettings) MarshalJSON() ([]byte, error) {
    type Alias Buadherenceadjustmentssettings

    if BuadherenceadjustmentssettingsMarshalled {
        return []byte("{}"), nil
    }
    BuadherenceadjustmentssettingsMarshalled = true

    return json.Marshal(&struct {
        
        SubmissionRangeConstraintDays int `json:"submissionRangeConstraintDays"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

