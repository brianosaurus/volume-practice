# Volume Practice Code

To use
```
go get
go run .
```

Then in another shell use it like this
```
curl -H 'Content-Type: application/json' -X POST -d '[["FOO", "BAR"],["BAR", "SFO"]]' http://localhost:8080/api/v1/flightPath
{"response":["FOO","SFO"],"status":200}

curl -H 'Content-Type: application/json -X POST -d '[["FOO", "BAR"],["BAR", "SFO"], ["XXX","BLEH"]]' http://localhost:8080/api/v1/flightPath
{"response":["",""],"status":200}
```

If the response is an empty set then there was a problem with the inputs. I haven't done everything to make error reporting pretty since this question was mainly to sort whether or not someone realizes that O(2^N) is not how to solve this.

My algorithm is O(2N)


