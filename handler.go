package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/ip2location/ip2proxy-go"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// addUser add users
func addUser(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	secret := req.URL.Query().Get("id")
	owner := "dc10ce85-5380-4ca1-8c0a-960eab4d5f22"
	if secret != owner {
		loadResp := Error{false, "no you (눈‸눈)"}
		resp, _ := json.Marshal(loadResp)
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write(resp)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	uid := uuid.New().String()
	newuser := Record{UUID: uid}
	db.Create(&newuser)
	log.Println(uid)
	loadResp := AddUser{true, uid}
	resp, _ := json.Marshal(loadResp)
	_, err := w.Write(resp)
	if err != nil {
		log.Fatal(err)
	}
}

func query(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	ip := req.URL.Query().Get("ip")
	id := req.URL.Query().Get("id")
	if ip == "" || id == "" {
		loadResp := Error{false, "missing parameter"}
		resp, _ := json.Marshal(loadResp)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write(resp)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// auth
	var record Record
	result := db.Find(&record, "uuid = ?", id)
	if result.RowsAffected == 0 {
		loadResp := Error{false, "who u (눈‸눈)"}
		resp, _ := json.Marshal(loadResp)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(resp)
		return
	}

	ipdb, err := ip2proxy.OpenDB("./IP2PROXY-IP-PROXYTYPE-COUNTRY.BIN")
	if err != nil {
		log.Fatal(err)
		return
	}
	all, err := ipdb.GetAll(ip)
	if err != nil {
		fmt.Print(err)
		return
	}
	isProxy, _ := strconv.ParseBool(all["isProxy"])
	queryResp := QueryResult{
		true,
		isProxy,
		all["ProxyType"],
		all["CountryShort"],
		ipdb.DatabaseVersion(),

	}
	resp, _ := json.Marshal(queryResp)
	w.Write(resp)
}