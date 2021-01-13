package states

type State string

const (
	Waiting    State = "waiting"
	Collecting State = "collecting"
	Sending    State = "sending"
	Error      State = "error"
)

func (s State) String() string {
	return string(s)
}

var currentState = Waiting

func SetState(state State) {
	currentState = state
}

func GetState() State {
	return currentState
}
