// This is a multi-purpose executable for the networking programming needed by docker
// docker-network-servicer SUBCOMMAND { COMMAND }
// SUBCOMMAND := redirecter | marker | resolver
package main

import (
	"flag"
	"log"
	"os"
)

type isDeleteFlag struct {
	set   bool
	value string
}

func (df *isDeleteFlag) Set(x string) error {
	df.value = x
	df.set = true
	return nil
}

func (df *isDeleteFlag) String() string {
	return df.value
}

func main() {
	var (
		path = flag.String("path", "", "netns path")
		vip  = flag.String("vip", "", "virtual ip")
		eip  = flag.String("eip", "", "endpoint ip")
		lip  = flag.String("lip", "", "local address")
		ltcp = flag.String("ltcp", "", "local tcp listen address")
		file = flag.String("ports-file", "", "ingress ports file")
		mark = flag.String("mark", "-1", "firewall mark")
		del  isDeleteFlag
	)

	flag.Var(&del, "del", "delete operation")

	flag.CommandLine.Parse(os.Args[2:])

	switch os.Args[1] {
	case "marker":
		checkNumArgs(5)
		if code, err := marker(*path, *vip, *eip, *mark, *file, del.set); err != nil {
			log.Println(err)
			os.Exit(code)
		}
	case "redirecter":
		checkNumArgs(3)
		if code, err := redirecter(*path, *eip, *file); err != nil {
			log.Println(err)
			os.Exit(code)
		}
	case "resolver":
		checkNumArgs(3)
		if code, err := setupResolver(*path, *lip, *ltcp); err != nil {
			log.Println(err)
			os.Exit(code)
		}
	default:
		log.Panicf("unrecognized sub command: %q. Usage: %s", os.Args[1], usage(os.Args[1]))
	}
}

func checkNumArgs(num int) {
	if flag.NFlag() < num {
		log.Panicf("Invalid number of arguments for %q sub command: %d. Usage: %s",
			os.Args[1], flag.NFlag(), usage(os.Args[1]))
	}
}

func usage(subc string) string {
	switch subc {
	case "redirecter":
		return "docker-network-servicer redirecter -eip <endpoint ip> -ports-file <ingress port file>"
	case "marker":
		return "docker-network-servicer marker -vip <vip> -eip <endpoint ip> -ports-file <ingress port file> -mark <mark> [-del]"
	case "resolver":
		return "docker-network-servicer resolver -lip <local address> -ltcp <local tcp listen address>"

	default:
		return "docker-network-servicer redirecter | marker | resolver { COMMAND }"
	}
}
