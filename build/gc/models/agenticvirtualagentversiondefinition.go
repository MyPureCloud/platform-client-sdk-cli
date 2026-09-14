package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentversiondefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentversiondefinitionDud struct { 
    


    


    


    


    


    


    


    

}

// Agenticvirtualagentversiondefinition - Definition of a virtual agent version.
type Agenticvirtualagentversiondefinition struct { 
    // Role - A brief description of the virtual agent's high-level capabilities. Specific policies go in the instructions field.
    Role string `json:"role"`


    // Instructions - List of instructions, rules, or guidelines the virtual agent must always follow.
    Instructions []string `json:"instructions"`


    // Guardrails - Custom guardrail rules used to detect and block matching user behavior.
    Guardrails Agenticvirtualagentguardrails `json:"guardrails"`


    // Tools - Tools available to the virtual agent. Each tool is a single action the virtual agent can take, and the virtual agent can call multiple tools in a single turn before responding to the user.
    Tools []Agenticvirtualagenttool `json:"tools"`


    // Types - Types the virtual agent can use for tool inputs, return types, and errors.
    Types []Agenticvirtualagenttypedefinition `json:"types"`


    // Events - Event settings for the virtual agent. Guardrails events configure generated guardrail behavior.
    Events []Agenticvirtualagenteventsettings `json:"events"`


    // Settings - Additional settings for the virtual agent version.
    Settings Agenticvirtualagentversionsettings `json:"settings"`


    // Model - The model powering the virtual agent version. Where a new model version is available, the Preview value can be used to opt into using it for your virtual agent. Note: Preview model versions will over time become the Stable version used.
    Model string `json:"model"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentversiondefinition) String() string {
    
     o.Instructions = []string{""} 
    
     o.Tools = []Agenticvirtualagenttool{{}} 
     o.Types = []Agenticvirtualagenttypedefinition{{}} 
     o.Events = []Agenticvirtualagenteventsettings{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentversiondefinition) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentversiondefinition

    if AgenticvirtualagentversiondefinitionMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentversiondefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Role string `json:"role"`
        
        Instructions []string `json:"instructions"`
        
        Guardrails Agenticvirtualagentguardrails `json:"guardrails"`
        
        Tools []Agenticvirtualagenttool `json:"tools"`
        
        Types []Agenticvirtualagenttypedefinition `json:"types"`
        
        Events []Agenticvirtualagenteventsettings `json:"events"`
        
        Settings Agenticvirtualagentversionsettings `json:"settings"`
        
        Model string `json:"model"`
        *Alias
    }{

        


        
        Instructions: []string{""},
        


        


        
        Tools: []Agenticvirtualagenttool{{}},
        


        
        Types: []Agenticvirtualagenttypedefinition{{}},
        


        
        Events: []Agenticvirtualagenteventsettings{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

