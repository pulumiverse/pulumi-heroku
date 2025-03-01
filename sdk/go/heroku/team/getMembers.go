// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package team

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-heroku/sdk/go/heroku/internal"
)

// Use this data source to get information about members for a Heroku Team.
//
// ## Example Usage
func GetMembers(ctx *pulumi.Context, args *GetMembersArgs, opts ...pulumi.InvokeOption) (*GetMembersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMembersResult
	err := ctx.Invoke("heroku:team/getMembers:getMembers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMembers.
type GetMembersArgs struct {
	// List of roles. Acceptable values are `admin`, `member`, `viewer`, `collaborator`, `owner`.
	// At least one role must be specified.
	Roles []string `pulumi:"roles"`
	// The team name.
	Team string `pulumi:"team"`
}

// A collection of values returned by getMembers.
type GetMembersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// All members of the team that have a specified role defined in the `roles` attribute above.
	Members []GetMembersMember `pulumi:"members"`
	Roles   []string           `pulumi:"roles"`
	Team    string             `pulumi:"team"`
}

func GetMembersOutput(ctx *pulumi.Context, args GetMembersOutputArgs, opts ...pulumi.InvokeOption) GetMembersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetMembersResultOutput, error) {
			args := v.(GetMembersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("heroku:team/getMembers:getMembers", args, GetMembersResultOutput{}, options).(GetMembersResultOutput), nil
		}).(GetMembersResultOutput)
}

// A collection of arguments for invoking getMembers.
type GetMembersOutputArgs struct {
	// List of roles. Acceptable values are `admin`, `member`, `viewer`, `collaborator`, `owner`.
	// At least one role must be specified.
	Roles pulumi.StringArrayInput `pulumi:"roles"`
	// The team name.
	Team pulumi.StringInput `pulumi:"team"`
}

func (GetMembersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMembersArgs)(nil)).Elem()
}

// A collection of values returned by getMembers.
type GetMembersResultOutput struct{ *pulumi.OutputState }

func (GetMembersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMembersResult)(nil)).Elem()
}

func (o GetMembersResultOutput) ToGetMembersResultOutput() GetMembersResultOutput {
	return o
}

func (o GetMembersResultOutput) ToGetMembersResultOutputWithContext(ctx context.Context) GetMembersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetMembersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMembersResult) string { return v.Id }).(pulumi.StringOutput)
}

// All members of the team that have a specified role defined in the `roles` attribute above.
func (o GetMembersResultOutput) Members() GetMembersMemberArrayOutput {
	return o.ApplyT(func(v GetMembersResult) []GetMembersMember { return v.Members }).(GetMembersMemberArrayOutput)
}

func (o GetMembersResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMembersResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o GetMembersResultOutput) Team() pulumi.StringOutput {
	return o.ApplyT(func(v GetMembersResult) string { return v.Team }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMembersResultOutput{})
}
