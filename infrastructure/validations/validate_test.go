package validations

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileCountriesJsonExits(t *testing.T) {
	if _, err := os.Stat("../../database/countries.json"); os.IsNotExist(err) {
		assert.Error(t,err)
	}else{
		assert.True(t,true)
	}
}

func TestCountriesJsonContainData(t *testing.T) {
	plan, _:= ioutil.ReadFile("../../database/countries.json")
	var countries []interface{}
	err := json.Unmarshal(plan, &countries)
	if err !=nil {
		assert.Error(t,err)
	}
	assert.NotEmpty(t,countries)
}

func TestGetValidCountryRegexByPhoneNumber(t *testing.T) {
	os.Chdir("../../")
	regex,err:=GetCountryRegexByPhone("(237) 697151594")
	if err !=nil {
		assert.Error(t,err)
	}
	assert.Equal(t,"\\(237\\)\\ ?[2368]\\d{7,8}$",*regex)
	os.Chdir("..")
}

func TestValidatePhoneValid(t *testing.T) {
	os.Chdir("../../")
	status:=ValidatePhone("(237) 697151594")
	assert.Equal(t,true,status)
	os.Chdir("..")
}

func TestValidatePhoneInValid(t *testing.T) {
	os.Chdir("../../")
	status:=ValidatePhone("(237) 697151594sss")
	assert.Equal(t,false,status)
	os.Chdir("..")
}