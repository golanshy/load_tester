# load_tester

## Goal: 
Allow exercising API calls for load testing

## Usage:
cd {repo_folder}/src

go run main.go

## URL:
http://localhost:8080/load_test/test

## Body:
{
    "test_type" : "Order",
    "repetitions": 1,
    "delay_in_milliseconds": 100
}

## JSON Params:
repetitions: How many times the API call will be made

delay_in_milliseconds: Deley in milliseconds between calls

test_type could be GetEntireMenu / Order / GetEntireMenuAndOrder

GetEntireMenu: Run the get full menu API call

Order: Run the place order API call

GetEntireMenuAndOrder: Does both

## curl:

curl --location --request POST 'http://localhost:8080/load_test/test' \
--header 'Content-Type: application/json' \
--data-raw '{
"test_type" : "Order",
"repetitions": 1,
"delay_in_milliseconds": 100
}'

