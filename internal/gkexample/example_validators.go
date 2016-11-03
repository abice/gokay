package gkexample

import (
	"errors"

	"github.com/zencoder/gokay/gokay"
)

func (s *HasValidateImplicit) Validate() error {
	em := make(gokay.ErrorMap)

	// BEGIN InvalidStruct Validations
	var errorsInvalidStruct gokay.ErrorSlice
	// dive

	if s.InvalidStruct != nil {
		if err := gokay.Validate(s.InvalidStruct); err != nil {
			errorsInvalidStruct = append(errorsInvalidStruct, err)
		}
	}

	if len(errorsInvalidStruct) > 0 {
		em["InvalidStruct"] = errorsInvalidStruct
	}
	// END InvalidStruct Validations

	// BEGIN ValidStruct Validations
	var errorsValidStruct gokay.ErrorSlice
	// dive

	if err := gokay.Validate(s.ValidStruct); err != nil {
		errorsValidStruct = append(errorsValidStruct, err)
	}

	if len(errorsValidStruct) > 0 {
		em["ValidStruct"] = errorsValidStruct
	}
	// END ValidStruct Validations

	if len(em) > 0 {
		return em
	}
	return nil
}

func (s *NotNilTestStruct) Validate() error {
	em := make(gokay.ErrorMap)

	// BEGIN NotNilMap Validations
	var errorsNotNilMap gokay.ErrorSlice
	// NotNil

	if s.NotNilMap == nil {
		errorsNotNilMap = append(errorsNotNilMap, errors.New("is Nil"))
	}

	if len(errorsNotNilMap) > 0 {
		em["NotNilMap"] = errorsNotNilMap
	}
	// END NotNilMap Validations

	// BEGIN NotNilSlice Validations
	var errorsNotNilSlice gokay.ErrorSlice
	// NotNil

	if s.NotNilSlice == nil {
		errorsNotNilSlice = append(errorsNotNilSlice, errors.New("is Nil"))
	}

	if len(errorsNotNilSlice) > 0 {
		em["NotNilSlice"] = errorsNotNilSlice
	}
	// END NotNilSlice Validations

	if len(em) > 0 {
		return em
	}
	return nil
}

func (s *ExampleStruct) Validate() error {
	em := make(gokay.ErrorMap)

	// BEGIN BCP47String Validations
	var errorsBCP47String gokay.ErrorSlice
	// BCP47

	if err := gokay.IsBCP47(&s.BCP47String); err != nil {
		errorsBCP47String = append(errorsBCP47String, err)
	}

	if len(errorsBCP47String) > 0 {
		em["BCP47String"] = errorsBCP47String
	}
	// END BCP47String Validations

	// BEGIN BCP47StringPtr Validations
	var errorsBCP47StringPtr gokay.ErrorSlice
	// NotNil

	if s.BCP47StringPtr == nil {
		errorsBCP47StringPtr = append(errorsBCP47StringPtr, errors.New("is Nil"))
	}

	// BCP47

	if err := gokay.IsBCP47(s.BCP47StringPtr); err != nil {
		errorsBCP47StringPtr = append(errorsBCP47StringPtr, err)
	}

	if len(errorsBCP47StringPtr) > 0 {
		em["BCP47StringPtr"] = errorsBCP47StringPtr
	}
	// END BCP47StringPtr Validations

	// BEGIN CanBeNilWithConstraints Validations
	var errorsCanBeNilWithConstraints gokay.ErrorSlice
	// Length

	if err := gokay.LengthString(12, s.CanBeNilWithConstraints); err != nil {
		errorsCanBeNilWithConstraints = append(errorsCanBeNilWithConstraints, err)
	}

	if len(errorsCanBeNilWithConstraints) > 0 {
		em["CanBeNilWithConstraints"] = errorsCanBeNilWithConstraints
	}
	// END CanBeNilWithConstraints Validations

	// BEGIN HexString Validations
	var errorsHexString gokay.ErrorSlice
	// Length

	if err := gokay.LengthString(12, &s.HexString); err != nil {
		errorsHexString = append(errorsHexString, err)
	}

	// Hex

	if err := gokay.IsHex(&s.HexString); err != nil {
		errorsHexString = append(errorsHexString, err)
	}

	if len(errorsHexString) > 0 {
		em["HexString"] = errorsHexString
	}
	// END HexString Validations

	// BEGIN HexStringPtr Validations
	var errorsHexStringPtr gokay.ErrorSlice
	// Length

	if err := gokay.LengthString(16, s.HexStringPtr); err != nil {
		errorsHexStringPtr = append(errorsHexStringPtr, err)
	}

	// NotNil

	if s.HexStringPtr == nil {
		errorsHexStringPtr = append(errorsHexStringPtr, errors.New("is Nil"))
	}

	// Hex

	if err := gokay.IsHex(s.HexStringPtr); err != nil {
		errorsHexStringPtr = append(errorsHexStringPtr, err)
	}

	if len(errorsHexStringPtr) > 0 {
		em["HexStringPtr"] = errorsHexStringPtr
	}
	// END HexStringPtr Validations

	if len(em) > 0 {
		return em
	}
	return nil
}

func (s *Example) Validate() error {
	em := make(gokay.ErrorMap)

	// BEGIN MapOfInterfaces Validations
	var errorsMapOfInterfaces gokay.ErrorSlice
	// NotNil

	if s.MapOfInterfaces == nil {
		errorsMapOfInterfaces = append(errorsMapOfInterfaces, errors.New("is Nil"))
	}

	if len(errorsMapOfInterfaces) > 0 {
		em["MapOfInterfaces"] = errorsMapOfInterfaces
	}
	// END MapOfInterfaces Validations

	if len(em) > 0 {
		return em
	}
	return nil
}
