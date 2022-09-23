# gin-trade-engine-example
The project for basic implementation of trade engine.
## Project structure
```
.
├── channel
│   └── order_channel.go // message queue for order
├── controller
|   └── controller.go // to handle the http request
├── domain
|   └── domain.go // define domain objects
├── engine
|   ├── engine_test.go // for testing the trade engine
|   └── engine.go // trade engine for matching the sell order and buy order
├── repository
|   └── repository.go // repository for orders, trades
├── util
|   └── util.go // utillities
├── main.go // for startup the application
...
```

## Design
### Flow for creating an order
1. Client send a request to POST /api/orders
2. Request handle by controller and controller will send a message to message queue
3. The listner receive the message and call the trade engine to create an order
4. Trade engine will try to match an order with sells or buys depends on order type
5. A trade would be created when order matched.

### Business logic of trade matching
The two scenarios here, 1) an buy order try to match the sells or, 2) a sell order try to match the buys.

#### Scenario 1 - a buy order request matched a sell:
1. Request a buy order
2. Loop over the sells
3. Create a trade when both price and quantity matched on buy and sell

#### Scenario 2 - a buy order request matched more than one sell:
1. Request a buy order price = 100, quantity = 4
2. There are 3 sells with price = 100 in the sells list, and the quantities are 1, 1, 2
4. Loop over the sells
3. Three trade would be created for the buy order

#### Scenario 3 - mo matched on a buy order
1. Request a buy order
2. Loop over the sells
3. No matched found then add the buy order to buy order list

## API Doc
| Endpoint | Response | Request  |
| --- | ----------- | --------- |
| POST /api/orders | order object | order object |
| GET /api/buys | list of orders | N/A |
| GET /api/sells | list of sells | N/A |
| GET /api/trades | list of trades | N/A |

``` shell
 # create an order
 curl --location --request POST 'http://localhost:8080/api/orders' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
--header 'Authorization: Basic dXNlcjpwYXNzd29yZA==' \
--data-raw '{
    "orderType": "sell",
    "priceType": "limit",
    "price": 20,
    "quantity": 3
}'

# list buys
curl --location --request GET 'http://localhost:8080/api/buys' \
--header 'Accept: application/json' \
--header 'Authorization: Basic dXNlcjpwYXNzd29yZA=='

# list sells
curl --location --request GET 'http://localhost:8080/api/sells' \
--header 'Accept: application/json' \
--header 'Authorization: Basic dXNlcjpwYXNzd29yZA=='

# list trades
curl --location --request GET 'http://localhost:8080/api/trades' \
--header 'Accept: application/json' \
--header 'Authorization: Basic dXNlcjpwYXNzd29yZA=='
```