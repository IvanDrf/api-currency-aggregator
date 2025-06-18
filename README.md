# API for currency-aggregator

## Methods
- GET
  ```shell
  curl -X GET http://localhost:8080/sources
  curl -X GET http://localhost:8080/currency
  ```

- ### /sources - returns JSON currency information sources
  ```json
  {
    "binance": true,
    "cbr": true,
    "coingecko": true
  }
  ```
  
- ### /currency - returns JSON information about the currency
  - EUR
    ```shell
    curl -X GET http://localhost:8080/currency?currency=EUR
    ```
    
    ```json
    {
      "currency": "EUR",
      "average_rate": 90.476,
      "sources": [
          {
              "source": "cbr",
              "rate": 90.755
          },
          {
              "source": "coingecko",
              "rate": 90.196
          }
      ]
    }
    ```
    
  - USD
    ```shell
    curl -X GET http://localhost:8080/currency?currency=USD
    ```
  
    ```json
    {
      "currency": "USD",
      "average_rate": 82.739,
      "sources": [
          {
              "source": "coingecko",
              "rate": 78.402
          },
          {
              "source": "binance",
              "rate": 91.1
          },
          {
              "source": "cbr",
              "rate": 78.714
          }
      ]
    }
    ``` 

***

## Supported currencies
- USD/RUB
- EUR/RUB

***

## Technology used
- Golang
- Web-framework Echo 
    
