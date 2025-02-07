package model

import "fmt"

type AlarmConfig struct {
	Channel   string               `json:"channel"`
	Door      string               `json:"door"`
	Priority  AlarmConfigPriority  `json:"priority"`
	Condition AlarmConfigCondition `json:"condition"`
	Receiver  AlarmConfigReceiver  `json:"receiver"`
}

type AlarmConfigCondition struct {
	Operator Operator `json:"operator"`
	Value    int      `json:"value"`
	Duration int      `json:"duration"`
	Unit     Unit     `json:"unit"`
}

func (a AlarmConfigCondition) String() string {

	operatorSymbol := "="

	if a.Operator == OperatorGreaterThan {
		operatorSymbol = ">"
	} else if a.Operator == OperatorLessThan {
		operatorSymbol = "<"
	} else if a.Operator == OperatorGreaterThanOrEqualTo {
		operatorSymbol = ">="
	} else if a.Operator == OperatorLessThanOrEqualTo {
		operatorSymbol = "<="
	}

	return fmt.Sprintf("%s %d Selama %d %s", operatorSymbol, a.Value, a.Duration, a.Unit)
}

type AlarmConfigReceiver string

const (
	AlarmConfigReceiverAdam      = "Adam"
	AlarmConfigReceiverHawa      = "Hawa"
	AlarmConfigReceiverWhatsapp  = "Whatsapp"
	AlarmConfigReceiverEmail     = "Email"
	AlarmConfigReceiverDashboard = "Dashboard"
)

type AlarmConfigPriority string

const (
	AlarmConfigPriorityWarning  = "Warning"
	AlarmConfigPriorityCritical = "Critical"
)

type Unit string

const (
	UnitSecond Unit = "Second"
	UnitMinute Unit = "Minute"
	UnitHour   Unit = "Hour"
)

type Operator string

const (
	OperatorLessThan             Operator = "LessThan"
	OperatorGreaterThan          Operator = "GreaterThan"
	OperatorLessThanOrEqualTo    Operator = "LessThanOrEqualTo"
	OperatorGreaterThanOrEqualTo Operator = "GreaterOrEqualTo"
	OperatorEqualTo              Operator = "EqualTo"
)
