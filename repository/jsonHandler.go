package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mercadolibre/model"
	"os"
)

type callerInfoType model.CallerInfo

func NewJSONRepository() *callerInfoType {
	return &callerInfoType{}
}

type JSONInterface interface {
	ReadJSONFile()
	WriteJSONFile()
	FailureId() []int
	FailureCnt() int
	ErrorCnt() int
	LastRequstID() int
	SetFailureId(requestID int)
	SetFailureCnt(failurecnt int)
	SetErrorCnt(errorcnt int)
	SetLastRequstID(LastRequestId int)
}

//read the JSON file
func (callerinfo *callerInfoType) ReadJSONFile() {
	if _, err := os.Stat("itemtrack.json"); os.IsNotExist(err) {
		fmt.Println("file not present")
		return
	}
	JSONFiledata, err := ioutil.ReadFile("itemtrack.json")
	if err != nil {
		fmt.Println("file can't be read", err)
		return
	}
	err = json.Unmarshal(JSONFiledata, &callerinfo)
	if err != nil {
		fmt.Println("invalid json ", err)
		return
	}
	fmt.Println("Assinged last request Id")
	fmt.Print("Last call info", string(JSONFiledata))
}

//write the JSON file
func (callerinfo *callerInfoType) WriteJSONFile() {
	file, err := json.MarshalIndent(callerinfo, "", " ")
	if err != nil {
		fmt.Println("invalid data json not created ", err)
		return
	}
	err = ioutil.WriteFile("itemtrack.json", file, 0644)
	if err != nil {
		fmt.Println("file can't be write ", err)
		return
	}
}

//Get failureIds
func (callerinfo *callerInfoType) FailureId() []int {
	return callerinfo.FailureIds
}

//Set failureIds
func (callerinfo *callerInfoType) SetFailureId(requestID int) {
	callerinfo.FailureIds = append(callerinfo.FailureIds, requestID)
}

//Get failureCnt
func (callerinfo *callerInfoType) FailureCnt() int {
	return callerinfo.FailureCount
}

//Set failureCnt
func (callerinfo *callerInfoType) SetFailureCnt(failurecnt int) {
	callerinfo.FailureCount = failurecnt
}

//Get ErrorCnt
func (callerinfo *callerInfoType) ErrorCnt() int {
	return callerinfo.ErrorCount
}

//Set ErrorCnt
func (callerinfo *callerInfoType) SetErrorCnt(errorcnt int) {
	callerinfo.ErrorCount = errorcnt
}

//Get LastRequestId
func (callerinfo *callerInfoType) LastRequstID() int {
	return callerinfo.LastRequestId
}

//Set LastRequestId
func (callerinfo *callerInfoType) SetLastRequstID(LastRequestId int) {
	callerinfo.LastRequestId = LastRequestId
}
