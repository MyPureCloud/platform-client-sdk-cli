package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NludomainversionqualityreportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NludomainversionqualityreportDud struct { 
    


    


    

}

// Nludomainversionqualityreport
type Nludomainversionqualityreport struct { 
    // Version - The domain and version details of the quality report
    Version Nludomainversion `json:"version"`


    // ConfusionMatrix - The confusion matrix for the Domain Version
    ConfusionMatrix []Nluconfusionmatrixrow `json:"confusionMatrix"`


    // Summary - The quality report summary for the Domain Version
    Summary Nluqualityreportsummary `json:"summary"`

}

// String returns a JSON representation of the model
func (o *Nludomainversionqualityreport) String() string {
    
    
    
    
    
    
     o.ConfusionMatrix = []Nluconfusionmatrixrow{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nludomainversionqualityreport) MarshalJSON() ([]byte, error) {
    type Alias Nludomainversionqualityreport

    if NludomainversionqualityreportMarshalled {
        return []byte("{}"), nil
    }
    NludomainversionqualityreportMarshalled = true

    return json.Marshal(&struct { 
        Version Nludomainversion `json:"version"`
        
        ConfusionMatrix []Nluconfusionmatrixrow `json:"confusionMatrix"`
        
        Summary Nluqualityreportsummary `json:"summary"`
        
        *Alias
    }{
        

        

        

        
        ConfusionMatrix: []Nluconfusionmatrixrow{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

