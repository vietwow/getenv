package model

type Request struct {
    AppName string `json:"app_name" form:"app_name" query:"app_name"`
}

