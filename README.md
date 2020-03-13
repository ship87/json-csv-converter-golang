# Convert JSON to CSV using golang

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub stars](https://img.shields.io/github/stars/ship87/json-csv-converter-golang.svg)](https://github.com/ship87/json-csv-converter-golang/stargazers)
[![Go Report Card](https://goreportcard.com/badge/github.com/ship87/json-csv-converter-golang)](https://goreportcard.com/report/github.com/ship87/json-csv-converter-golang)

## Build and run

`sudo chmod -R 777 public/download &&\
CGO_ENABLED=0 GOOS=linux go build -tags netgo -ldflags "-s" -a -installsuffix cgo -o service &&\
 docker build -t service . &&\
  docker run -e -it -p 8081:8081 -v /home/{user}/{app folder}/public/download:/public/download service`

## Test

`curl -X POST -H 'Content-Type:application/json' --data '[{"number":1,"columns":["Line1", "Test1", "Lorem ipsum dolor sit amet, consectetuer adipiscing elit,"]},{"number":2,"columns":["Line2", "Test2", "sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat."]}]' http://localhost:8081 -vvv`
