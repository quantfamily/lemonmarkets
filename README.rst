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

Usage
-----

Example getting Daily data for stock[1,2,3]

.. code-block:: golang

    client := lemonmarkets.NewClient(lemonmarkets.DATA, "API_KEY")

    isins := []string{"ISIN_STOCK1", "ISIN_STOCK2", "ISIN_STOCK3"}
    query := GetOHLCQuery{ISIN: isins}

    daily_data, err := GetOHLCPerDay(client, *query)
    if err != nil {
        return err
    }

Example placing order for 10 shares of asset with ISIN 123456789 on paper environment

.. code-block:: golang

    client := lemonmarkets.NewClient(lemonmarkets.PAPER, "API_KEY")

    order := Order{ISIN: "123456789", Quantity: 10, Side: "Buy"}

    created_order, err := CreateOrder(client, *order)
    if err != nil {
        return err 
    }

