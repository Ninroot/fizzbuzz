# Fizz-buzz REST server

This is a fizz-buzz REST server written in Golang made for the technical test at Leboncoin.

## Exercice

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz". The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...".

Your goal is to implement a web server that will expose a REST API endpoint that:
- Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
- Ready for production
- Easy to maintain by other developers

__Bonus__: add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request"

## Prerequisites

- go
- make (optional)

## Build / Run

To build and run server:
```bash
make && ./fizzbuzz
```

To run the test suit:
```bash
make test
```

Environment variables:
- `PORT`: port (default 8080)
- `STATS_WINDOW`: size of the requests window to be saved (default 10000)

## API

### Fizzbuzz

Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2. Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

`POST /fizzbuzz`

Input:
- `int1`: int, required
- `int2`: int, required
- `limit`: positive int, required
- `str1`: string, required
- `str2`: string, required

Output:
- Array of string or
- Error with `msg` field

Example:
```sh
curl -X POST 0.0.0.0:8080/fizzbuzz -H 'content-type: application/json' -d '{"int1":3, "int2":5,"limit":15,"str1":"Fizz","str2":"Buzz"}'
["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
```

### Statistiques

Most frequent request has been sent. Only `STATS_WINDOW` different requests are saved, the least frequent one are remove if the window is full. Note that stats are store in-memory and therefore reflect local requests.

`GET /stats`

Output:
- Dictionnary `{<limit> <int1> <int2> <str1> <str2>}` as key and number of occurence as value

Example:
```bash
curl 0.0.0.0:8080/stats
{"{100 3 5 Fizz Buzz}":1,"{15 3 5 Fizz Buzz}":1}
```
