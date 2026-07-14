package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a product warranty registry
type SmartContract struct {
	contractapi.Contract
}

// Warranty represents a product warranty record
type Warranty struct {
	WarrantyID string `json:"WarrantyID"` // unique warranty id, e.g. "w1"
	Product    string `json:"Product"`    // product name
	Owner      string `json:"Owner"`      // warranty holder name
	Status     string `json:"Status"`     // ACTIVE | CLAIMED
}

// RegisterWarranty creates a new warranty with status "ACTIVE".
// It must fail if a warranty with the same warrantyID already exists.
func (s *SmartContract) RegisterWarranty(ctx contractapi.TransactionContextInterface, warrantyID string, product string, owner string) error {

	return nil
}

// GetWarranty returns the warranty identified by warrantyID.
// It must fail if the warranty does not exist.
func (s *SmartContract) GetWarranty(ctx contractapi.TransactionContextInterface, warrantyID string) (*Warranty, error) {

	return nil, nil
}

// ClaimWarranty sets the warranty's status to "CLAIMED".
// It must fail if the warranty does not exist or is not currently ACTIVE.
func (s *SmartContract) ClaimWarranty(ctx contractapi.TransactionContextInterface, warrantyID string) error {

	return nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		panic("Error creating warranty chaincode: " + err.Error())
	}

	if err := chaincode.Start(); err != nil {
		panic("Error starting warranty chaincode: " + err.Error())
	}
}
