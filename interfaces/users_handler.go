package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/damasukma/sandbox_auth/application"
	"github.com/damasukma/sandbox_auth/domain/entity"
	"github.com/damasukma/sandbox_auth/utils"
	"github.com/gorilla/mux"
)

type Users struct {
	us application.UserAppInterface
}

func NewUsers(us application.UserAppInterface) *Users {
	return &Users{us: us}
}

func (u *Users) CreateUser(w http.ResponseWriter, r *http.Request) {

	validations := make(map[string]string)

	users := &entity.User{}
	response := &utils.Response{Status: http.StatusOK, Message: "Success"}
	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "Bad Request"
		response.ToJson(w)
		return
	}

	if email := u.us.EmailExist(users.Email); email {
		validations["email"] = "Already Exist"
		response.Status = http.StatusUnprocessableEntity
		response.Message = "Unprocessable Entity"
		response.Validation = &validations
		response.ToJson(w)
		return
	}

	if err := u.us.SaveUser(users); err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = "Internal Server Error"
	}

	response.ToJson(w)
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {

	users := &entity.User{}

	response := &utils.Response{Status: http.StatusOK, Message: "Success"}
	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "Bad Request"
		response.ToJson(w)
		return
	}

	result, err := u.us.Auth(users.Email, users.Password)

	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		response.ToJson(w)
		return
	}

	if !result {
		response.Status = http.StatusUnauthorized
		response.Message = "Unauthorized"
	}
	response.ToJson(w)

}

func (u *Users) Find(w http.ResponseWriter, r *http.Request) {
	var data interface{}

	response := &utils.Response{Status: http.StatusOK, Message: "Success"}

	results, err := u.us.Find()
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		response.ToJson(w)
		return
	}
	data = results
	response.Data = &data

	response.ToJson(w)

}

func (u *Users) FindID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var data interface{}
	response := &utils.Response{Status: http.StatusOK, Message: "Success"}

	id, _ := strconv.Atoi(params["id"])

	result, err := u.us.FindID(id)

	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		response.ToJson(w)
		return
	}

	data = result
	response.Data = &data
	response.ToJson(w)

}

func (u *Users) Update(w http.ResponseWriter, r *http.Request) {
	users := &entity.User{}

	response := &utils.Response{Status: http.StatusOK, Message: "Success"}

	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "Bad Request"
		response.ToJson(w)
		return
	}

	if err := u.us.Update(*users); err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		response.ToJson(w)
		return
	}

	response.ToJson(w)

}

func (u *Users) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	response := &utils.Response{Status: http.StatusOK, Message: "Success"}

	if err := u.us.Delete(id); err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		response.ToJson(w)
		return
	}
	response.ToJson(w)
}
