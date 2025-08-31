package maze

import (
	"bytes"
	"image"
	"log/slog"
)

func (m *Maze) Start() chan image.Image {

	signal := make(chan Maze, 10)

	go func() {
		m.Create(signal)
		m.FindExit(signal)
		m.FindExit(signal)
		m.Clear(signal)
		close(signal)
	}()

	img := make(chan image.Image, 100)

	go func() {
		for {
			select {
			case mazeUpdate, open := <-signal:
				if !open {
					break
				}
				slog.Info("sending maze")
				buffer := new(bytes.Buffer)
				RenderMaze(&mazeUpdate, buffer)
				decode, _, err := image.Decode(bytes.NewReader(buffer.Bytes()))
				if err != nil {
					return
				}

				img <- decode

			}
		}
	}()

	return img
}
