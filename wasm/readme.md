# Go WASM Game of Life

Implementation of the famous game of life running on the browser canvas and powered by Go Webassembly support.

Conway's game of life is a cellular automaton devised by the British mathematician John Horton Conway in 1970. For more information please visit https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life.

There are many variations, but in this case:
- 0 represents a dead cell
- 1 represents a living cell


## Running

To run:
```sh
$ make
```

This will build the go program into a webassebly wasm file and serve the required files on `http://localhost:8080`.

Tested on Go 1.24.

# Example

![Conway's Game of Life Vizualization](https://media.giphy.com/media/5Rtir4sTGj0mlQ9J65/giphy.gif)