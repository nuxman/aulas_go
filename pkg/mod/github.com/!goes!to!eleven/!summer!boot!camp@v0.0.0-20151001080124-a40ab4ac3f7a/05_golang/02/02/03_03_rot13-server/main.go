package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			line := strings.ToLower(scanner.Text())
			bs := []byte(line)
//			fmt.Fprintf(conn, "%T\n", bs)
//			fmt.Fprintf(conn, "%s\n", bs)
//			fmt.Fprintf(conn, "%v\n", bs)
//			fmt.Fprintf(conn, "%q\n", bs)
			var rot13 = make([]byte, len(bs))
			for k, v := range bs {
//				fmt.Fprintf(conn, "%v", k)
//				fmt.Fprintf(conn, "- %v\n", v)
				// ascii 97 - 122
				/// 109 + 13 = 122
				if v <= 109 {
					rot13[k] = v+13
				} else {
					// 110 + 13 = 123
					//123 - 26 = 97
					rot13[k] = v-26+13
				}
			}
//			fmt.Fprintf(conn, "rot13: %v\n", rot13)
			fmt.Fprintf(conn, "rot13: %s\n", rot13)
			}

	}
}