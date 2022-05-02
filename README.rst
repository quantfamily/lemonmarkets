Lemonmarkets (Unofficial) Go Client
===================================

.. image:: https://github.com/quantfamily/lemonmarkets/workflows/tests/badge.svg
    :target: https://github.com/quantfamily/lemonmarkets/actions?query=workflow%3Atests

.. image:: https://codecov.io/gh/quantfamily/lemonmarkets/branch/main/graph/badge.svg
    :target: https://codecov.io/gh/quantfamily/lemonmarkets
    :alt: Code coverage Status

.. image:: https://goreportcard.com/badge/github.com/quantfamily/lemonmarkets
    :target: https://goreportcard.com/report/github.com/quantfamily/lemonmarkets

.. image:: https://pkg.go.dev/badge/github.com/quantfamily/lemonmarkets.svg
    :target: https://pkg.go.dev/github.com/quantfamily/lemonmarkets


Background
----------

`LemonMarkets <https://www.lemon.markets>`__  is an API for stock trading and market data.

This library acts as client library to simplify the usage of LemonMarkets when using Go code.
This library is a Unofficial library and therefor the usage cannot always be guaranteed.
Nevertheless, the goal is to have full support for both Rest API as well as Streaming.

Usage (Trading module)
----------------------

Example getting your account information

.. code-block:: golang

    package main

    import (
        "fmt"

        "github.com/quantfamily/lemonmarkets/trading"
    )

    func main() {
        client := trading.NewClient("YOUR_API_KEY", trading.PAPER)
        account := client.GetAccount()

        fmt.Println(account)
    }

Example getting orders placed

.. code-block:: golang

    package main

    import (
        "fmt"

        "github.com/quantfamily/lemonmarkets/trading"
    )

    func main() {
        client := trading.NewClient("YOUR_API_KEY", trading.PAPER)
        orders := client.GetOrders(nil) // nil to not filter (i.e receive all)

        for order := range orders {
            fmt.Println(order.Data)
        }
    }


Usage (Trading module)
----------------------

Example getting OHLC Per Day of a share

.. code-block:: golang

    package main

    import (
        "fmt"

        "github.com/quantfamily/lemonmarkets/market_data"
    )

    func main() {
        client := market_data.NewClient("YOUR_API_KEY")

        ohlcQuery := market_data.GetOHLCQuery{ISIN: []string{"SE0000115446"}} // Get Volvo B
        ohlcs := client.GetOHLCPerDay(&ohlcQuery)

        for ohlc := range ohlcs {
            fmt.Println(ohlc.Data)
        }
    }

