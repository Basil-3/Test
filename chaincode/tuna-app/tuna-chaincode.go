// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
*/

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

//SmartContract definition
type SmartContract struct {
}

type Unemployed struct {
	UEid        string `json:"employee_id"`
	UEfirstname string `json:"employee_first_name"`
	UElastname  string `json:"employee_last_name"`
	Empstatus   string `json:"employment_status"`
	UEage       string `json:"employee_age"`
	UEregion    string `json:"employee_region"`
}

type Labor struct {
	Lid        string `json:"labor_id"`
	Lfirstname string `json:"labor_first_name"`
	Llastname  string `json:"labor_last_name"`
	Lage       string `json:"labor_age"`
	Lregion    string `json:"labor_region"`
}

type MSME struct {
	Cid     string `json:"company_registration_number"`
	Cname   string `json:"company_name"`
	Csize   string `json:"company_size"`
	Cregion string `json:"company_base_region"`
}

type NAV struct {
	NAVid        string `json:"NAV_officer_id"`
	NAVfirstname string `json:"NAV_officer_first_name"`
	NAVlastname  string `json:"NAV_officer_last_name"`
}

type GOVT struct {
	GOVid        string `json:"government_officer_id"`
	GOVfirstname string `json:"government_officer_first_name"`
	GOVlastname  string `json:"government_officer_last_name"`
}

type Bank struct {
	Bid       string `json:"bank_branch_code"`
	Bname     string `json:"bank_name"`
	Bcapacity string `json:"bank_monetary_capacity"`
	Bregion   string `json:""bank_region`
}

//Init Instantiation of chaincode
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke of chaincode
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "registerNAV" {
		return s.registerNAV(APIstub, args)
	} else if function == "registerEmployee" {
		return s.registerEmployee(APIstub, args)
	} else if function == "registerGovt" {
		return s.registerGovt(APIstub, args)
	} else if function == "registerBank" {
		return s.registerBank(APIstub, args)
	} else if function == "registerCompany" {
		return s.registerCompany(APIstub, args)
	} else if function == "registerLabor" {
		return s.registerLabor(APIstub, args)
	} else if function == "queryEmployee" {
		return s.queryEmployee(APIstub)
	} else if function == "queryLabor" {
		return s.queryLabor(APIstub)
	} else if function == "changeEmploymentStatus" {
		return s.changeEmploymentStatus(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/* The recordTuna method *
Fisherman like Sarah would use to record each of her tuna catches.
This method takes in five arguments (attributes to be saved in the ledger).
*/
func (s *SmartContract) registerNAV(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	var nav = NAV{NAVid: args[1], NAVfirstname: args[2], NAVlastname: args[3]}

	navAsBytes, _ := json.Marshal(nav)
	err := APIstub.PutState(args[0], navAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to register NAV officer: %s", args[1]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) registerEmployee(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	var ue = Unemployed{UEid: args[1], UEfirstname: args[2], UElastname: args[3], Empstatus: args[4], UEage: args[5], UEregion: args[6]}

	ueAsBytes, _ := json.Marshal(ue)
	err := APIstub.PutState(args[0], ueAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to register Employee: %s", args[1]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) registerGovt(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	var gov = GOVT{GOVid: args[1], GOVfirstname: args[2], GOVlastname: args[3]}

	govAsBytes, _ := json.Marshal(gov)
	err := APIstub.PutState(args[0], govAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to register government officer: %s", args[1]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) registerBank(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var bank = Bank{Bid: args[1], Bname: args[2], Bcapacity: args[3], Bregion: args[4]}

	bankAsBytes, _ := json.Marshal(bank)
	err := APIstub.PutState(args[0], bankAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to register Bank: %s", args[1]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) registerCompany(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var msme = MSME{Cid: args[1], Cname: args[2], Csize: args[3], Cregion: args[4]}

	msmeAsBytes, _ := json.Marshal(msme)
	err := APIstub.PutState(args[0], msmeAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to register Company: %s", args[1]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) registerLabor(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	var labor = Labor{Lid: args[1], Lfirstname: args[2], Llastname: args[3], Lage: args[5], Lregion: args[6]}

	laborAsBytes, _ := json.Marshal(labor)
	err := APIstub.PutState(args[0], laborAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to register Labor: %s", args[1]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) queryEmployee(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ueAsBytes, _ := APIstub.GetState(args[0])
	if ueAsBytes == nil {
		return shim.Error("Could not locate employee")
	}
	return shim.Success(ueAsBytes)
}

func (s *SmartContract) queryLabor(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	laborAsBytes, _ := APIstub.GetState(args[0])
	if laborAsBytes == nil {
		return shim.Error("Could not locate labor")
	}
	return shim.Success(laborAsBytes)
}

func (s *SmartContract) changeEmploymentStatus(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	ueAsBytes, _ := APIstub.GetState(args[0])
	if ueAsBytes == nil {
		return shim.Error("Could not locate employee")
	}
	ue := Unemployed{}

	json.Unmarshal(ueAsBytes, &ue)
	// Normally check that the specified argument is a valid holder of tuna
	// we are skipping this check for this example
	ue.Empstatus = args[1]

	ueAsBytes, _ = json.Marshal(ue)
	err := APIstub.PutState(args[0], ueAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change employment status: %s", args[0]))
	}

	return shim.Success(nil)
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
