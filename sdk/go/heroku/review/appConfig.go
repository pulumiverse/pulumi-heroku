// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package review

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-heroku/sdk/go/heroku/internal"
)

// Provides a resource for configuring review apps. Using this resource also enables review apps for a pipeline.
//
// > **IMPORTANT!**
// This resource can only be used after you create a pipeline, and the pipeline has been connected to a Github repository.
// Please visit this [help article](https://devcenter.heroku.com/articles/github-integration-review-apps#setup)
// for more information.
//
// ## Example Usage
//
// ## Import
//
// An Existing review app config using the combination of the pipeline UUID and the Github organization/repository
//
// separated by a colon.
//
// ```sh
// $ pulumi import heroku:review/appConfig:AppConfig foobar afd193fb-7c5a-4d8f-afad-2388f4e6049d:heroku/homebrew-brew
// ```
type AppConfig struct {
	pulumi.CustomResourceState

	// If true, this will trigger the creation of review apps when pull-requests
	// are opened in the repo. Defaults to `false`.
	AutomaticReviewApps pulumi.BoolPtrOutput `pulumi:"automaticReviewApps"`
	// A unique prefix that will be used to create review app names.
	BaseName pulumi.StringOutput `pulumi:"baseName"`
	// Provides a key/value pair to specify whether to use a common runtime or a private space.
	DeployTarget AppConfigDeployTargetOutput `pulumi:"deployTarget"`
	// If `true`, this will trigger automatic deletion of review apps when they’re stale.
	// Defaults to `false`.
	DestroyStaleApps pulumi.BoolPtrOutput `pulumi:"destroyStaleApps"`
	// The fullName of the repository that you want to enable review-apps from.
	// Example `heroku/homebrew-brew`.
	OrgRepo pulumi.StringOutput `pulumi:"orgRepo"`
	// The UUID of an existing pipeline.
	PipelineId pulumi.StringOutput `pulumi:"pipelineId"`
	RepoId     pulumi.IntOutput    `pulumi:"repoId"`
	// Destroy stale review apps automatically after these many days without any deploys.
	// Must be set with `destroyStaleApps` and value needs to be between `1` and `30` inclusive.
	StaleDays pulumi.IntOutput `pulumi:"staleDays"`
	// If true, review apps will only be created when CI passes. Defaults to `false`.
	WaitForCi pulumi.BoolPtrOutput `pulumi:"waitForCi"`
}

// NewAppConfig registers a new resource with the given unique name, arguments, and options.
func NewAppConfig(ctx *pulumi.Context,
	name string, args *AppConfigArgs, opts ...pulumi.ResourceOption) (*AppConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeployTarget == nil {
		return nil, errors.New("invalid value for required argument 'DeployTarget'")
	}
	if args.OrgRepo == nil {
		return nil, errors.New("invalid value for required argument 'OrgRepo'")
	}
	if args.PipelineId == nil {
		return nil, errors.New("invalid value for required argument 'PipelineId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppConfig
	err := ctx.RegisterResource("heroku:review/appConfig:AppConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppConfig gets an existing AppConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppConfigState, opts ...pulumi.ResourceOption) (*AppConfig, error) {
	var resource AppConfig
	err := ctx.ReadResource("heroku:review/appConfig:AppConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppConfig resources.
type appConfigState struct {
	// If true, this will trigger the creation of review apps when pull-requests
	// are opened in the repo. Defaults to `false`.
	AutomaticReviewApps *bool `pulumi:"automaticReviewApps"`
	// A unique prefix that will be used to create review app names.
	BaseName *string `pulumi:"baseName"`
	// Provides a key/value pair to specify whether to use a common runtime or a private space.
	DeployTarget *AppConfigDeployTarget `pulumi:"deployTarget"`
	// If `true`, this will trigger automatic deletion of review apps when they’re stale.
	// Defaults to `false`.
	DestroyStaleApps *bool `pulumi:"destroyStaleApps"`
	// The fullName of the repository that you want to enable review-apps from.
	// Example `heroku/homebrew-brew`.
	OrgRepo *string `pulumi:"orgRepo"`
	// The UUID of an existing pipeline.
	PipelineId *string `pulumi:"pipelineId"`
	RepoId     *int    `pulumi:"repoId"`
	// Destroy stale review apps automatically after these many days without any deploys.
	// Must be set with `destroyStaleApps` and value needs to be between `1` and `30` inclusive.
	StaleDays *int `pulumi:"staleDays"`
	// If true, review apps will only be created when CI passes. Defaults to `false`.
	WaitForCi *bool `pulumi:"waitForCi"`
}

type AppConfigState struct {
	// If true, this will trigger the creation of review apps when pull-requests
	// are opened in the repo. Defaults to `false`.
	AutomaticReviewApps pulumi.BoolPtrInput
	// A unique prefix that will be used to create review app names.
	BaseName pulumi.StringPtrInput
	// Provides a key/value pair to specify whether to use a common runtime or a private space.
	DeployTarget AppConfigDeployTargetPtrInput
	// If `true`, this will trigger automatic deletion of review apps when they’re stale.
	// Defaults to `false`.
	DestroyStaleApps pulumi.BoolPtrInput
	// The fullName of the repository that you want to enable review-apps from.
	// Example `heroku/homebrew-brew`.
	OrgRepo pulumi.StringPtrInput
	// The UUID of an existing pipeline.
	PipelineId pulumi.StringPtrInput
	RepoId     pulumi.IntPtrInput
	// Destroy stale review apps automatically after these many days without any deploys.
	// Must be set with `destroyStaleApps` and value needs to be between `1` and `30` inclusive.
	StaleDays pulumi.IntPtrInput
	// If true, review apps will only be created when CI passes. Defaults to `false`.
	WaitForCi pulumi.BoolPtrInput
}

func (AppConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*appConfigState)(nil)).Elem()
}

type appConfigArgs struct {
	// If true, this will trigger the creation of review apps when pull-requests
	// are opened in the repo. Defaults to `false`.
	AutomaticReviewApps *bool `pulumi:"automaticReviewApps"`
	// A unique prefix that will be used to create review app names.
	BaseName *string `pulumi:"baseName"`
	// Provides a key/value pair to specify whether to use a common runtime or a private space.
	DeployTarget AppConfigDeployTarget `pulumi:"deployTarget"`
	// If `true`, this will trigger automatic deletion of review apps when they’re stale.
	// Defaults to `false`.
	DestroyStaleApps *bool `pulumi:"destroyStaleApps"`
	// The fullName of the repository that you want to enable review-apps from.
	// Example `heroku/homebrew-brew`.
	OrgRepo string `pulumi:"orgRepo"`
	// The UUID of an existing pipeline.
	PipelineId string `pulumi:"pipelineId"`
	// Destroy stale review apps automatically after these many days without any deploys.
	// Must be set with `destroyStaleApps` and value needs to be between `1` and `30` inclusive.
	StaleDays *int `pulumi:"staleDays"`
	// If true, review apps will only be created when CI passes. Defaults to `false`.
	WaitForCi *bool `pulumi:"waitForCi"`
}

// The set of arguments for constructing a AppConfig resource.
type AppConfigArgs struct {
	// If true, this will trigger the creation of review apps when pull-requests
	// are opened in the repo. Defaults to `false`.
	AutomaticReviewApps pulumi.BoolPtrInput
	// A unique prefix that will be used to create review app names.
	BaseName pulumi.StringPtrInput
	// Provides a key/value pair to specify whether to use a common runtime or a private space.
	DeployTarget AppConfigDeployTargetInput
	// If `true`, this will trigger automatic deletion of review apps when they’re stale.
	// Defaults to `false`.
	DestroyStaleApps pulumi.BoolPtrInput
	// The fullName of the repository that you want to enable review-apps from.
	// Example `heroku/homebrew-brew`.
	OrgRepo pulumi.StringInput
	// The UUID of an existing pipeline.
	PipelineId pulumi.StringInput
	// Destroy stale review apps automatically after these many days without any deploys.
	// Must be set with `destroyStaleApps` and value needs to be between `1` and `30` inclusive.
	StaleDays pulumi.IntPtrInput
	// If true, review apps will only be created when CI passes. Defaults to `false`.
	WaitForCi pulumi.BoolPtrInput
}

func (AppConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appConfigArgs)(nil)).Elem()
}

type AppConfigInput interface {
	pulumi.Input

	ToAppConfigOutput() AppConfigOutput
	ToAppConfigOutputWithContext(ctx context.Context) AppConfigOutput
}

func (*AppConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**AppConfig)(nil)).Elem()
}

func (i *AppConfig) ToAppConfigOutput() AppConfigOutput {
	return i.ToAppConfigOutputWithContext(context.Background())
}

func (i *AppConfig) ToAppConfigOutputWithContext(ctx context.Context) AppConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppConfigOutput)
}

// AppConfigArrayInput is an input type that accepts AppConfigArray and AppConfigArrayOutput values.
// You can construct a concrete instance of `AppConfigArrayInput` via:
//
//	AppConfigArray{ AppConfigArgs{...} }
type AppConfigArrayInput interface {
	pulumi.Input

	ToAppConfigArrayOutput() AppConfigArrayOutput
	ToAppConfigArrayOutputWithContext(context.Context) AppConfigArrayOutput
}

type AppConfigArray []AppConfigInput

func (AppConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppConfig)(nil)).Elem()
}

func (i AppConfigArray) ToAppConfigArrayOutput() AppConfigArrayOutput {
	return i.ToAppConfigArrayOutputWithContext(context.Background())
}

func (i AppConfigArray) ToAppConfigArrayOutputWithContext(ctx context.Context) AppConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppConfigArrayOutput)
}

// AppConfigMapInput is an input type that accepts AppConfigMap and AppConfigMapOutput values.
// You can construct a concrete instance of `AppConfigMapInput` via:
//
//	AppConfigMap{ "key": AppConfigArgs{...} }
type AppConfigMapInput interface {
	pulumi.Input

	ToAppConfigMapOutput() AppConfigMapOutput
	ToAppConfigMapOutputWithContext(context.Context) AppConfigMapOutput
}

type AppConfigMap map[string]AppConfigInput

func (AppConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppConfig)(nil)).Elem()
}

func (i AppConfigMap) ToAppConfigMapOutput() AppConfigMapOutput {
	return i.ToAppConfigMapOutputWithContext(context.Background())
}

func (i AppConfigMap) ToAppConfigMapOutputWithContext(ctx context.Context) AppConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppConfigMapOutput)
}

type AppConfigOutput struct{ *pulumi.OutputState }

func (AppConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppConfig)(nil)).Elem()
}

func (o AppConfigOutput) ToAppConfigOutput() AppConfigOutput {
	return o
}

func (o AppConfigOutput) ToAppConfigOutputWithContext(ctx context.Context) AppConfigOutput {
	return o
}

// If true, this will trigger the creation of review apps when pull-requests
// are opened in the repo. Defaults to `false`.
func (o AppConfigOutput) AutomaticReviewApps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppConfig) pulumi.BoolPtrOutput { return v.AutomaticReviewApps }).(pulumi.BoolPtrOutput)
}

// A unique prefix that will be used to create review app names.
func (o AppConfigOutput) BaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConfig) pulumi.StringOutput { return v.BaseName }).(pulumi.StringOutput)
}

// Provides a key/value pair to specify whether to use a common runtime or a private space.
func (o AppConfigOutput) DeployTarget() AppConfigDeployTargetOutput {
	return o.ApplyT(func(v *AppConfig) AppConfigDeployTargetOutput { return v.DeployTarget }).(AppConfigDeployTargetOutput)
}

// If `true`, this will trigger automatic deletion of review apps when they’re stale.
// Defaults to `false`.
func (o AppConfigOutput) DestroyStaleApps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppConfig) pulumi.BoolPtrOutput { return v.DestroyStaleApps }).(pulumi.BoolPtrOutput)
}

// The fullName of the repository that you want to enable review-apps from.
// Example `heroku/homebrew-brew`.
func (o AppConfigOutput) OrgRepo() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConfig) pulumi.StringOutput { return v.OrgRepo }).(pulumi.StringOutput)
}

// The UUID of an existing pipeline.
func (o AppConfigOutput) PipelineId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConfig) pulumi.StringOutput { return v.PipelineId }).(pulumi.StringOutput)
}

func (o AppConfigOutput) RepoId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppConfig) pulumi.IntOutput { return v.RepoId }).(pulumi.IntOutput)
}

// Destroy stale review apps automatically after these many days without any deploys.
// Must be set with `destroyStaleApps` and value needs to be between `1` and `30` inclusive.
func (o AppConfigOutput) StaleDays() pulumi.IntOutput {
	return o.ApplyT(func(v *AppConfig) pulumi.IntOutput { return v.StaleDays }).(pulumi.IntOutput)
}

// If true, review apps will only be created when CI passes. Defaults to `false`.
func (o AppConfigOutput) WaitForCi() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppConfig) pulumi.BoolPtrOutput { return v.WaitForCi }).(pulumi.BoolPtrOutput)
}

type AppConfigArrayOutput struct{ *pulumi.OutputState }

func (AppConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppConfig)(nil)).Elem()
}

func (o AppConfigArrayOutput) ToAppConfigArrayOutput() AppConfigArrayOutput {
	return o
}

func (o AppConfigArrayOutput) ToAppConfigArrayOutputWithContext(ctx context.Context) AppConfigArrayOutput {
	return o
}

func (o AppConfigArrayOutput) Index(i pulumi.IntInput) AppConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppConfig {
		return vs[0].([]*AppConfig)[vs[1].(int)]
	}).(AppConfigOutput)
}

type AppConfigMapOutput struct{ *pulumi.OutputState }

func (AppConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppConfig)(nil)).Elem()
}

func (o AppConfigMapOutput) ToAppConfigMapOutput() AppConfigMapOutput {
	return o
}

func (o AppConfigMapOutput) ToAppConfigMapOutputWithContext(ctx context.Context) AppConfigMapOutput {
	return o
}

func (o AppConfigMapOutput) MapIndex(k pulumi.StringInput) AppConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppConfig {
		return vs[0].(map[string]*AppConfig)[vs[1].(string)]
	}).(AppConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppConfigInput)(nil)).Elem(), &AppConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppConfigArrayInput)(nil)).Elem(), AppConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppConfigMapInput)(nil)).Elem(), AppConfigMap{})
	pulumi.RegisterOutputType(AppConfigOutput{})
	pulumi.RegisterOutputType(AppConfigArrayOutput{})
	pulumi.RegisterOutputType(AppConfigMapOutput{})
}
