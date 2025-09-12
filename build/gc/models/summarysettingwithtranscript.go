package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummarysettingwithtranscriptMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummarysettingwithtranscriptDud struct { 
    


    


    

}

// Summarysettingwithtranscript
type Summarysettingwithtranscript struct { 
    // Transcript - Example transcript to preview with the setting.
    Transcript string `json:"transcript"`


    // SummarySetting - Summary setting to preview on the transcript.
    SummarySetting Summarysetting `json:"summarySetting"`


    // SummaryPreviewSessionId - Session identifier of the summary preview.
    SummaryPreviewSessionId string `json:"summaryPreviewSessionId"`

}

// String returns a JSON representation of the model
func (o *Summarysettingwithtranscript) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summarysettingwithtranscript) MarshalJSON() ([]byte, error) {
    type Alias Summarysettingwithtranscript

    if SummarysettingwithtranscriptMarshalled {
        return []byte("{}"), nil
    }
    SummarysettingwithtranscriptMarshalled = true

    return json.Marshal(&struct {
        
        Transcript string `json:"transcript"`
        
        SummarySetting Summarysetting `json:"summarySetting"`
        
        SummaryPreviewSessionId string `json:"summaryPreviewSessionId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

