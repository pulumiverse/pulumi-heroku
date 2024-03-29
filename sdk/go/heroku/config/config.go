// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/pulumiverse/pulumi-heroku/sdk/go/heroku/internal"
)

var _ = internal.GetEnvOrDefault

func GetApiKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "heroku:apiKey")
}
func GetCustomizations(ctx *pulumi.Context) string {
	return config.Get(ctx, "heroku:customizations")
}
func GetDelays(ctx *pulumi.Context) string {
	return config.Get(ctx, "heroku:delays")
}
func GetEmail(ctx *pulumi.Context) string {
	return config.Get(ctx, "heroku:email")
}
func GetHeaders(ctx *pulumi.Context) string {
	return config.Get(ctx, "heroku:headers")
}
func GetTimeouts(ctx *pulumi.Context) string {
	return config.Get(ctx, "heroku:timeouts")
}
func GetUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "heroku:url")
}
