//go:build js && wasm

package main

import (
	"fmt"
	"strconv"
	"syscall/js"
	"time"

	"github.com/joaonsantos/goplayground/wasm/pkg/glife"
)

const (
	cellSize     = 20 // px
	padding      = 2  // px
	worldColSize = 36 // cells
	worldRowSize = 18 // cells
)

func main() {
	worldState := glife.NewWorldState(worldColSize, worldRowSize)
	glife.Seed(worldState, worldColSize, worldRowSize)

	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "world")
	fullWidth := (worldColSize + padding) * cellSize
	fullHeight := (worldRowSize + padding) * cellSize
	canvas.Set("width", strconv.Itoa(fullWidth))
	canvas.Set("height", strconv.Itoa(fullHeight))
	ctx := canvas.Call("getContext", "2d")

	var renderFrame js.Func
	var tmark float64
	var tdiffSum float64
	var tmarkCount float64

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		now := args[0].Float()
		tdiff := now - tmark
		tdiffSum += tdiff
		tmarkCount += 1
		if tmarkCount > 10 {
			avgFrameTime := tdiffSum / tmarkCount
			avgFPS := 1000 / avgFrameTime

			doc.Call("getElementById", "frame-time").Set("innerHTML", fmt.Sprintf("%.01f", avgFrameTime))
			doc.Call("getElementById", "fps").Set("innerHTML", fmt.Sprintf("%.01f", avgFPS))
			tmarkCount = 0
			tdiffSum = 0
		}
		tmark = now

		for i := 0; i < worldRowSize; i++ {
			for j := 0; j < worldColSize; j++ {
				cellStatus := worldState[i][j]

				posX := j * cellSize
				posY := i * cellSize

				ctx.Set("fillStyle", "#000000")
				if cellStatus == 1 {
					ctx.Set("fillStyle", "#00c800")
				}

				realCellSize := cellSize - padding
				ctx.Call("fillRect", posX, posY, realCellSize, realCellSize)
			}
		}

		// keep render loop going
		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})

	// start render loop
	js.Global().Call("requestAnimationFrame", renderFrame)

	// logic update loop
	for {
		worldState = glife.WorldTick(worldState, worldColSize, worldRowSize)
		time.Sleep(260 * time.Millisecond)
	}
}
