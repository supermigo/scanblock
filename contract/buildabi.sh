#!/bin/bash
#this cmd  is make abi file to go file
abigen --abi ido.abi --pkg contract --type ido --out ido.go
abigen --abi idfactory.abi --pkg contract --type idfactory --out idfactory.go