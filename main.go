package main

type (

	// func(ctx context.Context, e *scheduler.Event) ( context.Context, *scheduler.Event, error)
	// 향후 아래와 같이 구성되어야 할듯
	// Handlers[] func(ctx context.Context, e *scheduler.Event) ( context.Context, *scheduler.Event, error)
	Handlers []func() (string, error)
	//Handlers[] interface{}
	// Handlers1 func(funcs ...Function) *Chainable
	// 추후에는 scheduler.Event_Type 로 해야 하지만 테스트를 위해 임시로 이렇게 함.
	MapHandlers map[int]Handlers
	// Handlers interface{}
	// MapHandlers map[int]Handlers
	//MapHandlers map[scheduler.Event_Type]Handlers

	Error string

	Function interface{}
)

const (
	// 향후 errors.go 함수제작하기
	ErrNotFoundA = Error("test a : error not founded")
	ErrNotFoundB = Error("test b : error not founded")
	ErrNotFoundC = Error("test c : error not founded")
	ErrNotFoundD = Error("test d : error not founded")
	ErrNotFoundE = Error("test e : error not founded")
	ErrNotFoundF = Error("test f : error not founded")
	ErrNotFoundG = Error("test g : error not founded")

	ErrA = Error("test a : error founded")
	ErrB = Error("test b : error founded")
	ErrC = Error("test c : error founded")
	ErrD = Error("test d : error founded")
	ErrE = Error("test e : error founded")
	ErrF = Error("test f : error founded")
	ErrG = Error("test g : error founded")
)

func main() {
	// TDD 방식으로 최대한 method 단위를 쪼개서 제작한다.

	s := EventHandlers()
	EventLooP(1, s)
}

func (err Error) Error() string { return string(err) }

func EventHandlers() (handlers *MapHandlers) {

	handlers = &MapHandlers{
		1: Handlers{
			testA(),
			//testB(),
		},

		/*2: Handlers{
			testC(),
			testD(),
		},

		3: Handlers{
			testE(),
			testF(),
			testG(),
		},*/
	}
	return
}

// link represents a function in a chain
// fn  은 향후 func(ctx context.Context, e *scheduler.Event) ( context.Context, *scheduler.Event, error) 로 되어야 함.
// link, Chainabl 을 수정해야함. chain 이 될때마다 결과값들을 저장해두어야 한다. 이 결과 값이 서로 연결이 안되게 하고 리턴만 하고, 다음으로 넘어가는 구조를 가져야 한다.
type link struct {
	fn          func() (string, error) // 비교 테스트를 하자
	handleError bool                   // ??
}

// Chainable defines the chain logic.
type Chainable struct {
	links []link
}

// Chain appends a new function to the chain, with error
// handling enabled
func (c *Chainable) Chain(funcs ...Function) *Chainable {
	if c == nil {
		n := &Chainable{}
		n.chainFuncs(funcs, true)
		return n
	}

	c.chainFuncs(funcs, true)
	return c
}

// chainFuncs add new functions to the chain, creating the underlying link
func (c *Chainable) chainFuncs(funcs []Function, handleError bool) {
	for _, f := range funcs {
		c.links = append(c.links, link{
			fn:          f.(func() (string, error)),
			handleError: handleError,
		})
	}
}

/*func EventHandlers() (handlers *MapHandlers) {
	return &MapHandlers{
		scheduler.Event_SUBSCRIBED: Handlers{
			testA(),
			testB(),
		},

		scheduler.Event_OFFERS: Handlers{
			testC(),
			testD(),
		},

		scheduler.Event_UPDATE: Handlers{
			testE(),
			testF(),
			testG(),
		},
	}
}*/
// 향후 이렇게 가야 하지만 테스트를 위해 다른 파라미터는 일단 제거
// 에러가 발생하면 해당 에러를 내보내고 정지??
// 에러가 발생하더라도 에러 메세지만 리턴하고, 해당 함수는 정지하지 않는다. 다른 이벤트들을 처리한다.
// 이벤트 핸들러의 타입은 func(ctx context.Context, e *scheduler.Event) (err error) 이렇게 정한다.
// 추가적인 파라미터가 생길경우는
/*
	func SomeEventHandler(param ...params) {
		return func(ctx context.Context, e *scheduler.Event) ( context.Context, *scheduler.Event, error){

			param...

			return ctx, event, err
		}
}
*/

func EventLooP(n int, handlers *MapHandlers) (err error) {
	return nil
}

/*func EventLooP(ctx context.Context, event scheduler.Event, handlers *Handlers) (err error) {
	return nil
}*/

func testA() func() (string, error) {
	return func() (string, error) {
		str := "test A run"
		err := ErrNotFoundA
		return str, err
	}
}
func testB() (str string, err error) {
	str = "test B run"
	err = ErrB
	return
}
func testC() (str string, err error) {
	str = "test C run"
	err = ErrNotFoundC
	return
}
func testD() (str string, err error) {
	str = "test D run"
	err = ErrD
	return
}
func testE() (str string, err error) {
	str = "test E run"
	err = ErrE
	return
}
func testF() (str string, err error) {
	str = "test F run"
	err = ErrNotFoundF
	return
}
func testG() (str string, err error) {
	str = "test G run"
	err = ErrNotFoundG
	return
}
