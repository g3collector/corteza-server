package event

import (
	"encoding/json"
)

// Match returns false if given conditions do not match event & resource internals
func (res pageBase) Match(name string, op string, values ...string) bool {
	// By default we match no mather what kind of constraints we receive
	//
	// Function will be called multiple times - once for every trigger constraint
	// All should match (return true):
	//   constraint#1 AND constraint#2 AND constraint#3 ...
	//
	// When there are multiple values, Match() can decide how to treat them (OR, AND...)
	return true
}

// Encode internal data to be passed as event params & arguments to triggered Corredor script
func (res pageBase) Encode() (args map[string][]byte, err error) {
	args = make(map[string][]byte)

	if args["page"], err = json.Marshal(res.page); err != nil {
		return nil, err
	}

	if args["oldPage"], err = json.Marshal(res.oldPage); err != nil {
		return nil, err
	}

	if args["namespace"], err = json.Marshal(res.namespace); err != nil {
		return nil, err
	}

	if args["invoker"], err = json.Marshal(res.invoker); err != nil {
		return nil, err
	}

	return
}

// Decode return values from Corredor script into struct props
func (res *pageBase) Decode(results map[string][]byte) (err error) {
	if r, ok := results["result"]; ok && len(results) == 1 {
		if err = json.Unmarshal(r, res.page); err != nil {
			return
		}
	}

	if r, ok := results["page"]; ok && len(results) == 1 {
		if err = json.Unmarshal(r, res.page); err != nil {
			return
		}
	}

	if r, ok := results["invoker"]; ok && len(results) == 1 {
		if err = json.Unmarshal(r, res.invoker); err != nil {
			return
		}
	}
	return
}
