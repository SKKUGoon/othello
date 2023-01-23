package game

type GameError struct {
	Code    string
	Message string
}

func (e *GameError) Error() string {
	return e.Code + ", " + e.Message
}
