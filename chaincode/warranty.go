package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a product warranty registry
type SmartContract struct {
	contractapi.Contract
}

// Warranty represents a product warranty record
type Warranty struct {
	WarrantyID string `json:"WarrantyID"`
	Product    string `json:"Product"`
	Owner      string `json:"Owner"`
	Status     string `json:"Status"`
}

const (
	statusActive  = "ACTIVE"
	statusClaimed = "CLAIMED"
)

// RegisterWarranty creates a new warranty with status "ACTIVE".
func (s *SmartContract) RegisterWarranty(ctx contractapi.TransactionContextInterface, warrantyID string, product string, owner string) error {
	existing, err := ctx.GetStub().GetState(warrantyID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if existing != nil {
		return fmt.Errorf("warranty %s already exists", warrantyID)
	}

	warranty := Warranty{
		WarrantyID: warrantyID,
		Product:    product,
		Owner:      owner,
		Status:     statusActive,
	}

	bytes, err := json.Marshal(warranty)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(warrantyID, bytes)
}

// GetWarranty returns the warranty identified by warrantyID.
func (s *SmartContract) GetWarranty(ctx contractapi.TransactionContextInterface, warrantyID string) (*Warranty, error) {
	data, err := ctx.GetStub().GetState(warrantyID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if data == nil {
		return nil, fmt.Errorf("warranty %s does not exist", warrantyID)
	}

	var warranty Warranty
	if err := json.Unmarshal(data, &warranty); err != nil {
		return nil, err
	}
	return &warranty, nil
}

// ClaimWarranty sets the warranty's status to "CLAIMED".
func (s *SmartContract) ClaimWarranty(ctx contractapi.TransactionContextInterface, warrantyID string) error {
	warranty, err := s.GetWarranty(ctx, warrantyID)
	if err != nil {
		return err
	}
	if warranty.Status != statusActive {
		return fmt.Errorf("warranty %s is not ACTIVE (current status: %s)", warrantyID, warranty.Status)
	}

	warranty.Status = statusClaimed

	bytes, err := json.Marshal(warranty)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(warrantyID, bytes)
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
