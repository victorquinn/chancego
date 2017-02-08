# Chance Go

This is a quasi-port of my other library ChanceJS to Go.

I say quasi-port because as of right now it doesn't adhere to the same principles/ideals.

In the JavaScript version I used a Mersenne Twister under the hood as the source of random.

For Go, that doesn't seem to make as much sense. Instead, we are using the `math/rand` package as the source of random.

Also, this is a WIP so it is lacking many of the generators that the JS version currently has. Some of them will make sense to port over (e.g. address) but some less so.

Particularly all of the numerical generators make less sense in the statically typed Go than they did in JavaScript and the `math/rand` package already has methods for each of the Go numeric types. We may wrap them in this library for ease, but for now they are omitted.

[![Chance Logo](http://chancejs.com/logo.png)](http://chancejs.com)
