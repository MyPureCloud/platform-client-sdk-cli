package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OperandMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OperandDud struct { 
    


    


    


    


    


    


    


    


    

}

// Operand
type Operand struct { 
    // VarType - The Operand type of the category
    VarType string `json:"type"`


    // Occurrence - The minimum number of occurrences of the defined operand type
    Occurrence int `json:"occurrence"`


    // Inverted - Applies a NOT modifier to the operand group
    Inverted bool `json:"inverted"`


    // Term - Filter interaction by word(s)
    Term Term `json:"term"`


    // TopicId - Filter interaction by topic ID
    TopicId string `json:"topicId"`


    // VoiceSecondsPosition - Dictates when the operand must occur in a voice interaction
    VoiceSecondsPosition Operandposition `json:"voiceSecondsPosition"`


    // DigitalWordsPosition - Dictates when the operand must occur in a digital interaction
    DigitalWordsPosition Operandposition `json:"digitalWordsPosition"`


    // InfixOperator - Defines a logical operation that is applied on the current operand, against the following operand
    InfixOperator Infixoperator `json:"infixOperator"`


    // Operands - Contains a new level of operands
    Operands []Operand `json:"operands"`

}

// String returns a JSON representation of the model
func (o *Operand) String() string {
    
    
    
    
    
    
    
    
     o.Operands = []Operand{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Operand) MarshalJSON() ([]byte, error) {
    type Alias Operand

    if OperandMarshalled {
        return []byte("{}"), nil
    }
    OperandMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Occurrence int `json:"occurrence"`
        
        Inverted bool `json:"inverted"`
        
        Term Term `json:"term"`
        
        TopicId string `json:"topicId"`
        
        VoiceSecondsPosition Operandposition `json:"voiceSecondsPosition"`
        
        DigitalWordsPosition Operandposition `json:"digitalWordsPosition"`
        
        InfixOperator Infixoperator `json:"infixOperator"`
        
        Operands []Operand `json:"operands"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        Operands: []Operand{{}},
        

        Alias: (*Alias)(u),
    })
}

