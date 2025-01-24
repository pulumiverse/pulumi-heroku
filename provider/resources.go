// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package heroku

import (
	"fmt"
	"path/filepath"

	"github.com/heroku/terraform-provider-heroku/v5/heroku"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumiverse/pulumi-heroku/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "heroku"
	// modules:
	mainMod = "index" // the heroku module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(heroku.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "heroku",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Heroku",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "pulumiverse - Marcel Arns",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/pulumiverse",
		Description:       "A Pulumi package for creating and managing heroku cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "heroku", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/pulumiverse/pulumi-heroku",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg:               "heroku",
		TFProviderModuleVersion: "v5",
		Config:                  map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"heroku_account_feature":                   {Tok: tfbridge.MakeResource(mainPkg, "account", "Feature")},
			"heroku_addon":                             {Tok: tfbridge.MakeResource(mainPkg, "addon", "Addon")},
			"heroku_addon_attachment":                  {Tok: tfbridge.MakeResource(mainPkg, "addon", "Attachment")},
			"heroku_app":                               {Tok: tfbridge.MakeResource(mainPkg, "app", "App")},
			"heroku_app_config_association":            {Tok: tfbridge.MakeResource(mainPkg, "app", "ConfigAssociation")},
			"heroku_config":                            {Tok: tfbridge.MakeResource(mainPkg, "app", "Config")},
			"heroku_app_feature":                       {Tok: tfbridge.MakeResource(mainPkg, "app", "Feature")},
			"heroku_app_release":                       {Tok: tfbridge.MakeResource(mainPkg, "app", "Release")},
			"heroku_app_webhook":                       {Tok: tfbridge.MakeResource(mainPkg, "app", "Webhook")},
			"heroku_build":                             {Tok: tfbridge.MakeResource(mainPkg, "build", "Build")},
			"heroku_cert":                              {Tok: tfbridge.MakeResource(mainPkg, "cert", "Cert")},
			"heroku_collaborator":                      {Tok: tfbridge.MakeResource(mainPkg, "collaborator", "Collaborator")},
			"heroku_domain":                            {Tok: tfbridge.MakeResource(mainPkg, "domain", "Domain")},
			"heroku_drain":                             {Tok: tfbridge.MakeResource(mainPkg, "drain", "Drain")},
			"heroku_formation":                         {Tok: tfbridge.MakeResource(mainPkg, "formation", "Formation")},
			"heroku_pipeline":                          {Tok: tfbridge.MakeResource(mainPkg, "pipeline", "Pipeline")},
			"heroku_pipeline_config_var":               {Tok: tfbridge.MakeResource(mainPkg, "pipeline", "ConfigVar")},
			"heroku_pipeline_coupling":                 {Tok: tfbridge.MakeResource(mainPkg, "pipeline", "Coupling")},
			"heroku_review_app_config":                 {Tok: tfbridge.MakeResource(mainPkg, "review", "AppConfig")},
			"heroku_slug":                              {Tok: tfbridge.MakeResource(mainPkg, "slug", "Slug")},
			"heroku_space":                             {Tok: tfbridge.MakeResource(mainPkg, "space", "Space")},
			"heroku_space_app_access":                  {Tok: tfbridge.MakeResource(mainPkg, "space", "AppAccess")},
			"heroku_space_inbound_ruleset":             {Tok: tfbridge.MakeResource(mainPkg, "space", "InboundRuleset")},
			"heroku_space_peering_connection_accepter": {Tok: tfbridge.MakeResource(mainPkg, "space", "PeeringConnectionAccepter")},
			"heroku_space_vpn_connection":              {Tok: tfbridge.MakeResource(mainPkg, "space", "VpnConnection")},
			"heroku_ssl":                               {Tok: tfbridge.MakeResource(mainPkg, "ssl", "Ssl")},
			"heroku_team_collaborator":                 {Tok: tfbridge.MakeResource(mainPkg, "team", "Collaborator")},
			"heroku_team_member":                       {Tok: tfbridge.MakeResource(mainPkg, "team", "Member")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"heroku_addon":              {Tok: tfbridge.MakeDataSource(mainPkg, "addon", "getAddon")},
			"heroku_app":                {Tok: tfbridge.MakeDataSource(mainPkg, "app", "getApp")},
			"heroku_pipeline":           {Tok: tfbridge.MakeDataSource(mainPkg, "pipeline", "getPipeline")},
			"heroku_space":              {Tok: tfbridge.MakeDataSource(mainPkg, "space", "getSpace")},
			"heroku_space_peering_info": {Tok: tfbridge.MakeDataSource(mainPkg, "space", "getPeeringInfo")},
			"heroku_team":               {Tok: tfbridge.MakeDataSource(mainPkg, "team", "getTeam")},
			"heroku_team_members":       {Tok: tfbridge.MakeDataSource(mainPkg, "team", "getMembers")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/heroku",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumiverse_heroku",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumiverse",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
