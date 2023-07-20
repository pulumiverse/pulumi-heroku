// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package team

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-heroku/sdk/go/heroku/internal"
)

var _ = internal.GetEnvOrDefault

type GetMembersMember struct {
	Email                   string `pulumi:"email"`
	Federated               bool   `pulumi:"federated"`
	Role                    string `pulumi:"role"`
	TeamMemberId            string `pulumi:"teamMemberId"`
	TwoFactorAuthentication bool   `pulumi:"twoFactorAuthentication"`
	UserId                  string `pulumi:"userId"`
}

// GetMembersMemberInput is an input type that accepts GetMembersMemberArgs and GetMembersMemberOutput values.
// You can construct a concrete instance of `GetMembersMemberInput` via:
//
//	GetMembersMemberArgs{...}
type GetMembersMemberInput interface {
	pulumi.Input

	ToGetMembersMemberOutput() GetMembersMemberOutput
	ToGetMembersMemberOutputWithContext(context.Context) GetMembersMemberOutput
}

type GetMembersMemberArgs struct {
	Email                   pulumi.StringInput `pulumi:"email"`
	Federated               pulumi.BoolInput   `pulumi:"federated"`
	Role                    pulumi.StringInput `pulumi:"role"`
	TeamMemberId            pulumi.StringInput `pulumi:"teamMemberId"`
	TwoFactorAuthentication pulumi.BoolInput   `pulumi:"twoFactorAuthentication"`
	UserId                  pulumi.StringInput `pulumi:"userId"`
}

func (GetMembersMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMembersMember)(nil)).Elem()
}

func (i GetMembersMemberArgs) ToGetMembersMemberOutput() GetMembersMemberOutput {
	return i.ToGetMembersMemberOutputWithContext(context.Background())
}

func (i GetMembersMemberArgs) ToGetMembersMemberOutputWithContext(ctx context.Context) GetMembersMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetMembersMemberOutput)
}

// GetMembersMemberArrayInput is an input type that accepts GetMembersMemberArray and GetMembersMemberArrayOutput values.
// You can construct a concrete instance of `GetMembersMemberArrayInput` via:
//
//	GetMembersMemberArray{ GetMembersMemberArgs{...} }
type GetMembersMemberArrayInput interface {
	pulumi.Input

	ToGetMembersMemberArrayOutput() GetMembersMemberArrayOutput
	ToGetMembersMemberArrayOutputWithContext(context.Context) GetMembersMemberArrayOutput
}

type GetMembersMemberArray []GetMembersMemberInput

func (GetMembersMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetMembersMember)(nil)).Elem()
}

func (i GetMembersMemberArray) ToGetMembersMemberArrayOutput() GetMembersMemberArrayOutput {
	return i.ToGetMembersMemberArrayOutputWithContext(context.Background())
}

func (i GetMembersMemberArray) ToGetMembersMemberArrayOutputWithContext(ctx context.Context) GetMembersMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetMembersMemberArrayOutput)
}

type GetMembersMemberOutput struct{ *pulumi.OutputState }

func (GetMembersMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMembersMember)(nil)).Elem()
}

func (o GetMembersMemberOutput) ToGetMembersMemberOutput() GetMembersMemberOutput {
	return o
}

func (o GetMembersMemberOutput) ToGetMembersMemberOutputWithContext(ctx context.Context) GetMembersMemberOutput {
	return o
}

func (o GetMembersMemberOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetMembersMember) string { return v.Email }).(pulumi.StringOutput)
}

func (o GetMembersMemberOutput) Federated() pulumi.BoolOutput {
	return o.ApplyT(func(v GetMembersMember) bool { return v.Federated }).(pulumi.BoolOutput)
}

func (o GetMembersMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v GetMembersMember) string { return v.Role }).(pulumi.StringOutput)
}

func (o GetMembersMemberOutput) TeamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMembersMember) string { return v.TeamMemberId }).(pulumi.StringOutput)
}

func (o GetMembersMemberOutput) TwoFactorAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v GetMembersMember) bool { return v.TwoFactorAuthentication }).(pulumi.BoolOutput)
}

func (o GetMembersMemberOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMembersMember) string { return v.UserId }).(pulumi.StringOutput)
}

type GetMembersMemberArrayOutput struct{ *pulumi.OutputState }

func (GetMembersMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetMembersMember)(nil)).Elem()
}

func (o GetMembersMemberArrayOutput) ToGetMembersMemberArrayOutput() GetMembersMemberArrayOutput {
	return o
}

func (o GetMembersMemberArrayOutput) ToGetMembersMemberArrayOutputWithContext(ctx context.Context) GetMembersMemberArrayOutput {
	return o
}

func (o GetMembersMemberArrayOutput) Index(i pulumi.IntInput) GetMembersMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetMembersMember {
		return vs[0].([]GetMembersMember)[vs[1].(int)]
	}).(GetMembersMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetMembersMemberInput)(nil)).Elem(), GetMembersMemberArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetMembersMemberArrayInput)(nil)).Elem(), GetMembersMemberArray{})
	pulumi.RegisterOutputType(GetMembersMemberOutput{})
	pulumi.RegisterOutputType(GetMembersMemberArrayOutput{})
}
