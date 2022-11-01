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


Usage (Market Data Module)
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

Usage (Streaming Module)
----------------------

Streaming is based on "Ably Realtime", the SDK itself does not contain the library.
Users who wish to implement live streaming of data from lemon.markets can obtain a token and implement the functionality as example below

.. code-block:: golang
    
    package main

    import (
        "context"
        "fmt"
        "time"

        "github.com/ably/ably-go/ably"
        "github.com/quantfamily/lemonmarkets/streaming"
    )

    func main() {
        client := streaming.NewClient("YOUR_API_KEY")

        token := client.GetToken()
        if token.Error != nil {
            panic(token.Error)
        }

        // Get a connection to ably using token from Lemon.markets
        conn, err := ably.NewRealtime(ably.WithToken(token.Data.Token))
        if err != nil {
            panic(err)
        }

        // Get main channel where we will receive quotes and just print every message in a callback function
        ch := conn.Channels.Get(token.Data.UserID)
        stop, err := ch.SubscribeAll(context.TODO(), func(m *ably.Message) {
            fmt.Println(m.Data)
        })
        if err != nil {
            panic(err)
        }

        // Get a subscriptions- channel where we will publish what we want to listen to, ISINs that are comma separated
        subCh := conn.Channels.Get(fmt.Sprintf("%s.subscriptions", token.Data.UserID))
        err = subCh.Publish(context.TODO(), "isins", "US64110L1061,US88160R1014")
        if err != nil {
            panic(err)
        }

        // Sleep just to get some events before we stop the callback function and exit.
        time.Sleep(time.Second * 5)
        stop()
    }
