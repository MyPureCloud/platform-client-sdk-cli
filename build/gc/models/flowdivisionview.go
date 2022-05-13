package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowdivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowdivisionviewDud struct { 
    


    


    


    


    


    


    


    SupportedLanguages []Supportedlanguage `json:"supportedLanguages"`


    


    


    SelfUri string `json:"selfUri"`

}

// Flowdivisionview
type Flowdivisionview struct { 
    // Id - The flow identifier
    Id string `json:"id"`


    // Name - The flow name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // VarType
    VarType string `json:"type"`


    // Description - the flow description
    Description string `json:"description"`


    // InputSchema - json schema describing the inputs for the flow
    InputSchema Jsonschemadocument `json:"inputSchema"`


    // OutputSchema - json schema describing the outputs for the flow
    OutputSchema Jsonschemadocument `json:"outputSchema"`


    


    // PublishedVersion - published version information if there is a published version
    PublishedVersion Flowversion `json:"publishedVersion"`


    // DebugVersion - debug version information if there is a debug version
    DebugVersion Flowversion `json:"debugVersion"`


    

}

// String returns a JSON representation of the model
func (o *Flowdivisionview) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowdivisionview) MarshalJSON() ([]byte, error) {
    type Alias Flowdivisionview

    if FlowdivisionviewMarshalled {
        return []byte("{}"), nil
    }
    FlowdivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        VarType string `json:"type"`
        
        Description string `json:"description"`
        
        InputSchema Jsonschemadocument `json:"inputSchema"`
        
        OutputSchema Jsonschemadocument `json:"outputSchema"`
        
        PublishedVersion Flowversion `json:"publishedVersion"`
        
        DebugVersion Flowversion `json:"debugVersion"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

