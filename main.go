package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println(`
		usage: myskill <command> [options]
		commands: add, get, list, delete, search
		
		Examples:
		
		`)
}

func main() {

	args := os.Args

	if len(args) < 2 {
		usage()

		os.Exit(1)
	}

	switch args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("title", "", "task title")

		addCmd.Parse(args[2:])

		fmt.Println("Adding:", *title)

	case "get":
		getCmd := flag.NewFlagSet("get", flag.ExitOnError)
		id := getCmd.String("id", "", "item id")

		getCmd.Parse(args[2:])

		fmt.Println("Getting:", *id)

	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		category := listCmd.String("cat", "", "category filter")

		listCmd.Parse(os.Args[2:])
		fmt.Println("Category:", *category)

	case "delete":
		delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := delCmd.String("id", "", "delete item")

		delCmd.Parse(os.Args[2:])
		fmt.Println("id", *id)

	case "search":
		searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
		find := searchCmd.String("find", "", "find text")

		searchCmd.Parse(args[2:])

		fmt.Println("Searching:", *find)

	default:
		fmt.Println(`unknow command`)
		usage()
		os.Exit(1)
	}
}
