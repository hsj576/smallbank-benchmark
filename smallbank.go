/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	//"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmallBank Chaincode implementation
type SmallBank struct {
	contractapi.Contract
}

func (t *SmallBank) CreateAccount(ctx contractapi.TransactionContextInterface, Account string, Saving int, Check int) error {
	fmt.Println("SmallBank Create Account")
	var err error
	// Initialize the chaincode
	err = ctx.GetStub().PutPrivateData("collectionSaving", Account, []byte(strconv.Itoa(Saving)))
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData("collectionCheck", Account, []byte(strconv.Itoa(Check)))
	if err != nil {
		return err
	}
	fmt.Printf("Saving = %d, Check = %d\n", Saving, Check)
	// Write the state to the ledger
	return nil
}

func (t *SmallBank) AMALGAMATE(ctx contractapi.TransactionContextInterface, A string, B string) error {
	var err error
	var Acheck int
	var Bsaving int
	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Acheckbytes, err := ctx.GetStub().GetPrivateData("collectionCheck", A)
	if err != nil {
		return fmt.Errorf("Failed to get state")
	}
	if Acheckbytes == nil {
		return fmt.Errorf("Entity not found")
	}
	Acheck, _ = strconv.Atoi(string(Acheckbytes))
	
	err = ctx.GetStub().PutPrivateData("collectionCheck", A, []byte(strconv.Itoa(0)))
	if err != nil {
		return err
	}

	Bsavingbytes, err := ctx.GetStub().GetPrivateData("collectionSaving", B)
	if err != nil {
		return fmt.Errorf("Failed to get state")
	}
	if Bsavingbytes == nil {
		return fmt.Errorf("Entity not found")
	}
	Bsaving, _ = strconv.Atoi(string(Bsavingbytes))
	Bsaving = Bsaving + Acheck

	err = ctx.GetStub().PutPrivateData("collectionSaving", B, []byte(strconv.Itoa(Bsaving)))
	if err != nil {
		return err
	}

	return nil
}

func (t *SmallBank) DEPOSIT_CHECKING(ctx contractapi.TransactionContextInterface, A string, X int) error {
	var err error
	var Acheck int
	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Acheckbytes, err := ctx.GetStub().GetPrivateData("collectionCheck", A)
	if err != nil {
		return fmt.Errorf("Failed to get state")
	}
	if Acheckbytes == nil {
		return fmt.Errorf("Entity not found")
	}
	Acheck, _ = strconv.Atoi(string(Acheckbytes))
	Acheck = Acheck + X
	fmt.Printf("Acheck = %d\n", Acheck)
	// Write the state back to the ledger
	err = ctx.GetStub().PutPrivateData("collectionCheck", A, []byte(strconv.Itoa(Acheck)))
	if err != nil {
		return err
	}
	return nil
}

func (t *SmallBank) SEND_PAYMENT(ctx contractapi.TransactionContextInterface, A string, B string, X int) error {
	var err error
	var Acheck int
	var Bcheck int
	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Acheckbytes, err := ctx.GetStub().GetPrivateData("collectionCheck", A)
	if err != nil {
		return fmt.Errorf("Failed to get state")
	}
	if Acheckbytes == nil {
		return fmt.Errorf("Entity not found")
	}
	Acheck, _ = strconv.Atoi(string(Acheckbytes))

	Bcheckbytes, err := ctx.GetStub().GetPrivateData("collectionCheck", B)
	if err != nil {
		return fmt.Errorf("Failed to get state")
	}
	if Bcheckbytes == nil {
		return fmt.Errorf("Entity not found")
	}
	Bcheck, _ = strconv.Atoi(string(Bcheckbytes))

	Bcheck = Bcheck + X
	Acheck = Acheck - X

	err = ctx.GetStub().PutPrivateData("collectionCheck", A, []byte(strconv.Itoa(Acheck)))
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData("collectionCheck", B, []byte(strconv.Itoa(Bcheck)))
	if err != nil {
		return err
	}

	return nil
}

func (t *SmallBank) TRANSACT_SAVINGS(ctx contractapi.TransactionContextInterface, A string, X int) error {
	var err error
	var Asaving int
	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Asavingbytes, err := ctx.GetStub().GetPrivateData("collectionSaving", A)
	if err != nil {
		return fmt.Errorf("Failed to get state")
	}
	if Asavingbytes == nil {
		return fmt.Errorf("Entity not found")
	}
	Asaving, _ = strconv.Atoi(string(Asavingbytes))
	Asaving = Asaving + X
	fmt.Printf("Asaving = %d\n", Asaving)
	// Write the state back to the ledger
	err = ctx.GetStub().PutPrivateData("collectionSaving", A, []byte(strconv.Itoa(Asaving)))
	if err != nil {
		return err
	}
	return nil
}

func (t *SmallBank) WRITE_CHECK(ctx contractapi.TransactionContextInterface, A string, X int) error {
	var err error
	var Acheck int
	var Asaving int
	var Asum int
	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Acheckbytes, err := ctx.GetStub().GetPrivateData("collectionCheck", A)
	if err != nil {
		return fmt.Errorf("Failed to get state")
	}
	if Acheckbytes == nil {
		return fmt.Errorf("Entity not found")
	}
	Acheck, _ = strconv.Atoi(string(Acheckbytes))

	Asavingbytes, err := ctx.GetStub().GetPrivateData("collectionSaving", A)
	if err != nil {
		return fmt.Errorf("Failed to get state")
	}
	if Asavingbytes == nil {
		return fmt.Errorf("Entity not found")
	}
	Asaving, _ = strconv.Atoi(string(Asavingbytes))
	Asum = Asaving + Acheck
	if Asum < X {
		Acheck = Acheck - X + 1
	} else {
		Acheck = Acheck - X
	}
	err = ctx.GetStub().PutPrivateData("collectionCheck", A, []byte(strconv.Itoa(Acheck)))
	if err != nil {
		return err
	}

	return nil
}

// Query callback representing the query of a chaincode
func (t *SmallBank) BALANCE(ctx contractapi.TransactionContextInterface, A string) (string, error) {
	var err error
	var Acheck int
	var Asaving int
	var Asum int
	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Acheckbytes, err := ctx.GetStub().GetPrivateData("collectionCheck", A)
	if err != nil {
		return "Failed to get state", err
	}
	if Acheckbytes == nil {
		return "Entity not found", nil
	}
	Acheck, _ = strconv.Atoi(string(Acheckbytes))

	Asavingbytes, err := ctx.GetStub().GetPrivateData("collectionSaving", A)
	if err != nil {
		return "Failed to get state", err
	}
	if Asavingbytes == nil {
		return "Entity not found", nil
	}
	Asaving, _ = strconv.Atoi(string(Asavingbytes))
	Asum = Acheck + Asaving
	jsonResp := "{\"Account\":\"" + A + "\",\"Sum\":\"" + strconv.Itoa(Asum) + "\",\"Saving\":\"" + string(Asavingbytes)+ "\",\"Check\":\"" + string(Acheckbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return jsonResp, nil
}

func main() {
	cc, err := contractapi.NewChaincode(new(SmallBank))
	if err != nil {
		panic(err.Error())
	}
	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting SmallBank chaincode: %s", err)
	}
}
