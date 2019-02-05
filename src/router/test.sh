#!/bin/bash
go test --test.run TestAddressHttpFunc1 -v
go test --test.run TestAddressHttpFunc2 -v
go test --test.run TestGetAddressCsv -v
go test --test.run TestPostAddressCsv1 -v
go test --test.run TestPostAddressCsv2 -v
