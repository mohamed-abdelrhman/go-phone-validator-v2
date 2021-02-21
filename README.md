To Start the GoLang APP

## open the project Directory and run
docker build . --tag golang:phoneValidator
then
docker run -it -p 8000:8000 golang:phoneValidator

##entry point
localhost:8000/customers

##filter with country code
localhost:8000/customers?country_code=(237)


##filter with status
localhost:8000/customers?status=valid
localhost:8000/customers?status=invalid

##filter with countryCode and status
localhost:8000/customers?country_code=(237)&status=valid