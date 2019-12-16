package cmd_test

import (
	// "ci_plugin_travis_go/cmd"
	"github.com/stretchr/testify/assert"
	"testing"
)

//  (*travis.Build)(0xc000512780)({
//   Id: (*uint)(0xc0002b8f50)(569902352),
//   Number: (*string)(0xc0002a9a60)((len=1) "2"),
//   State: (*string)(0xc0002a9a70)((len=7) "errored"),
//   Duration: (*uint)(0xc0002b8f70)(52),
//   EventType: (*string)(0xc0002a9a80)((len=4) "push"),
//   PreviousState: (*string)(0xc0002a9a90)((len=6) "failed"),
//   PullRequestTitle: (*string)(<nil>),
//   PullRequestNumber: (*uint)(<nil>),
//   StartedAt: (*string)(0xc0002a9aa0)((len=20) "2019-08-09T15:51:13Z"),
//   FinishedAt: (*string)(0xc0002a9ab0)((len=20) "2019-08-09T15:52:05Z"),
//   UpdatedAt: (*string)(0xc0002a9c90)((len=24) "2019-08-09T15:52:05.476Z"),
//   Private: (*bool)(0xc0002b8f86)(false),
//   Repository: (*travis.Repository)(0xc000102880)({
//    Id: (*uint)(0xc0002b8fb0)(25641832),
//    Name: (*string)(0xc0002a9af0)((len=8) "testproj"),
//    Slug: (*string)(0xc0002a9b00)((len=21) "crosscloudci/testproj"),
//    Description: (*string)(<nil>),
//    GitHubId: (*uint)(<nil>),
//    GitHubLanguage: (*string)(<nil>),
//    Active: (*bool)(<nil>),
//    Private: (*bool)(<nil>),
//    Owner: (*travis.Owner)(<nil>),
//    DefaultBranch: (*travis.Branch)(<nil>),
//    Starred: (*bool)(<nil>),
//    ManagedByInstallation: (*bool)(<nil>),
//    ActiveOnOrg: (*bool)(<nil>),
//    MigrationStatus: (*string)(<nil>),
//    AllowMigration: (*bool)(<nil>),
//    Metadata: (*travis.Metadata)(0xc0002917a0)({
//     Type: (*string)(0xc0002a9ac0)((len=10) "repository"),
//     Href: (*string)(0xc0002a9ad0)((len=14) "/repo/25641832"),
//     Representation: (*string)(0xc0002a9ae0)((len=7) "minimal"),
//     Permissions: (*travis.Permissions)(<nil>)
//    })
//   }),

func TestStatusCmd(t *testing.T) {
	assert.True(t, false, "True is true!")
	cmd
	// cmd
	// status
}
