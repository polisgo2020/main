package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// http://127.0.0.1:8080/?priority=low&subject=Hello!&inner=ignored&id=12&recipient=test@host.ru

type SendMessage struct {
	Id        int    `validate:"gte=0,lte=100"`
	Priority  string `validate:"required,oneof=low normal high"`
	Recipient string `validate:"required,email"`
	Subject   string `validate:"msgSubject"`
	Inner     string
	flag      int
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("request " + r.URL.String() + "\n\n"))

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	msg := &SendMessage{
		Id:        id,
		Priority:  r.URL.Query().Get("priority"),
		Recipient: r.URL.Query().Get("recipient"),
		Subject:   r.URL.Query().Get("subject"),
		Inner:     r.URL.Query().Get("inner"),
	}

	w.Write([]byte(fmt.Sprintf("Msg: %#v\n\n", msg)))

	err := validate.Struct(msg)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			data := []byte(fmt.Sprintf("field: %#v\n\n", err.Field()))
			w.Write(data)
		}

		w.Write([]byte(fmt.Sprintf("error: %s\n\n", err)))
	} else {
		w.Write([]byte(fmt.Sprintf("msg is correct\n\n")))
	}
}

func init() {
	validate = validator.New()
	validate.RegisterValidation("msgSubject", msgSubjectValidator)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

func msgSubjectValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if len(value) == 0 || len(value) > 10 {
		return false
	}
	return true
}
