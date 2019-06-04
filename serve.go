package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	//"fmt"
	//"io/ioutil"
	"log"
	"net/http"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/index.html")
	err := t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func handlerGenerate(w http.ResponseWriter, r *http.Request) {
	s := generate()
	resp := Body{
		Code: 1000,
		Msg:  "请求成功",
		Data: s.board,
	}
	json.NewEncoder(w).Encode(&resp)
}

func handlerAdjust(w http.ResponseWriter, r *http.Request) {
	var req, resp Body
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	json.NewDecoder(r.Body).Decode(&req)
	var s State
	// interface 转 slice
	for _, num := range req.Data.([]interface{}) {
		s.board = append(s.board, int(num.(float64)))
	}
	s.block = s.Block()
	if s.SolveAble() {
		resp.Code = 1004
		resp.Data = "已经可解,无需重新求解"
		json.NewEncoder(w).Encode(&resp)
		return
	}
	s.Adjust()
	resp.Code = 1000
	resp.Msg = "调整成功"
	resp.Data = s.board
	json.NewEncoder(w).Encode(&resp)
}

func handlerSolution(w http.ResponseWriter, r *http.Request) {
	var req, resp Body
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	json.NewDecoder(r.Body).Decode(&req)
	var s State
	// interface 转 slice
	for _, num := range req.Data.([]interface{}) {
		s.board = append(s.board, int(num.(float64)))
	}
	s.block = s.Block()
	//s.Show()
	if s.Solved() {
		resp.Code = 1001
		resp.Msg = "该状态已解"
		json.NewEncoder(w).Encode(&resp)
		return
	}
	if !s.SolveAble() {
		resp.Code = 1002
		resp.Msg = "该状态不可解答"
		json.NewEncoder(w).Encode(&resp)
		return
	}
	solution, _ := s.Solution()
	if solution == "" {
		resp.Code = 1003
		resp.Msg = "求解失败"
		json.NewEncoder(w).Encode(&resp)
		return
	}
	resp.Code = 1000
	resp.Msg = "求解成功"
	resp.Data = solution
	json.NewEncoder(w).Encode(&resp)

}
func server() {
	// 启动静态文件服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/generate", handlerGenerate)
	http.HandleFunc("/solution", handlerSolution)
	http.HandleFunc("/adjust", handlerAdjust)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
