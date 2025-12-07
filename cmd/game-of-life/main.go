// Package main
package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gentmaks/game-of-life/internal/game"
)

const (
	screenWidth   = 800
	screenHeight  = 800
	cellSize      = 20
	buttonWidth   = 200
	buttonHeight  = 50
	buttonX       = (screenWidth - buttonWidth) / 2
	buttonY       = screenHeight - 80
	resumeButtonX = buttonX
	resumeButtonY = buttonY
)

func main() {
	board := game.BoardInit(40, 40)
	width, height := board.GetDimensions()

	// Calculate grid dimensions to fit in screen
	gridWidth := width * cellSize
	gridHeight := height * cellSize
	gridOffsetX := (screenWidth - gridWidth) / 2
	gridOffsetY := (screenHeight - gridHeight - 100) / 2

	rl.InitWindow(screenWidth, screenHeight, "Game of Life")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	running := true

	for !rl.WindowShouldClose() {
		// Handle button clicks
		mousePos := rl.GetMousePosition()
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			// Stop button (when running)
			if running {
				if mousePos.X >= float32(buttonX) && mousePos.X <= float32(buttonX+buttonWidth) &&
					mousePos.Y >= float32(buttonY) && mousePos.Y <= float32(buttonY+buttonHeight) {
					running = false
				}
			} else {
				// Resume button (when stopped)
				if mousePos.X >= float32(resumeButtonX) && mousePos.X <= float32(resumeButtonX+buttonWidth) &&
					mousePos.Y >= float32(resumeButtonY) && mousePos.Y <= float32(resumeButtonY+buttonHeight) {
					running = true
				}
			}
		}

		if running {
			board.Advance()
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				cellX := float32(gridOffsetX + x*cellSize)
				cellY := float32(gridOffsetY + y*cellSize)

				if board.GetCellState(x, y) == 1 {
					rl.DrawRectangle(int32(cellX), int32(cellY), cellSize, cellSize, rl.Black)
				} else {
					rl.DrawRectangle(int32(cellX), int32(cellY), cellSize, cellSize, rl.White)
				}

				// Draw grid lines
				rl.DrawRectangleLines(int32(cellX), int32(cellY), cellSize, cellSize, rl.LightGray)
			}
		}

		if running {
			buttonColor := rl.DarkGray
			if mousePos.X >= float32(buttonX) && mousePos.X <= float32(buttonX+buttonWidth) &&
				mousePos.Y >= float32(buttonY) && mousePos.Y <= float32(buttonY+buttonHeight) {
				buttonColor = rl.Gray
			}
			rl.DrawRectangle(buttonX, buttonY, buttonWidth, buttonHeight, buttonColor)
			rl.DrawRectangleLines(buttonX, buttonY, buttonWidth, buttonHeight, rl.Black)
			text := "Stop Simulation"
			textWidth := rl.MeasureText(text, 20)
			rl.DrawText(text, buttonX+(buttonWidth-textWidth)/2, buttonY+15, 20, rl.White)
		} else {
			// Draw resume button
			buttonColor := rl.DarkGreen
			if mousePos.X >= float32(resumeButtonX) && mousePos.X <= float32(resumeButtonX+buttonWidth) &&
				mousePos.Y >= float32(resumeButtonY) && mousePos.Y <= float32(resumeButtonY+buttonHeight) {
				buttonColor = rl.Green
			}
			rl.DrawRectangle(resumeButtonX, resumeButtonY, buttonWidth, buttonHeight, buttonColor)
			rl.DrawRectangleLines(resumeButtonX, resumeButtonY, buttonWidth, buttonHeight, rl.Black)
			text := "Resume Simulation"
			textWidth := rl.MeasureText(text, 20)
			rl.DrawText(text, resumeButtonX+(buttonWidth-textWidth)/2, resumeButtonY+15, 20, rl.White)
		}

		rl.EndDrawing()
	}
}
