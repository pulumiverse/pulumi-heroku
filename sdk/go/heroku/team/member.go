// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package team

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-heroku/sdk/go/heroku/internal"
)

type Member struct {
	pulumi.CustomResourceState

	Email     pulumi.StringOutput  `pulumi:"email"`
	Federated pulumi.BoolPtrOutput `pulumi:"federated"`
	Role      pulumi.StringOutput  `pulumi:"role"`
	Team      pulumi.StringOutput  `pulumi:"team"`
}

// NewMember registers a new resource with the given unique name, arguments, and options.
func NewMember(ctx *pulumi.Context,
	name string, args *MemberArgs, opts ...pulumi.ResourceOption) (*Member, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Team == nil {
		return nil, errors.New("invalid value for required argument 'Team'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Member
	err := ctx.RegisterResource("heroku:team/member:Member", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMember gets an existing Member resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MemberState, opts ...pulumi.ResourceOption) (*Member, error) {
	var resource Member
	err := ctx.ReadResource("heroku:team/member:Member", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Member resources.
type memberState struct {
	Email     *string `pulumi:"email"`
	Federated *bool   `pulumi:"federated"`
	Role      *string `pulumi:"role"`
	Team      *string `pulumi:"team"`
}

type MemberState struct {
	Email     pulumi.StringPtrInput
	Federated pulumi.BoolPtrInput
	Role      pulumi.StringPtrInput
	Team      pulumi.StringPtrInput
}

func (MemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*memberState)(nil)).Elem()
}

type memberArgs struct {
	Email     string `pulumi:"email"`
	Federated *bool  `pulumi:"federated"`
	Role      string `pulumi:"role"`
	Team      string `pulumi:"team"`
}

// The set of arguments for constructing a Member resource.
type MemberArgs struct {
	Email     pulumi.StringInput
	Federated pulumi.BoolPtrInput
	Role      pulumi.StringInput
	Team      pulumi.StringInput
}

func (MemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*memberArgs)(nil)).Elem()
}

type MemberInput interface {
	pulumi.Input

	ToMemberOutput() MemberOutput
	ToMemberOutputWithContext(ctx context.Context) MemberOutput
}

func (*Member) ElementType() reflect.Type {
	return reflect.TypeOf((**Member)(nil)).Elem()
}

func (i *Member) ToMemberOutput() MemberOutput {
	return i.ToMemberOutputWithContext(context.Background())
}

func (i *Member) ToMemberOutputWithContext(ctx context.Context) MemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberOutput)
}

// MemberArrayInput is an input type that accepts MemberArray and MemberArrayOutput values.
// You can construct a concrete instance of `MemberArrayInput` via:
//
//	MemberArray{ MemberArgs{...} }
type MemberArrayInput interface {
	pulumi.Input

	ToMemberArrayOutput() MemberArrayOutput
	ToMemberArrayOutputWithContext(context.Context) MemberArrayOutput
}

type MemberArray []MemberInput

func (MemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Member)(nil)).Elem()
}

func (i MemberArray) ToMemberArrayOutput() MemberArrayOutput {
	return i.ToMemberArrayOutputWithContext(context.Background())
}

func (i MemberArray) ToMemberArrayOutputWithContext(ctx context.Context) MemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberArrayOutput)
}

// MemberMapInput is an input type that accepts MemberMap and MemberMapOutput values.
// You can construct a concrete instance of `MemberMapInput` via:
//
//	MemberMap{ "key": MemberArgs{...} }
type MemberMapInput interface {
	pulumi.Input

	ToMemberMapOutput() MemberMapOutput
	ToMemberMapOutputWithContext(context.Context) MemberMapOutput
}

type MemberMap map[string]MemberInput

func (MemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Member)(nil)).Elem()
}

func (i MemberMap) ToMemberMapOutput() MemberMapOutput {
	return i.ToMemberMapOutputWithContext(context.Background())
}

func (i MemberMap) ToMemberMapOutputWithContext(ctx context.Context) MemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberMapOutput)
}

type MemberOutput struct{ *pulumi.OutputState }

func (MemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Member)(nil)).Elem()
}

func (o MemberOutput) ToMemberOutput() MemberOutput {
	return o
}

func (o MemberOutput) ToMemberOutputWithContext(ctx context.Context) MemberOutput {
	return o
}

func (o MemberOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

func (o MemberOutput) Federated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Member) pulumi.BoolPtrOutput { return v.Federated }).(pulumi.BoolPtrOutput)
}

func (o MemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o MemberOutput) Team() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.Team }).(pulumi.StringOutput)
}

type MemberArrayOutput struct{ *pulumi.OutputState }

func (MemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Member)(nil)).Elem()
}

func (o MemberArrayOutput) ToMemberArrayOutput() MemberArrayOutput {
	return o
}

func (o MemberArrayOutput) ToMemberArrayOutputWithContext(ctx context.Context) MemberArrayOutput {
	return o
}

func (o MemberArrayOutput) Index(i pulumi.IntInput) MemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Member {
		return vs[0].([]*Member)[vs[1].(int)]
	}).(MemberOutput)
}

type MemberMapOutput struct{ *pulumi.OutputState }

func (MemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Member)(nil)).Elem()
}

func (o MemberMapOutput) ToMemberMapOutput() MemberMapOutput {
	return o
}

func (o MemberMapOutput) ToMemberMapOutputWithContext(ctx context.Context) MemberMapOutput {
	return o
}

func (o MemberMapOutput) MapIndex(k pulumi.StringInput) MemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Member {
		return vs[0].(map[string]*Member)[vs[1].(string)]
	}).(MemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MemberInput)(nil)).Elem(), &Member{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemberArrayInput)(nil)).Elem(), MemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemberMapInput)(nil)).Elem(), MemberMap{})
	pulumi.RegisterOutputType(MemberOutput{})
	pulumi.RegisterOutputType(MemberArrayOutput{})
	pulumi.RegisterOutputType(MemberMapOutput{})
}
