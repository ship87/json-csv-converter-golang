# Convert JSON to CSV using golang

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub stars](https://img.shields.io/github/stars/ship87/json-csv-converter-golang.svg)](https://github.com/ship87/json-csv-converter-golang/stargazers)
[![Go Report Card](https://goreportcard.com/badge/github.com/ship87/json-csv-converter-golang)](https://goreportcard.com/report/github.com/ship87/json-csv-converter-golang)

## Build and run

`CGO_ENABLED=0 GOOS=linux go build -tags netgo -ldflags "-s" -a -installsuffix cgo -o service && docker build -t service . && docker run -e -it -p 8080:8080 -v /home/{user}/{app folder}/download:/download service`

## Test

`curl -X POST -H 'Content-Type:application/json' --data '[{"number":1,"columns":["Line1", "Test1", "Lorem ipsum dolor sit amet, consectetuer adipiscing elit,"]},{"number":2,"columns":["Line2", "Test2", "sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat."]},{"number":3,"columns":["Line3", "Test3", "Ut wisi enim ad minim veniam,"]},{"number":4,"columns":["Line4", "Test4", "quis nostrud exerci tation ullamcorper suscipit lobortis nisl ut aliquip ex ea commodo consequat."]},{"number":5,"columns":["Line5", "Test5", "Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat,"]},{"number":6,"columns":["Line6", "Test6", "vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi."]},{"number":7,"columns":["Line7", "Test7", "Lorem ipsum dolor sit amet, consectetuer adipiscing elit,"]},{"number":8,"columns":["Line8", "Test8", "sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat."]},{"number":9,"columns":["Line9", "Test9", "Ut wisi enim ad minim veniam,"]},{"number":10,"columns":["Line10", "Test10", "quis nostrud exerci tation ullamcorper suscipit lobortis nisl ut aliquip ex ea commodo consequat."]},{"number":11,"columns":["Line11", "Test11", "Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat,"]},{"number":12,"columns":["Line12", "Test12", "vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi."]}]' http://localhost:8080 -vvv`
