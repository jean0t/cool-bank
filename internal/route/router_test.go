package route

import (
	"net/http"
	"net/http/httptest"

	"testing"
	"reflect"
)

func TestAddHandler(t *testing.T) {
	var tests = []struct{
		title string
		input []map[string]http.HandlerFunc
		expectedOutput map[string]http.HandlerFunc
		expectFail bool
	} {
		{
			title: "Only one handleFunc",
			input: []map[string]http.HandlerFunc{
				{
					"/test": func(w http.ResponseWriter, r *http.Request){},
				},
			},
			expectedOutput: map[string]http.HandlerFunc{
					"/test": func(w http.ResponseWriter, r *http.Request){},
					},
			expectFail: false,
		},
		{
			title: "Adding more than one handleFunc",
			input: []map[string]http.HandlerFunc{
				{
					"/test": func(w http.ResponseWriter, r *http.Request){},
				},
				{
					"/another": func(w http.ResponseWriter, r *http.Request){},
				},
				{
					"/ping": func(w http.ResponseWriter, r *http.Request){},
				},
			},
			expectedOutput: map[string]http.HandlerFunc{
					"/test": func(w http.ResponseWriter, r *http.Request){},
					"/another": func(w http.ResponseWriter, r *http.Request){},
					"/ping": func(w http.ResponseWriter, r *http.Request){},
					},
			expectFail: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T){
			var handler Handler = Handler{handlers: make(map[string]http.HandlerFunc)}

			for _, element := range tt.input {
				for route, function := range element {
					handler.AddHandler(route, function)
				}
			}

			if reflect.DeepEqual(tt.expectedOutput, handler.handlers) && !tt.expectFail {
				t.Fatal("Error adding handlers with the function AddHandler")
			}
		})
	}

}


func TestGetHandler(t *testing.T) {

	var tests = []struct{
		title string
		expectedAnswer map[string]string // map a router to an answer
		inputHandlers []map[string]http.HandlerFunc
		expectFail bool
	}{
		{
			title: "single get request",
			expectedAnswer: map[string]string {
					"/test": "test succeeded",
			},
			inputHandlers: []map[string]http.HandlerFunc{
				{
					"/test": func(w http.ResponseWriter, r *http.Request) {
						w.Write([]byte("test succeeded"))
					},
				},
			},
			expectFail: false,
		},
		{
			title: "multiple get request",
			expectedAnswer: map[string]string {
					"/test": "test succeeded",
					"/ping": "pong pong",
					"/echo": "heellooo",
			},
			inputHandlers: []map[string]http.HandlerFunc{
				{
					"/test": func(w http.ResponseWriter, r *http.Request) {
						w.Write([]byte("test succeeded"))
					},
				},
				{
					"/ping": func(w http.ResponseWriter, r *http.Request) {
						w.Write([]byte("pong pong"))
					},
				},
				{
					"/echo": func(w http.ResponseWriter, r *http.Request) {
						w.Write([]byte("heellooo"))
					},
				},
			},
			expectFail: false,
		},
	}

	for _, tt := range tests {
		var (
			handler Handler
			routesRegistered []string
			handlerToTest http.HandlerFunc
			req *http.Request
			w *httptest.ResponseRecorder
			response *http.Response
		)

		handler = Handler{handlers: make(map[string]http.HandlerFunc)}
		
		for _, element := range tt.inputHandlers {
			for route, function := range element {
				handler.AddHandler(route, function)
				routesRegistered = append(routesRegistered, route)
			}
		}

		for _, routeToTest := range routesRegistered {
			handlerToTest = handler.GetHandler(routeToTest)
			req = httptest.NewRequest(http.MethodGet, routeToTest, nil)
			w = httptest.NewRecorder()
			
			handlerToTest(w, req)

			response = w.Result()
			defer response.Body.Close()

			if w.Body.String() != tt.expectedAnswer[routeToTest] && !tt.expectFail {
				t.Errorf("Handler error, expected %q got %q", tt.expectedAnswer[routeToTest], w.Body.String())
			}
		}
	}
}
