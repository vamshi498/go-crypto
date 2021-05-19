# go-crypto

This is used to fetch real time crypto prices

# Run the project 
```go run main.go```
# This contains two endpoints

1.curl http://localhost:8081/currency
 This gives all the data of currencies available

2. curl http://localhost:8081/currency/{symbol}
 This gives data for a given symbol  
ex: curl http://localhost:8081/currency/BTCUSD
should give you something like 
  ```{"symbol":"BTCUSD","FullName":"","Ask":"38393.06","Bid":"38322.84","Last":"38333.08","Open":"42337.56","Low":"29257.95","High":"43575.35","FeeCurrency":""}```

