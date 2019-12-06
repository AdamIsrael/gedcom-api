# A Go-based microservice for parsing GEDCOM files

[![Build Status](https://travis-ci.org/AdamIsrael/gedcom-api.svg?branch=master)](https://travis-ci.org/AdamIsrael/gedcom-api)
[![codecov](https://codecov.io/gh/AdamIsrael/gedcom-api/branch/master/graph/badge.svg)](https://codecov.io/gh/AdamIsrael/gedcom-api)
[![Report Card](https://goreportcard.com/badge/github.com/adamisrael/gedcom-api)](https://goreportcard.com/report/github.com/adamisrael/gedcom-api)

## Building

```bash
go build -o gedcom-api cmd/gedcom-api/main.go && ./gedcom-api -verbose
```

## Endpoints

| Endpoint      | Description   | Method | Status |
| ------------- |:-------------:| -----:| -----:|
| /upload | Upload a GEDCOM file | PUT | Done |
| /upload/{gedcom id} | Get information about an uploaded GEDCOM | GET | TBD |
| /individual{gedcom id} | Get a list of all individuals | GET | TBD |
| /individual/{gedcom id}/1 | Get a list of individual 1 | GET | Done |
| /surname/{gedcom id} | Get a dictionary of surnames and how frequently they occur | GET | Done |


### Upload

To upload a GEDCOM:

```bash
$ curl -F 'gedcom=@/home/stone/Downloads/AdamIsrael.ged' http://localhost:8000/upload
{"status":"OK","uuid":"8174da41-183f-11ea-a9f3-28d244449944"}
```

To get information about an uploaded GEDCOM file:

```bash
$ curl http://localhost:8000/gedcom/8174da41-183f-11ea-a9f3-28d244449944
{}
```


curl http://localhost:8000/individual/8174da41-183f-11ea-a9f3-28d244449944/1