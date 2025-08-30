package maze

func (m *Maze) Start() chan Maze {

	signal := make(chan Maze, 100)

	go func() {
		m.Create(signal)
		m.FindExit(signal)
		m.FindExit(signal)
		m.Clear(signal)
		close(signal)
	}()

	return signal
}
