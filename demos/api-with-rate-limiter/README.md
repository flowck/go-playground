# `api-with-rate-limiter`

This demo uses [`go-chi/chi`](https://github.com/go-chi/chi) as a custom router and [`go-chi/httprate`](https://github.com/go-chi/httprate) to limit the amount of http requests a single client can perform per minute. 