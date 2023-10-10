/*
Copyright © 2021 ax-i-om <addressaxiom@pm.me>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/TwiN/go-color"
	"github.com/ax-i-om/bitcrook/cmd"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("   ___  __________________  ____  ____  __ __")
	fmt.Println("  / _ )/  _/_  __/ ___/ _ \\/ __ \\/ __ \\/ //_/")
	fmt.Println(" / _  |/ /  / / / /__/ , _/ /_/ / /_/ / ,<   ")
	fmt.Println("/____/___/ /_/  \\___/_/|_|\\____/\\____/_/|_|  " + color.Colorize(color.Green, "v1.0.0"))
	fmt.Println()
	fmt.Println("Bitcrook Copyright (C) 2021 Axiom\nThis program comes with ABSOLUTELY NO WARRANTY")
	fmt.Println()
	time.Sleep(2500 * time.Millisecond)
	cmd.Execute()
}
