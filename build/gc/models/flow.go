package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    NluInfo Nluinfo `json:"nluInfo"`


    SupportedLanguages []Supportedlanguage `json:"supportedLanguages"`


    CompatibleFlowTypes []string `json:"compatibleFlowTypes"`


    


    


    SelfUri string `json:"selfUri"`

}

// Flow
type Flow struct { 
    // Id - The flow identifier
    Id string `json:"id"`


    // Name - The flow name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description
    Description string `json:"description"`


    // VarType
    VarType string `json:"type"`


    // LockedUser - User that has the flow locked.
    LockedUser User `json:"lockedUser"`


    // LockedClient - OAuth client that has the flow locked.
    LockedClient Domainentityref `json:"lockedClient"`


    // Active
    Active bool `json:"active"`


    // System
    System bool `json:"system"`


    // Deleted
    Deleted bool `json:"deleted"`


    // PublishedVersion
    PublishedVersion Flowversion `json:"publishedVersion"`


    // SavedVersion
    SavedVersion Flowversion `json:"savedVersion"`


    // InputSchema - json schema describing the inputs for the flow
    InputSchema interface{} `json:"inputSchema"`


    // OutputSchema - json schema describing the outputs for the flow
    OutputSchema interface{} `json:"outputSchema"`


    // CheckedInVersion
    CheckedInVersion Flowversion `json:"checkedInVersion"`


    // DebugVersion
    DebugVersion Flowversion `json:"debugVersion"`


    // PublishedBy
    PublishedBy User `json:"publishedBy"`


    // CurrentOperation
    CurrentOperation Operation `json:"currentOperation"`


    


    


    


    // WorktypeId
    WorktypeId string `json:"worktypeId"`


    // VirtualAgentEnabled - Indicates whether virtual agent is enabled for this flow.
    VirtualAgentEnabled bool `json:"virtualAgentEnabled"`


    

}

// String returns a JSON representation of the model
func (o *Flow) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.InputSchema = Interface{} 
     o.OutputSchema = Interface{} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flow) MarshalJSON() ([]byte, error) {
    type Alias Flow

    if FlowMarshalled {
        return []byte("{}"), nil
    }
    FlowMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        
        VarType string `json:"type"`
        
        LockedUser User `json:"lockedUser"`
        
        LockedClient Domainentityref `json:"lockedClient"`
        
        Active bool `json:"active"`
        
        System bool `json:"system"`
        
        Deleted bool `json:"deleted"`
        
        PublishedVersion Flowversion `json:"publishedVersion"`
        
        SavedVersion Flowversion `json:"savedVersion"`
        
        InputSchema interface{} `json:"inputSchema"`
        
        OutputSchema interface{} `json:"outputSchema"`
        
        CheckedInVersion Flowversion `json:"checkedInVersion"`
        
        DebugVersion Flowversion `json:"debugVersion"`
        
        PublishedBy User `json:"publishedBy"`
        
        CurrentOperation Operation `json:"currentOperation"`
        
        WorktypeId string `json:"worktypeId"`
        
        VirtualAgentEnabled bool `json:"virtualAgentEnabled"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        InputSchema: Interface{},
        


        
        OutputSchema: Interface{},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

