// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package job_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PredicateOp Op is the operation to apply.
//
//  - EQUALS: Operators on scalar values. Only applies to one of |int_value|,
// |long_value|, |string_value| or |timestamp_value|.
//  - IN: Checks if the value is a member of a given array, which should be one of
// |int_values|, |long_values| or |string_values|.
// swagger:model PredicateOp
type PredicateOp string

const (

	// PredicateOpUNKNOWN captures enum value "UNKNOWN"
	PredicateOpUNKNOWN PredicateOp = "UNKNOWN"

	// PredicateOpEQUALS captures enum value "EQUALS"
	PredicateOpEQUALS PredicateOp = "EQUALS"

	// PredicateOpNOTEQUALS captures enum value "NOT_EQUALS"
	PredicateOpNOTEQUALS PredicateOp = "NOT_EQUALS"

	// PredicateOpGREATERTHAN captures enum value "GREATER_THAN"
	PredicateOpGREATERTHAN PredicateOp = "GREATER_THAN"

	// PredicateOpGREATERTHANEQUALS captures enum value "GREATER_THAN_EQUALS"
	PredicateOpGREATERTHANEQUALS PredicateOp = "GREATER_THAN_EQUALS"

	// PredicateOpLESSTHAN captures enum value "LESS_THAN"
	PredicateOpLESSTHAN PredicateOp = "LESS_THAN"

	// PredicateOpLESSTHANEQUALS captures enum value "LESS_THAN_EQUALS"
	PredicateOpLESSTHANEQUALS PredicateOp = "LESS_THAN_EQUALS"

	// PredicateOpIN captures enum value "IN"
	PredicateOpIN PredicateOp = "IN"
)

// for schema
var predicateOpEnum []interface{}

func init() {
	var res []PredicateOp
	if err := json.Unmarshal([]byte(`["UNKNOWN","EQUALS","NOT_EQUALS","GREATER_THAN","GREATER_THAN_EQUALS","LESS_THAN","LESS_THAN_EQUALS","IN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		predicateOpEnum = append(predicateOpEnum, v)
	}
}

func (m PredicateOp) validatePredicateOpEnum(path, location string, value PredicateOp) error {
	if err := validate.Enum(path, location, value, predicateOpEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this predicate op
func (m PredicateOp) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePredicateOpEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
