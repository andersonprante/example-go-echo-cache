# Simple example

This example show how to use `gitsight/go-echo-cache` as a cache for echo and add Header control that's allow to control when route not should be cached.

## CURL route with cache
`curl localhost:8080/`

## CURL route with NO cache
Just need send header key (cache) with false value

`curl -H cache:false localhost:8080/`

