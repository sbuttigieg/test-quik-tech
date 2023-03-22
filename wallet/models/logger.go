package models

import "strconv"

type LogReqAPI struct {
	Method     string
	Path       string
	StatusCode int
}

func (l LogReqAPI) String() string {
	return "Method: " + l.Method + ", Path: " + l.Path + ", StatusCode: " + strconv.Itoa(l.StatusCode)
}

type LogService struct {
	Layer       string
	Duration    string
	ServiceName string
	Method      string
	Data        interface{}
}

func (l LogService) String() string {
	return "Layer: " + l.Layer + ", Duration: " + l.Duration + ", Service-name: " + l.ServiceName + ", Method: " + l.Method + ", Data: " + l.Data.(string)
}
