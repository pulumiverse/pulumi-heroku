// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package space

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-heroku/sdk/go/heroku/internal"
)

type VpnConnection struct {
	pulumi.CustomResourceState

	IkeVersion     pulumi.IntOutput               `pulumi:"ikeVersion"`
	Name           pulumi.StringOutput            `pulumi:"name"`
	PublicIp       pulumi.StringOutput            `pulumi:"publicIp"`
	RoutableCidrs  pulumi.StringArrayOutput       `pulumi:"routableCidrs"`
	Space          pulumi.StringOutput            `pulumi:"space"`
	SpaceCidrBlock pulumi.StringOutput            `pulumi:"spaceCidrBlock"`
	Tunnels        VpnConnectionTunnelArrayOutput `pulumi:"tunnels"`
}

// NewVpnConnection registers a new resource with the given unique name, arguments, and options.
func NewVpnConnection(ctx *pulumi.Context,
	name string, args *VpnConnectionArgs, opts ...pulumi.ResourceOption) (*VpnConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicIp == nil {
		return nil, errors.New("invalid value for required argument 'PublicIp'")
	}
	if args.RoutableCidrs == nil {
		return nil, errors.New("invalid value for required argument 'RoutableCidrs'")
	}
	if args.Space == nil {
		return nil, errors.New("invalid value for required argument 'Space'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpnConnection
	err := ctx.RegisterResource("heroku:space/vpnConnection:VpnConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnConnection gets an existing VpnConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnConnectionState, opts ...pulumi.ResourceOption) (*VpnConnection, error) {
	var resource VpnConnection
	err := ctx.ReadResource("heroku:space/vpnConnection:VpnConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnConnection resources.
type vpnConnectionState struct {
	IkeVersion     *int                  `pulumi:"ikeVersion"`
	Name           *string               `pulumi:"name"`
	PublicIp       *string               `pulumi:"publicIp"`
	RoutableCidrs  []string              `pulumi:"routableCidrs"`
	Space          *string               `pulumi:"space"`
	SpaceCidrBlock *string               `pulumi:"spaceCidrBlock"`
	Tunnels        []VpnConnectionTunnel `pulumi:"tunnels"`
}

type VpnConnectionState struct {
	IkeVersion     pulumi.IntPtrInput
	Name           pulumi.StringPtrInput
	PublicIp       pulumi.StringPtrInput
	RoutableCidrs  pulumi.StringArrayInput
	Space          pulumi.StringPtrInput
	SpaceCidrBlock pulumi.StringPtrInput
	Tunnels        VpnConnectionTunnelArrayInput
}

func (VpnConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionState)(nil)).Elem()
}

type vpnConnectionArgs struct {
	Name          *string               `pulumi:"name"`
	PublicIp      string                `pulumi:"publicIp"`
	RoutableCidrs []string              `pulumi:"routableCidrs"`
	Space         string                `pulumi:"space"`
	Tunnels       []VpnConnectionTunnel `pulumi:"tunnels"`
}

// The set of arguments for constructing a VpnConnection resource.
type VpnConnectionArgs struct {
	Name          pulumi.StringPtrInput
	PublicIp      pulumi.StringInput
	RoutableCidrs pulumi.StringArrayInput
	Space         pulumi.StringInput
	Tunnels       VpnConnectionTunnelArrayInput
}

func (VpnConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionArgs)(nil)).Elem()
}

type VpnConnectionInput interface {
	pulumi.Input

	ToVpnConnectionOutput() VpnConnectionOutput
	ToVpnConnectionOutputWithContext(ctx context.Context) VpnConnectionOutput
}

func (*VpnConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnConnection)(nil)).Elem()
}

func (i *VpnConnection) ToVpnConnectionOutput() VpnConnectionOutput {
	return i.ToVpnConnectionOutputWithContext(context.Background())
}

func (i *VpnConnection) ToVpnConnectionOutputWithContext(ctx context.Context) VpnConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionOutput)
}

// VpnConnectionArrayInput is an input type that accepts VpnConnectionArray and VpnConnectionArrayOutput values.
// You can construct a concrete instance of `VpnConnectionArrayInput` via:
//
//	VpnConnectionArray{ VpnConnectionArgs{...} }
type VpnConnectionArrayInput interface {
	pulumi.Input

	ToVpnConnectionArrayOutput() VpnConnectionArrayOutput
	ToVpnConnectionArrayOutputWithContext(context.Context) VpnConnectionArrayOutput
}

type VpnConnectionArray []VpnConnectionInput

func (VpnConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnConnection)(nil)).Elem()
}

func (i VpnConnectionArray) ToVpnConnectionArrayOutput() VpnConnectionArrayOutput {
	return i.ToVpnConnectionArrayOutputWithContext(context.Background())
}

func (i VpnConnectionArray) ToVpnConnectionArrayOutputWithContext(ctx context.Context) VpnConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionArrayOutput)
}

// VpnConnectionMapInput is an input type that accepts VpnConnectionMap and VpnConnectionMapOutput values.
// You can construct a concrete instance of `VpnConnectionMapInput` via:
//
//	VpnConnectionMap{ "key": VpnConnectionArgs{...} }
type VpnConnectionMapInput interface {
	pulumi.Input

	ToVpnConnectionMapOutput() VpnConnectionMapOutput
	ToVpnConnectionMapOutputWithContext(context.Context) VpnConnectionMapOutput
}

type VpnConnectionMap map[string]VpnConnectionInput

func (VpnConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnConnection)(nil)).Elem()
}

func (i VpnConnectionMap) ToVpnConnectionMapOutput() VpnConnectionMapOutput {
	return i.ToVpnConnectionMapOutputWithContext(context.Background())
}

func (i VpnConnectionMap) ToVpnConnectionMapOutputWithContext(ctx context.Context) VpnConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionMapOutput)
}

type VpnConnectionOutput struct{ *pulumi.OutputState }

func (VpnConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnConnection)(nil)).Elem()
}

func (o VpnConnectionOutput) ToVpnConnectionOutput() VpnConnectionOutput {
	return o
}

func (o VpnConnectionOutput) ToVpnConnectionOutputWithContext(ctx context.Context) VpnConnectionOutput {
	return o
}

func (o VpnConnectionOutput) IkeVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.IntOutput { return v.IkeVersion }).(pulumi.IntOutput)
}

func (o VpnConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VpnConnectionOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.StringOutput { return v.PublicIp }).(pulumi.StringOutput)
}

func (o VpnConnectionOutput) RoutableCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.StringArrayOutput { return v.RoutableCidrs }).(pulumi.StringArrayOutput)
}

func (o VpnConnectionOutput) Space() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.StringOutput { return v.Space }).(pulumi.StringOutput)
}

func (o VpnConnectionOutput) SpaceCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.StringOutput { return v.SpaceCidrBlock }).(pulumi.StringOutput)
}

func (o VpnConnectionOutput) Tunnels() VpnConnectionTunnelArrayOutput {
	return o.ApplyT(func(v *VpnConnection) VpnConnectionTunnelArrayOutput { return v.Tunnels }).(VpnConnectionTunnelArrayOutput)
}

type VpnConnectionArrayOutput struct{ *pulumi.OutputState }

func (VpnConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnConnection)(nil)).Elem()
}

func (o VpnConnectionArrayOutput) ToVpnConnectionArrayOutput() VpnConnectionArrayOutput {
	return o
}

func (o VpnConnectionArrayOutput) ToVpnConnectionArrayOutputWithContext(ctx context.Context) VpnConnectionArrayOutput {
	return o
}

func (o VpnConnectionArrayOutput) Index(i pulumi.IntInput) VpnConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnConnection {
		return vs[0].([]*VpnConnection)[vs[1].(int)]
	}).(VpnConnectionOutput)
}

type VpnConnectionMapOutput struct{ *pulumi.OutputState }

func (VpnConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnConnection)(nil)).Elem()
}

func (o VpnConnectionMapOutput) ToVpnConnectionMapOutput() VpnConnectionMapOutput {
	return o
}

func (o VpnConnectionMapOutput) ToVpnConnectionMapOutputWithContext(ctx context.Context) VpnConnectionMapOutput {
	return o
}

func (o VpnConnectionMapOutput) MapIndex(k pulumi.StringInput) VpnConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnConnection {
		return vs[0].(map[string]*VpnConnection)[vs[1].(string)]
	}).(VpnConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnConnectionInput)(nil)).Elem(), &VpnConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnConnectionArrayInput)(nil)).Elem(), VpnConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnConnectionMapInput)(nil)).Elem(), VpnConnectionMap{})
	pulumi.RegisterOutputType(VpnConnectionOutput{})
	pulumi.RegisterOutputType(VpnConnectionArrayOutput{})
	pulumi.RegisterOutputType(VpnConnectionMapOutput{})
}
