package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
	//"github.com/ghodss/yaml"

	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type conf struct {
	Action         string `yaml:"action"`
	Name           string `yaml:"eks_cluster_name"`
	Region         string `yaml:"region"`
	Vpc_name       string `yaml:"vpc_name"`
	Cidr           string `yaml:"cidr"`
	NodeGroup      string `yaml:"NodeGroup"`
	Public_Access  bool   `yaml:"public_access"`
	Private_Access bool   `yaml:"private_access"`
	Zones []string     `json:"zones"`
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func main() {
	var c conf
	c.getConf()
	if c.Action == "" {
		fmt.Println("Cluster config: invalid `Action`, should be create, update, or delete")
		os.Exit(1)
	} else if c.Action == "create" {
		fmt.Println("Cluster Creation In Progress")
	} else if c.Action == "update" {
		fmt.Println("Cluster update In Progress")
	} else if c.Action == "delete" {
		fmt.Println("Cluster Deletion In Progress")
	} else {

	}

	tpl, err := template.ParseFiles("templates/cluster.yaml.tmpl")

	if err != nil {
		fmt.Print("Template Not Found")
	}

	f, err := os.Create("cluster.yaml")

	if err != nil {
		log.Println("create file: ", err)
		return
	}
	err = tpl.Execute(f, c)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()
}
