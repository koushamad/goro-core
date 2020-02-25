package exception

type (
	Exception struct {
		Message string
		Code int
		Err error
	}
)

func Init(Message string, Code int, Err error) Exception  {
	return Exception{Message:Message, Code:Code, Err:Err}
}
