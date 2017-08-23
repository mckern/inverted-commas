# Inverted Commas

## Dumb up smart quotes with a simple plaintext filter

I decided to learn [Go](https://golang.org) because it's 2017 and interpreted languages
have fallen out of vogue. To that end, I started by writing a simple filter to replace
[smart quotes](http://practicaltypography.com/straight-and-curly-quotes.html) with their dumb cousins. And thus, inverted-commas
(`ic`) was born.

## Caveat Lector!

This is an academic project that may or may not be useful to you.
It's a work in progress, and it's expected to be rough around the edges as I learn the
ins and outs of a new language that's built around very different mores and conventions
than Ruby or Bash.

## Usage

`ic` is meant to be run as a simple [plaintext filter](https://en.wikipedia.org/wiki/Pipeline_(Unix)),
accepting input from `STDIN` and printing processed input back across `STDOUT`:

```
$ echo \‟This was some fancy text\” | ./ic
"This was some fancy text"
```

## License

inverted-commas is licensed under the MIT License (MIT), and it's a work in progress.
It was written to scratch an itch, and as such it has some rough edges and will experience churn.
