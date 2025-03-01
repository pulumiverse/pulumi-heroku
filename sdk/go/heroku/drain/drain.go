// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package drain

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-heroku/sdk/go/heroku/internal"
)

// Provides a Heroku Drain resource. This can be used to
// create and manage Log Drains on Heroku.
//
// ## Example Usage
//
// ## Importing
//
// When importing a Heroku drain resource, the ID must be built using the app name colon the unique ID from the Heroku API.
// For an app named `production-api` with a drain ID of `b85d9224-310b-409b-891e-c903f5a40568` and the `url` attribute value
// defined for the resource, you would import it as:
//
// When importing a Heroku drain resource, the ID must be built using the app name colon the unique ID from the Heroku API.
// For an app named `production-api` with a drain ID of `b85d9224-310b-409b-891e-c903f5a40568` and the `sensitiveUrl` attribute value
// defined for the resource, you would import it as:
type Drain struct {
	pulumi.CustomResourceState

	// Heroku app ID (do not use app name)
	AppId        pulumi.StringOutput    `pulumi:"appId"`
	SensitiveUrl pulumi.StringPtrOutput `pulumi:"sensitiveUrl"`
	// The unique token for your created drain.
	Token pulumi.StringOutput `pulumi:"token"`
	// The URL for Heroku to drain your logs to. Either `url` or `sensitiveUrl` must be defined.
	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewDrain registers a new resource with the given unique name, arguments, and options.
func NewDrain(ctx *pulumi.Context,
	name string, args *DrainArgs, opts ...pulumi.ResourceOption) (*Drain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.SensitiveUrl != nil {
		args.SensitiveUrl = pulumi.ToSecret(args.SensitiveUrl).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"sensitiveUrl",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Drain
	err := ctx.RegisterResource("heroku:drain/drain:Drain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDrain gets an existing Drain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDrain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DrainState, opts ...pulumi.ResourceOption) (*Drain, error) {
	var resource Drain
	err := ctx.ReadResource("heroku:drain/drain:Drain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Drain resources.
type drainState struct {
	// Heroku app ID (do not use app name)
	AppId        *string `pulumi:"appId"`
	SensitiveUrl *string `pulumi:"sensitiveUrl"`
	// The unique token for your created drain.
	Token *string `pulumi:"token"`
	// The URL for Heroku to drain your logs to. Either `url` or `sensitiveUrl` must be defined.
	Url *string `pulumi:"url"`
}

type DrainState struct {
	// Heroku app ID (do not use app name)
	AppId        pulumi.StringPtrInput
	SensitiveUrl pulumi.StringPtrInput
	// The unique token for your created drain.
	Token pulumi.StringPtrInput
	// The URL for Heroku to drain your logs to. Either `url` or `sensitiveUrl` must be defined.
	Url pulumi.StringPtrInput
}

func (DrainState) ElementType() reflect.Type {
	return reflect.TypeOf((*drainState)(nil)).Elem()
}

type drainArgs struct {
	// Heroku app ID (do not use app name)
	AppId        string  `pulumi:"appId"`
	SensitiveUrl *string `pulumi:"sensitiveUrl"`
	// The URL for Heroku to drain your logs to. Either `url` or `sensitiveUrl` must be defined.
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a Drain resource.
type DrainArgs struct {
	// Heroku app ID (do not use app name)
	AppId        pulumi.StringInput
	SensitiveUrl pulumi.StringPtrInput
	// The URL for Heroku to drain your logs to. Either `url` or `sensitiveUrl` must be defined.
	Url pulumi.StringPtrInput
}

func (DrainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*drainArgs)(nil)).Elem()
}

type DrainInput interface {
	pulumi.Input

	ToDrainOutput() DrainOutput
	ToDrainOutputWithContext(ctx context.Context) DrainOutput
}

func (*Drain) ElementType() reflect.Type {
	return reflect.TypeOf((**Drain)(nil)).Elem()
}

func (i *Drain) ToDrainOutput() DrainOutput {
	return i.ToDrainOutputWithContext(context.Background())
}

func (i *Drain) ToDrainOutputWithContext(ctx context.Context) DrainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrainOutput)
}

// DrainArrayInput is an input type that accepts DrainArray and DrainArrayOutput values.
// You can construct a concrete instance of `DrainArrayInput` via:
//
//	DrainArray{ DrainArgs{...} }
type DrainArrayInput interface {
	pulumi.Input

	ToDrainArrayOutput() DrainArrayOutput
	ToDrainArrayOutputWithContext(context.Context) DrainArrayOutput
}

type DrainArray []DrainInput

func (DrainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Drain)(nil)).Elem()
}

func (i DrainArray) ToDrainArrayOutput() DrainArrayOutput {
	return i.ToDrainArrayOutputWithContext(context.Background())
}

func (i DrainArray) ToDrainArrayOutputWithContext(ctx context.Context) DrainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrainArrayOutput)
}

// DrainMapInput is an input type that accepts DrainMap and DrainMapOutput values.
// You can construct a concrete instance of `DrainMapInput` via:
//
//	DrainMap{ "key": DrainArgs{...} }
type DrainMapInput interface {
	pulumi.Input

	ToDrainMapOutput() DrainMapOutput
	ToDrainMapOutputWithContext(context.Context) DrainMapOutput
}

type DrainMap map[string]DrainInput

func (DrainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Drain)(nil)).Elem()
}

func (i DrainMap) ToDrainMapOutput() DrainMapOutput {
	return i.ToDrainMapOutputWithContext(context.Background())
}

func (i DrainMap) ToDrainMapOutputWithContext(ctx context.Context) DrainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrainMapOutput)
}

type DrainOutput struct{ *pulumi.OutputState }

func (DrainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Drain)(nil)).Elem()
}

func (o DrainOutput) ToDrainOutput() DrainOutput {
	return o
}

func (o DrainOutput) ToDrainOutputWithContext(ctx context.Context) DrainOutput {
	return o
}

// Heroku app ID (do not use app name)
func (o DrainOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *Drain) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

func (o DrainOutput) SensitiveUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Drain) pulumi.StringPtrOutput { return v.SensitiveUrl }).(pulumi.StringPtrOutput)
}

// The unique token for your created drain.
func (o DrainOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *Drain) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// The URL for Heroku to drain your logs to. Either `url` or `sensitiveUrl` must be defined.
func (o DrainOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Drain) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

type DrainArrayOutput struct{ *pulumi.OutputState }

func (DrainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Drain)(nil)).Elem()
}

func (o DrainArrayOutput) ToDrainArrayOutput() DrainArrayOutput {
	return o
}

func (o DrainArrayOutput) ToDrainArrayOutputWithContext(ctx context.Context) DrainArrayOutput {
	return o
}

func (o DrainArrayOutput) Index(i pulumi.IntInput) DrainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Drain {
		return vs[0].([]*Drain)[vs[1].(int)]
	}).(DrainOutput)
}

type DrainMapOutput struct{ *pulumi.OutputState }

func (DrainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Drain)(nil)).Elem()
}

func (o DrainMapOutput) ToDrainMapOutput() DrainMapOutput {
	return o
}

func (o DrainMapOutput) ToDrainMapOutputWithContext(ctx context.Context) DrainMapOutput {
	return o
}

func (o DrainMapOutput) MapIndex(k pulumi.StringInput) DrainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Drain {
		return vs[0].(map[string]*Drain)[vs[1].(string)]
	}).(DrainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DrainInput)(nil)).Elem(), &Drain{})
	pulumi.RegisterInputType(reflect.TypeOf((*DrainArrayInput)(nil)).Elem(), DrainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DrainMapInput)(nil)).Elem(), DrainMap{})
	pulumi.RegisterOutputType(DrainOutput{})
	pulumi.RegisterOutputType(DrainArrayOutput{})
	pulumi.RegisterOutputType(DrainMapOutput{})
}
