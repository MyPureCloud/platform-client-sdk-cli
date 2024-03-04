package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowcharacteristicsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowcharacteristicsDud struct { 
    


    


    


    


    


    


    


    

}

// Flowcharacteristics - This is a set of enabled characteristics for the loglevel
type Flowcharacteristics struct { 
    // ExecutionItems - Whether to report execution data about individual actions, menus, states, tasks, etc. etc. that ran during execution of the flow.
    ExecutionItems bool `json:"executionItems"`


    // ExecutionInputOutputs - Whether to report input setting input setting values and output data values for individual execution items above.  For example, if you have FlowExecutionInputOutputs and a Call Data Action ran in a flow, if FlowExecutionItems was enabled you'd see the fact a Call Data Action ran and the output path it took but nothing about which Data Action it ran, the input data sent to it at flow runtime and the data returned from it.  If you enable this characteristic, execution data will contain this additional detail.
    ExecutionInputOutputs bool `json:"executionInputOutputs"`


    // Communications - Communications are either audio or digital communications sent to or received from a participant.  An example here would be the initial greeting in an inbound call flow where it plays a greeting message to the participant.
    Communications bool `json:"communications"`


    // EventError - Whether to report flow error events.
    EventError bool `json:"eventError"`


    // EventWarning - Whether to report flow warning events.
    EventWarning bool `json:"eventWarning"`


    // EventOther - Whether to report events other than errors or warnings such as a language change, loop event.
    EventOther bool `json:"eventOther"`


    // Variables - Whether to report assignment of values to variables in flow execution data. It's important to remember there is a difference between variable value assignments and output data from an action.  If you have a Call Digital Bot flow action in an Inbound Message flow and there is no variable bound to the Exit Reason output but FlowExecutionInputOutputs is enabled, you will still be able to see the exit reason from the digital bot flow in execution data even though it is not bound to a variable.
    Variables bool `json:"variables"`


    // Names - This characteristic specifies whether or not name information should be emitted in execution data such as action, task, state or even the flow name itself.  Names are very handy from a readability standpoint but they do take up additional space in flow execution data instances.
    Names bool `json:"names"`

}

// String returns a JSON representation of the model
func (o *Flowcharacteristics) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowcharacteristics) MarshalJSON() ([]byte, error) {
    type Alias Flowcharacteristics

    if FlowcharacteristicsMarshalled {
        return []byte("{}"), nil
    }
    FlowcharacteristicsMarshalled = true

    return json.Marshal(&struct {
        
        ExecutionItems bool `json:"executionItems"`
        
        ExecutionInputOutputs bool `json:"executionInputOutputs"`
        
        Communications bool `json:"communications"`
        
        EventError bool `json:"eventError"`
        
        EventWarning bool `json:"eventWarning"`
        
        EventOther bool `json:"eventOther"`
        
        Variables bool `json:"variables"`
        
        Names bool `json:"names"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

