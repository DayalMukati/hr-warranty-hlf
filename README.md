# hr-warranty-hlf

Solution repository for the **Warranty Registry** Hyperledger Fabric chaincode
challenge (NPCI / HackerRank, Basic).

Standard Fabric **test-network** plus a chaincode skeleton at
[`chaincode/warranty.go`](chaincode/warranty.go). Cloned into the candidate's
environment by the HackerRank Setup Script (via [`setup.sh`](setup.sh)).

## Candidate task
1. Implement the functions in `chaincode/warranty.go`.
2. Deploy: `cd test-network && ./network.sh deployCC -ccn warrantycc -ccp ../chaincode -ccl go`
3. Register w1 (Laptop, Alice), then claim it.

---

Authored by **Dayal Mukati** — [dayalmukati.com](https://dayalmukati.com)
