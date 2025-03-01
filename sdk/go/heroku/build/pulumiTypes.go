// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package build

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-heroku/sdk/go/heroku/internal"
)

var _ = internal.GetEnvOrDefault

type BuildSource struct {
	// SHA256 hash of the tarball archive to verify its integrity, example:
	// `SHA256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
	Checksum *string `pulumi:"checksum"`
	// Local path to the source directory or tarball archive for the app
	Path *string `pulumi:"path"`
	// `https` location of the source archive for the app
	Url *string `pulumi:"url"`
	// Use to track what version of your source originated this build. If you are creating builds
	// from git-versioned source code, for example, the commit hash, or release tag would be a good value to use for the
	// version parameter.
	Version *string `pulumi:"version"`
}

// BuildSourceInput is an input type that accepts BuildSourceArgs and BuildSourceOutput values.
// You can construct a concrete instance of `BuildSourceInput` via:
//
//	BuildSourceArgs{...}
type BuildSourceInput interface {
	pulumi.Input

	ToBuildSourceOutput() BuildSourceOutput
	ToBuildSourceOutputWithContext(context.Context) BuildSourceOutput
}

type BuildSourceArgs struct {
	// SHA256 hash of the tarball archive to verify its integrity, example:
	// `SHA256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
	Checksum pulumi.StringPtrInput `pulumi:"checksum"`
	// Local path to the source directory or tarball archive for the app
	Path pulumi.StringPtrInput `pulumi:"path"`
	// `https` location of the source archive for the app
	Url pulumi.StringPtrInput `pulumi:"url"`
	// Use to track what version of your source originated this build. If you are creating builds
	// from git-versioned source code, for example, the commit hash, or release tag would be a good value to use for the
	// version parameter.
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (BuildSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildSource)(nil)).Elem()
}

func (i BuildSourceArgs) ToBuildSourceOutput() BuildSourceOutput {
	return i.ToBuildSourceOutputWithContext(context.Background())
}

func (i BuildSourceArgs) ToBuildSourceOutputWithContext(ctx context.Context) BuildSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildSourceOutput)
}

func (i BuildSourceArgs) ToBuildSourcePtrOutput() BuildSourcePtrOutput {
	return i.ToBuildSourcePtrOutputWithContext(context.Background())
}

func (i BuildSourceArgs) ToBuildSourcePtrOutputWithContext(ctx context.Context) BuildSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildSourceOutput).ToBuildSourcePtrOutputWithContext(ctx)
}

// BuildSourcePtrInput is an input type that accepts BuildSourceArgs, BuildSourcePtr and BuildSourcePtrOutput values.
// You can construct a concrete instance of `BuildSourcePtrInput` via:
//
//	        BuildSourceArgs{...}
//
//	or:
//
//	        nil
type BuildSourcePtrInput interface {
	pulumi.Input

	ToBuildSourcePtrOutput() BuildSourcePtrOutput
	ToBuildSourcePtrOutputWithContext(context.Context) BuildSourcePtrOutput
}

type buildSourcePtrType BuildSourceArgs

func BuildSourcePtr(v *BuildSourceArgs) BuildSourcePtrInput {
	return (*buildSourcePtrType)(v)
}

func (*buildSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildSource)(nil)).Elem()
}

func (i *buildSourcePtrType) ToBuildSourcePtrOutput() BuildSourcePtrOutput {
	return i.ToBuildSourcePtrOutputWithContext(context.Background())
}

func (i *buildSourcePtrType) ToBuildSourcePtrOutputWithContext(ctx context.Context) BuildSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildSourcePtrOutput)
}

type BuildSourceOutput struct{ *pulumi.OutputState }

func (BuildSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildSource)(nil)).Elem()
}

func (o BuildSourceOutput) ToBuildSourceOutput() BuildSourceOutput {
	return o
}

func (o BuildSourceOutput) ToBuildSourceOutputWithContext(ctx context.Context) BuildSourceOutput {
	return o
}

func (o BuildSourceOutput) ToBuildSourcePtrOutput() BuildSourcePtrOutput {
	return o.ToBuildSourcePtrOutputWithContext(context.Background())
}

func (o BuildSourceOutput) ToBuildSourcePtrOutputWithContext(ctx context.Context) BuildSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BuildSource) *BuildSource {
		return &v
	}).(BuildSourcePtrOutput)
}

// SHA256 hash of the tarball archive to verify its integrity, example:
// `SHA256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
func (o BuildSourceOutput) Checksum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildSource) *string { return v.Checksum }).(pulumi.StringPtrOutput)
}

// Local path to the source directory or tarball archive for the app
func (o BuildSourceOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildSource) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// `https` location of the source archive for the app
func (o BuildSourceOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildSource) *string { return v.Url }).(pulumi.StringPtrOutput)
}

// Use to track what version of your source originated this build. If you are creating builds
// from git-versioned source code, for example, the commit hash, or release tag would be a good value to use for the
// version parameter.
func (o BuildSourceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildSource) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type BuildSourcePtrOutput struct{ *pulumi.OutputState }

func (BuildSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildSource)(nil)).Elem()
}

func (o BuildSourcePtrOutput) ToBuildSourcePtrOutput() BuildSourcePtrOutput {
	return o
}

func (o BuildSourcePtrOutput) ToBuildSourcePtrOutputWithContext(ctx context.Context) BuildSourcePtrOutput {
	return o
}

func (o BuildSourcePtrOutput) Elem() BuildSourceOutput {
	return o.ApplyT(func(v *BuildSource) BuildSource {
		if v != nil {
			return *v
		}
		var ret BuildSource
		return ret
	}).(BuildSourceOutput)
}

// SHA256 hash of the tarball archive to verify its integrity, example:
// `SHA256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
func (o BuildSourcePtrOutput) Checksum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildSource) *string {
		if v == nil {
			return nil
		}
		return v.Checksum
	}).(pulumi.StringPtrOutput)
}

// Local path to the source directory or tarball archive for the app
func (o BuildSourcePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildSource) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

// `https` location of the source archive for the app
func (o BuildSourcePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildSource) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

// Use to track what version of your source originated this build. If you are creating builds
// from git-versioned source code, for example, the commit hash, or release tag would be a good value to use for the
// version parameter.
func (o BuildSourcePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildSource) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type BuildUser struct {
	Email *string `pulumi:"email"`
	Id    *string `pulumi:"id"`
}

// BuildUserInput is an input type that accepts BuildUserArgs and BuildUserOutput values.
// You can construct a concrete instance of `BuildUserInput` via:
//
//	BuildUserArgs{...}
type BuildUserInput interface {
	pulumi.Input

	ToBuildUserOutput() BuildUserOutput
	ToBuildUserOutputWithContext(context.Context) BuildUserOutput
}

type BuildUserArgs struct {
	Email pulumi.StringPtrInput `pulumi:"email"`
	Id    pulumi.StringPtrInput `pulumi:"id"`
}

func (BuildUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildUser)(nil)).Elem()
}

func (i BuildUserArgs) ToBuildUserOutput() BuildUserOutput {
	return i.ToBuildUserOutputWithContext(context.Background())
}

func (i BuildUserArgs) ToBuildUserOutputWithContext(ctx context.Context) BuildUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildUserOutput)
}

// BuildUserArrayInput is an input type that accepts BuildUserArray and BuildUserArrayOutput values.
// You can construct a concrete instance of `BuildUserArrayInput` via:
//
//	BuildUserArray{ BuildUserArgs{...} }
type BuildUserArrayInput interface {
	pulumi.Input

	ToBuildUserArrayOutput() BuildUserArrayOutput
	ToBuildUserArrayOutputWithContext(context.Context) BuildUserArrayOutput
}

type BuildUserArray []BuildUserInput

func (BuildUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildUser)(nil)).Elem()
}

func (i BuildUserArray) ToBuildUserArrayOutput() BuildUserArrayOutput {
	return i.ToBuildUserArrayOutputWithContext(context.Background())
}

func (i BuildUserArray) ToBuildUserArrayOutputWithContext(ctx context.Context) BuildUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildUserArrayOutput)
}

type BuildUserOutput struct{ *pulumi.OutputState }

func (BuildUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildUser)(nil)).Elem()
}

func (o BuildUserOutput) ToBuildUserOutput() BuildUserOutput {
	return o
}

func (o BuildUserOutput) ToBuildUserOutputWithContext(ctx context.Context) BuildUserOutput {
	return o
}

func (o BuildUserOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildUser) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o BuildUserOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildUser) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type BuildUserArrayOutput struct{ *pulumi.OutputState }

func (BuildUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildUser)(nil)).Elem()
}

func (o BuildUserArrayOutput) ToBuildUserArrayOutput() BuildUserArrayOutput {
	return o
}

func (o BuildUserArrayOutput) ToBuildUserArrayOutputWithContext(ctx context.Context) BuildUserArrayOutput {
	return o
}

func (o BuildUserArrayOutput) Index(i pulumi.IntInput) BuildUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BuildUser {
		return vs[0].([]BuildUser)[vs[1].(int)]
	}).(BuildUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BuildSourceInput)(nil)).Elem(), BuildSourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildSourcePtrInput)(nil)).Elem(), BuildSourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildUserInput)(nil)).Elem(), BuildUserArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildUserArrayInput)(nil)).Elem(), BuildUserArray{})
	pulumi.RegisterOutputType(BuildSourceOutput{})
	pulumi.RegisterOutputType(BuildSourcePtrOutput{})
	pulumi.RegisterOutputType(BuildUserOutput{})
	pulumi.RegisterOutputType(BuildUserArrayOutput{})
}
