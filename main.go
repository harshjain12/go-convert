// Copyright 2022 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/drone/go-convert/command"
	"github.com/google/subcommands"
	"net/http"
	"os/exec"
)

func runBinaryHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters for two string inputs
	query := r.URL.Query()
	param1 := query.Get("input1")
	param2 := query.Get("input2")

	if param1 == "" || param2 == "" {
		http.Error(w, "Both input1 and input2 are required", http.StatusBadRequest)
		return
	}

	flag.Parse()
	ctx := context.Background()

	// Here, replace with the actual path to your binary executable
	// For this example, let's assume we have a simple binary that takes two strings as arguments
	// cmd := exec.Command("./go-convert")
	os.Exit(int(subcommands.Execute(ctx, param1, param2)))

	// Run the command and capture the output
	//output, err := cmd.CombinedOutput()
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Error running binary: %v", err), http.StatusInternalServerError)
	//	return
	//}
	//
	//// Send the output back to the client
	//w.WriteHeader(http.StatusOK)
	//w.Write(output)
}

func main() {
	subcommands.Register(new(command.Azure), "")
	subcommands.Register(new(command.Bitbucket), "")
	subcommands.Register(new(command.Circle), "")
	subcommands.Register(new(command.Cloudbuild), "")
	subcommands.Register(new(command.Drone), "")
	subcommands.Register(new(command.Github), "")
	subcommands.Register(new(command.Gitlab), "")
	subcommands.Register(new(command.Jenkins), "")
	subcommands.Register(new(command.Travis), "")
	subcommands.Register(new(command.Downgrade), "")
	subcommands.Register(new(command.JenkinsJson), "")
	subcommands.Register(new(command.JenkinsXml), "")

	http.HandleFunc("/run", runBinaryHandler)

	// Start the server
	fmt.Println("Server starting on :8990...")
	if err := http.ListenAndServe(":8990", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
