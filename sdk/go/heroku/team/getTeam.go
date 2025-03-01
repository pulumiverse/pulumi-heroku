// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package team

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-heroku/sdk/go/heroku/internal"
)

// Use this data source to get information about a Heroku Team.
//
// ## Example Usage
func GetTeam(ctx *pulumi.Context, args *GetTeamArgs, opts ...pulumi.InvokeOption) (*GetTeamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTeamResult
	err := ctx.Invoke("heroku:team/getTeam:getTeam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTeam.
type GetTeamArgs struct {
	// The team name
	Name string `pulumi:"name"`
}

// A collection of values returned by getTeam.
type GetTeamResult struct {
	// Whether to use this team when none is specified
	Default bool `pulumi:"default"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Upper limit of members allowed in a team
	MembershipLimit int    `pulumi:"membershipLimit"`
	Name            string `pulumi:"name"`
	// Whether the team is provisioned licenses by Salesforce
	ProvisionedLicenses bool `pulumi:"provisionedLicenses"`
	// type of team Will likely be either "enterprise" or "team"
	Type string `pulumi:"type"`
}

func GetTeamOutput(ctx *pulumi.Context, args GetTeamOutputArgs, opts ...pulumi.InvokeOption) GetTeamResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetTeamResultOutput, error) {
			args := v.(GetTeamArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("heroku:team/getTeam:getTeam", args, GetTeamResultOutput{}, options).(GetTeamResultOutput), nil
		}).(GetTeamResultOutput)
}

// A collection of arguments for invoking getTeam.
type GetTeamOutputArgs struct {
	// The team name
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetTeamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTeamArgs)(nil)).Elem()
}

// A collection of values returned by getTeam.
type GetTeamResultOutput struct{ *pulumi.OutputState }

func (GetTeamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTeamResult)(nil)).Elem()
}

func (o GetTeamResultOutput) ToGetTeamResultOutput() GetTeamResultOutput {
	return o
}

func (o GetTeamResultOutput) ToGetTeamResultOutputWithContext(ctx context.Context) GetTeamResultOutput {
	return o
}

// Whether to use this team when none is specified
func (o GetTeamResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v GetTeamResult) bool { return v.Default }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTeamResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTeamResult) string { return v.Id }).(pulumi.StringOutput)
}

// Upper limit of members allowed in a team
func (o GetTeamResultOutput) MembershipLimit() pulumi.IntOutput {
	return o.ApplyT(func(v GetTeamResult) int { return v.MembershipLimit }).(pulumi.IntOutput)
}

func (o GetTeamResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetTeamResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether the team is provisioned licenses by Salesforce
func (o GetTeamResultOutput) ProvisionedLicenses() pulumi.BoolOutput {
	return o.ApplyT(func(v GetTeamResult) bool { return v.ProvisionedLicenses }).(pulumi.BoolOutput)
}

// type of team Will likely be either "enterprise" or "team"
func (o GetTeamResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetTeamResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTeamResultOutput{})
}
