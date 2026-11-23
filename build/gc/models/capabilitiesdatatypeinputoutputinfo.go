package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesdatatypeinputoutputinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesdatatypeinputoutputinfoDud struct { 
    


    


    


    


    


    


    


    


    

}

// Capabilitiesdatatypeinputoutputinfo
type Capabilitiesdatatypeinputoutputinfo struct { 
    // InputAny
    InputAny map[string]Capabilitiesavailability `json:"inputAny"`


    // InputColl
    InputColl map[string]Capabilitiesavailability `json:"inputColl"`


    // InputOne
    InputOne map[string]Capabilitiesavailability `json:"inputOne"`


    // InputOutputAny
    InputOutputAny map[string]Capabilitiesavailability `json:"inputOutputAny"`


    // InputOutputColl
    InputOutputColl map[string]Capabilitiesavailability `json:"inputOutputColl"`


    // InputOutputOne
    InputOutputOne map[string]Capabilitiesavailability `json:"inputOutputOne"`


    // OutputAny
    OutputAny map[string]Capabilitiesavailability `json:"outputAny"`


    // OutputColl
    OutputColl map[string]Capabilitiesavailability `json:"outputColl"`


    // OutputOne
    OutputOne map[string]Capabilitiesavailability `json:"outputOne"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatypeinputoutputinfo) String() string {
     o.InputAny = map[string]Capabilitiesavailability{"": {}} 
     o.InputColl = map[string]Capabilitiesavailability{"": {}} 
     o.InputOne = map[string]Capabilitiesavailability{"": {}} 
     o.InputOutputAny = map[string]Capabilitiesavailability{"": {}} 
     o.InputOutputColl = map[string]Capabilitiesavailability{"": {}} 
     o.InputOutputOne = map[string]Capabilitiesavailability{"": {}} 
     o.OutputAny = map[string]Capabilitiesavailability{"": {}} 
     o.OutputColl = map[string]Capabilitiesavailability{"": {}} 
     o.OutputOne = map[string]Capabilitiesavailability{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesdatatypeinputoutputinfo) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesdatatypeinputoutputinfo

    if CapabilitiesdatatypeinputoutputinfoMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesdatatypeinputoutputinfoMarshalled = true

    return json.Marshal(&struct {
        
        InputAny map[string]Capabilitiesavailability `json:"inputAny"`
        
        InputColl map[string]Capabilitiesavailability `json:"inputColl"`
        
        InputOne map[string]Capabilitiesavailability `json:"inputOne"`
        
        InputOutputAny map[string]Capabilitiesavailability `json:"inputOutputAny"`
        
        InputOutputColl map[string]Capabilitiesavailability `json:"inputOutputColl"`
        
        InputOutputOne map[string]Capabilitiesavailability `json:"inputOutputOne"`
        
        OutputAny map[string]Capabilitiesavailability `json:"outputAny"`
        
        OutputColl map[string]Capabilitiesavailability `json:"outputColl"`
        
        OutputOne map[string]Capabilitiesavailability `json:"outputOne"`
        *Alias
    }{

        
        InputAny: map[string]Capabilitiesavailability{"": {}},
        


        
        InputColl: map[string]Capabilitiesavailability{"": {}},
        


        
        InputOne: map[string]Capabilitiesavailability{"": {}},
        


        
        InputOutputAny: map[string]Capabilitiesavailability{"": {}},
        


        
        InputOutputColl: map[string]Capabilitiesavailability{"": {}},
        


        
        InputOutputOne: map[string]Capabilitiesavailability{"": {}},
        


        
        OutputAny: map[string]Capabilitiesavailability{"": {}},
        


        
        OutputColl: map[string]Capabilitiesavailability{"": {}},
        


        
        OutputOne: map[string]Capabilitiesavailability{"": {}},
        

        Alias: (*Alias)(u),
    })
}

