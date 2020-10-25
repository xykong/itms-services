
.PHONY : itms-services

itms-services :
	@go-bindata -o ./generator/bindata.go -pkg generator template
	@go build
