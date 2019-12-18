/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"
	"os"

	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/vulk/go-travis"
	// "github.com/koshatul/go-travis"
	// "github.com/shuheiktgw/go-travis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

// "https://travis-ci.org/crosscloudci/testproj/builds/572521581"
type CliResponse struct {
	JobUrl          string
	BuildUrl        string
	BuildStatus     string
	OptionalMessage string
}

func (c *CliResponse) output() (output string) {
	//TODO if -q parameter don't add header
	fmt.Printf("status\tbuild_url\n")
	fmt.Printf("%v\t%v \n", c.BuildStatus, c.BuildUrl)
	// fmt.Printf("{'build_url': '%v', 'status': '%v'}", c.BuildUrl, c.BuildStatus)
	return
}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:              "status",
	TraverseChildren: true,
	Short:            "This command retrieves the status of a travis-ci project build",
	Long:             `This command takes a project name, commit ref, or tag and return success, failure, or running.`,
	Run: func(cmd *cobra.Command, args []string) {
		// client := travis.NewClient(travis.ApiOrgUrl, os.Getenv("TRAVIS_API_KEY"))
		client := travis.NewClient(travis.ApiComUrl, os.Getenv("TRAVIS_API_KEY"))
		opt := &travis.BuildsByRepoOption{Limit: 100, Include: []string{"build.commit"}}
		// opt := &travis.BuildsByRepoOption{Limit: 100, SortBy: "finished_at:desc", Include: []string{"build.commit"}}
		// opt := &travis.BuildsByRepoOption{Limit: 100, SortBy: "id:asc", Include: []string{"build.commit"}}
		var returned_build_status string
		var returned_build_url string
		var cli_response CliResponse

		// Generics anyone?
		var reverse = func(lst []*travis.Build) chan *travis.Build {
			ret := make(chan *travis.Build)
			go func() {
				for i, _ := range lst {
					ret <- lst[len(lst)-1-i]
				}
				close(ret)
			}()
			return ret
		}

		var retrieveMatchingBuildStatus = func() {
			var done bool
			done = false
			for {
				build, resp, err := client.Builds.ListByRepoSlug(context.Background(), viper.GetString("project"), opt)
				// spew.Dump("err %v", err)
				// spew.Dump("build %v", build)
				if err != nil {
					panic(err)
				}
				for b := range reverse(build) {
					// spew.Dump("build by rep commit.sha %v", *b.Commit.Sha)
					// spew.Dump("build by rep build.number %v", *b.Number)
					// spew.Dump("build by rep build id %v", *b.Id)
					// spew.Dump("build by rep build.commit.id %v", *b.Commit.Id)
					arg_commit := viper.GetString("commit")
					if (*b.Commit.Sha)[:6] == arg_commit[:6] {
						returned_build_status = *b.State
						returned_build_url = *b.Href
						if viper.GetBool("verbose") {
							spew.Dump("travis build", b)
						}
						done = true
					}
				}
				if resp.NextPage == nil {
					done = true
				}
				if done == true {
					break
				}
				opt.Limit = resp.NextPage.Limit
				opt.Offset = resp.NextPage.Offset
			}

			switch returned_build_status {
			case "received":
				returned_build_status = "running"
			case "created":
				returned_build_status = "running"
			case "started":
				returned_build_status = "running"
			case "passed":
				returned_build_status = "success"
			case "errored":
				returned_build_status = "failed"
			case "failed":
				returned_build_status = "failed"
			case "canceled":
				returned_build_status = "failed"
			default:
				os.Stdout.Sync()
				fmt.Fprintf(os.Stderr, "ERROR: %v \n", "failed to find project with given commit")
				os.Exit(1)
			}

			url_prefix := fmt.Sprintf("https://travis-ci.com/%s/builds", viper.GetString("project"))
			cli_response.BuildUrl = strings.Replace(returned_build_url, "/build", url_prefix, 1)
			cli_response.BuildStatus = returned_build_status

			fmt.Printf(cli_response.output())
		}
		retrieveMatchingBuildStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.PersistentFlags().StringP("project", "p", "", "travis-ci project name")
	statusCmd.PersistentFlags().StringP("commit", "c", "", "travis-ci project commit sha")
	statusCmd.PersistentFlags().StringP("tag", "t", "", "travis-ci project tag")
	statusCmd.PersistentFlags().BoolP("verbose", "v", false, "travis-ci verbose output")
	viper.BindPFlag("project", statusCmd.PersistentFlags().Lookup("project"))
	viper.BindPFlag("commit", statusCmd.PersistentFlags().Lookup("commit"))
	viper.BindPFlag("tag", statusCmd.PersistentFlags().Lookup("tag"))
	viper.BindPFlag("verbose", statusCmd.PersistentFlags().Lookup("verbose"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
