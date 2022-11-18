// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ssl struct {
	pulumi.CustomResourceState

	AppId            pulumi.StringOutput    `pulumi:"appId"`
	CertificateChain pulumi.StringOutput    `pulumi:"certificateChain"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	PrivateKey       pulumi.StringPtrOutput `pulumi:"privateKey"`
}

// NewSsl registers a new resource with the given unique name, arguments, and options.
func NewSsl(ctx *pulumi.Context,
	name string, args *SslArgs, opts ...pulumi.ResourceOption) (*Ssl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.CertificateChain == nil {
		return nil, errors.New("invalid value for required argument 'CertificateChain'")
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Ssl
	err := ctx.RegisterResource("heroku:ssl/ssl:Ssl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSsl gets an existing Ssl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSsl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslState, opts ...pulumi.ResourceOption) (*Ssl, error) {
	var resource Ssl
	err := ctx.ReadResource("heroku:ssl/ssl:Ssl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ssl resources.
type sslState struct {
	AppId            *string `pulumi:"appId"`
	CertificateChain *string `pulumi:"certificateChain"`
	Name             *string `pulumi:"name"`
	PrivateKey       *string `pulumi:"privateKey"`
}

type SslState struct {
	AppId            pulumi.StringPtrInput
	CertificateChain pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	PrivateKey       pulumi.StringPtrInput
}

func (SslState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslState)(nil)).Elem()
}

type sslArgs struct {
	AppId            string  `pulumi:"appId"`
	CertificateChain string  `pulumi:"certificateChain"`
	PrivateKey       *string `pulumi:"privateKey"`
}

// The set of arguments for constructing a Ssl resource.
type SslArgs struct {
	AppId            pulumi.StringInput
	CertificateChain pulumi.StringInput
	PrivateKey       pulumi.StringPtrInput
}

func (SslArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslArgs)(nil)).Elem()
}

type SslInput interface {
	pulumi.Input

	ToSslOutput() SslOutput
	ToSslOutputWithContext(ctx context.Context) SslOutput
}

func (*Ssl) ElementType() reflect.Type {
	return reflect.TypeOf((**Ssl)(nil)).Elem()
}

func (i *Ssl) ToSslOutput() SslOutput {
	return i.ToSslOutputWithContext(context.Background())
}

func (i *Ssl) ToSslOutputWithContext(ctx context.Context) SslOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslOutput)
}

// SslArrayInput is an input type that accepts SslArray and SslArrayOutput values.
// You can construct a concrete instance of `SslArrayInput` via:
//
//          SslArray{ SslArgs{...} }
type SslArrayInput interface {
	pulumi.Input

	ToSslArrayOutput() SslArrayOutput
	ToSslArrayOutputWithContext(context.Context) SslArrayOutput
}

type SslArray []SslInput

func (SslArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ssl)(nil)).Elem()
}

func (i SslArray) ToSslArrayOutput() SslArrayOutput {
	return i.ToSslArrayOutputWithContext(context.Background())
}

func (i SslArray) ToSslArrayOutputWithContext(ctx context.Context) SslArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslArrayOutput)
}

// SslMapInput is an input type that accepts SslMap and SslMapOutput values.
// You can construct a concrete instance of `SslMapInput` via:
//
//          SslMap{ "key": SslArgs{...} }
type SslMapInput interface {
	pulumi.Input

	ToSslMapOutput() SslMapOutput
	ToSslMapOutputWithContext(context.Context) SslMapOutput
}

type SslMap map[string]SslInput

func (SslMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ssl)(nil)).Elem()
}

func (i SslMap) ToSslMapOutput() SslMapOutput {
	return i.ToSslMapOutputWithContext(context.Background())
}

func (i SslMap) ToSslMapOutputWithContext(ctx context.Context) SslMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslMapOutput)
}

type SslOutput struct{ *pulumi.OutputState }

func (SslOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ssl)(nil)).Elem()
}

func (o SslOutput) ToSslOutput() SslOutput {
	return o
}

func (o SslOutput) ToSslOutputWithContext(ctx context.Context) SslOutput {
	return o
}

func (o SslOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

func (o SslOutput) CertificateChain() pulumi.StringOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringOutput { return v.CertificateChain }).(pulumi.StringOutput)
}

func (o SslOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SslOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringPtrOutput { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

type SslArrayOutput struct{ *pulumi.OutputState }

func (SslArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ssl)(nil)).Elem()
}

func (o SslArrayOutput) ToSslArrayOutput() SslArrayOutput {
	return o
}

func (o SslArrayOutput) ToSslArrayOutputWithContext(ctx context.Context) SslArrayOutput {
	return o
}

func (o SslArrayOutput) Index(i pulumi.IntInput) SslOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ssl {
		return vs[0].([]*Ssl)[vs[1].(int)]
	}).(SslOutput)
}

type SslMapOutput struct{ *pulumi.OutputState }

func (SslMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ssl)(nil)).Elem()
}

func (o SslMapOutput) ToSslMapOutput() SslMapOutput {
	return o
}

func (o SslMapOutput) ToSslMapOutputWithContext(ctx context.Context) SslMapOutput {
	return o
}

func (o SslMapOutput) MapIndex(k pulumi.StringInput) SslOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ssl {
		return vs[0].(map[string]*Ssl)[vs[1].(string)]
	}).(SslOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SslInput)(nil)).Elem(), &Ssl{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslArrayInput)(nil)).Elem(), SslArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslMapInput)(nil)).Elem(), SslMap{})
	pulumi.RegisterOutputType(SslOutput{})
	pulumi.RegisterOutputType(SslArrayOutput{})
	pulumi.RegisterOutputType(SslMapOutput{})
}