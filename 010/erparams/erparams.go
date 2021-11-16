package erparams

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type SendConfig struct {
	Key   string `json:"key"`
	Api   string `json:"api"`
	Param string `json:"param"`
}

type ErParams struct {
	StartDate  string `json:"start_date"`
	Hour       string `json:"hour"`
	FinalId    string `json:"final_id"`
	DateForm   string `json:"date_form"`
	DateBefore string `json:"date_before"`
	DateAfter  string `json:"date_after"`
	HourForm   string `json:"hour_form"`
	Dim        string `json:"dim"`
	Map        string `json:"map"`
	Const      string `json:"const"`
	Cal        string `json:"cal"`
	Group      string `json:"group"`
	Aggregate  string `json:"_aggregate"`
}

type ErMapping struct {
	MappingField string            `json:"mapping_field"`
	MappingName  string            `json:"mapping_name"`
	DefaultValue string            `json:"default_value"`
	Mapping      map[string]string `json:"mapping"`
}

type ErConst struct {
	CustomField string             `json:"custom_field"`
	Expression  *ErConstExpression `json:"expression"`
}

type ErCal struct {
	CustomField string        `json:"custom_field"`
	Expression  *ErExpression `json:"expression"`
}

type ErConstExpression struct {
	Op     string `json:"op"`
	Params string `json:"params"`
}
type ErExpression struct {
	Op     string   `json:"op"`
	Params []string `json:"params"`
}

type ErAggregate struct {
	FieldName  string `json:"fieldName"`
	Expression string `json:"expression"`
	Method     string `json:"method"`
}


const (
	All         = "all"
	START_DATE  = "start_date"
	HOUR        = "hour"
	FINAL_ID    = "final_id"
	DATE_FORM   = "date_form"
	DATE_BEFORE = "date_before"
	DATE_AFTER  = "date_after"
	HOUR_FORM   = "hour_form"
	DIM         = "dim"
	GROUP       = "group"
	MAP         = "map"
	CONST       = "const"
	CAL         = "cal"
	AGGREGATE   = "_aggregate"
)

var fieldMap = map[string]func([]*SendConfig, string) error{
	All:         printAll,
	START_DATE:  printBase,
	HOUR:        printBase,
	FINAL_ID:    printBase,
	DATE_FORM:   printBase,
	DATE_BEFORE: printBase,
	DATE_AFTER:  printBase,
	HOUR_FORM:   printBase,
	DIM:         printBase,
	GROUP:       printBase,
	MAP:         printNested,
	CONST:       printNested,
	CAL:         printNested,
	AGGREGATE:   printNested,
}

func DoParse(value string, field string) error {
	sendConfigs,err := parseSendConfig(value)
	if err != nil {
	 	return err
	}

	err = fieldMap[field](sendConfigs, field)
	return err
}

func parseSendConfig(str string) ([]*SendConfig, error) {
	sendConfigs := []*SendConfig{}
	err := json.Unmarshal([]byte(str), &sendConfigs)

	return sendConfigs, errors.Wrapf(err, "解析SendConfig出错")
}

func printAll(sendConfigs []*SendConfig, _ string) error {
	for _, config := range sendConfigs {
		var out bytes.Buffer
		err := json.Indent(&out, []byte(config.Param), "", "\t")
		if err != nil {
			return errors.Wrapf(err, "解析erParams出错！")
		}

		//fmt.Printf("request: %s\n", config.Key)
		fmt.Printf("%v\n", out.String())
	}
	return nil
}

func printBase(sendConfigs []*SendConfig, field string) error {
	for _, config := range sendConfigs {
		var erParams ErParams

		err := json.Unmarshal([]byte(config.Param), &erParams)
		if err != nil {
			return errors.Wrapf(err, "解析erParams出错！")
		}

		var v string
		switch field {
		case START_DATE:
			v = erParams.StartDate
		case HOUR:
			v = erParams.Hour
		case FINAL_ID:
			v = erParams.FinalId
		case DATE_FORM:
			v = erParams.DateForm
		case DATE_BEFORE:
			v = erParams.DateBefore
		case DATE_AFTER:
			v = erParams.DateAfter
		case HOUR_FORM:
			v = erParams.HourForm
		case DIM:
			v = erParams.Dim
		case GROUP:
			v = erParams.Group
		default:
			return errors.New("none field")
		}

		//fmt.Printf("request: %s\n", config.Key)
		fmt.Printf("%v\n", v)
	}
	return nil
}

func printNested(sendConfigs []*SendConfig, field string) error {
	for _, config := range sendConfigs {
		var erParams ErParams

		err := json.Unmarshal([]byte(config.Param), &erParams)
		if err != nil {
			return errors.Wrapf(err, "解析erParams出错！")
		}

		var v []byte
		switch field {
		case MAP:
			v = []byte(erParams.Map)
		case CONST:
			v = []byte(erParams.Const)
		case CAL:
			v = []byte(erParams.Cal)
		case AGGREGATE:
			v = []byte(erParams.Aggregate)
		default:
			return errors.New("none field")
		}

		var out bytes.Buffer
		err = json.Indent(&out, v, "", "\t")
		if err != nil {
			return errors.Wrapf(err, "解析erParams出错！")
		}

		//fmt.Printf("request: %s\n", config.Key)
		fmt.Printf("%v\n", out.String())
	}
	return nil
}