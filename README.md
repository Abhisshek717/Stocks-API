# Stocks-API
repo is a simple REST API that allows you to retrieve, create, update, and delete stock data. The API is written in Go and uses PostgreSQL as the database.

To use the API, you will need to first install Go and PostgreSQL. Once you have installed Go and PostgreSQL, you can clone the Anu-Ra-g / Stocks-API GitHub repo and run the following command to start the API:

```
go run main.go
```

The API will be listening on port 8080. You can then use a tool like Postman to make requests to the API. For example, the following request will retrieve all stocks from the database:


```
GET http://localhost:8080/stocks
```

The following request will create a new stock in the database:

```
POST http://localhost:8080/stocks

{
  "name": "Apple",
  "company": "Apple Inc.",
  "price": 100
}
```

The following request will update an existing stock in the database:

```
PUT http://localhost:8080/stocks/1

{
  "name": "Google",
  "company": "Google Inc.",
  "price": 200
}
```

The following request will delete a stock from the database:

```
DELETE http://localhost:8080/stocks/1
```
