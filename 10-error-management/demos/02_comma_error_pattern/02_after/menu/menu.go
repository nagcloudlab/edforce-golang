package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

var in = bufio.NewReader(os.Stdin)

func Print() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
	}
}
func Add() error {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)

	for _, item := range menu {
		if item.name == name {
			return errors.New("menu item already exists")
		}
	}
	menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
	return nil
}
