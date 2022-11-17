// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pipeline

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PipelineOwner struct {
	Id   string `pulumi:"id"`
	Type string `pulumi:"type"`
}

// PipelineOwnerInput is an input type that accepts PipelineOwnerArgs and PipelineOwnerOutput values.
// You can construct a concrete instance of `PipelineOwnerInput` via:
//
//          PipelineOwnerArgs{...}
type PipelineOwnerInput interface {
	pulumi.Input

	ToPipelineOwnerOutput() PipelineOwnerOutput
	ToPipelineOwnerOutputWithContext(context.Context) PipelineOwnerOutput
}

type PipelineOwnerArgs struct {
	Id   pulumi.StringInput `pulumi:"id"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (PipelineOwnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineOwner)(nil)).Elem()
}

func (i PipelineOwnerArgs) ToPipelineOwnerOutput() PipelineOwnerOutput {
	return i.ToPipelineOwnerOutputWithContext(context.Background())
}

func (i PipelineOwnerArgs) ToPipelineOwnerOutputWithContext(ctx context.Context) PipelineOwnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOwnerOutput)
}

func (i PipelineOwnerArgs) ToPipelineOwnerPtrOutput() PipelineOwnerPtrOutput {
	return i.ToPipelineOwnerPtrOutputWithContext(context.Background())
}

func (i PipelineOwnerArgs) ToPipelineOwnerPtrOutputWithContext(ctx context.Context) PipelineOwnerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOwnerOutput).ToPipelineOwnerPtrOutputWithContext(ctx)
}

// PipelineOwnerPtrInput is an input type that accepts PipelineOwnerArgs, PipelineOwnerPtr and PipelineOwnerPtrOutput values.
// You can construct a concrete instance of `PipelineOwnerPtrInput` via:
//
//          PipelineOwnerArgs{...}
//
//  or:
//
//          nil
type PipelineOwnerPtrInput interface {
	pulumi.Input

	ToPipelineOwnerPtrOutput() PipelineOwnerPtrOutput
	ToPipelineOwnerPtrOutputWithContext(context.Context) PipelineOwnerPtrOutput
}

type pipelineOwnerPtrType PipelineOwnerArgs

func PipelineOwnerPtr(v *PipelineOwnerArgs) PipelineOwnerPtrInput {
	return (*pipelineOwnerPtrType)(v)
}

func (*pipelineOwnerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineOwner)(nil)).Elem()
}

func (i *pipelineOwnerPtrType) ToPipelineOwnerPtrOutput() PipelineOwnerPtrOutput {
	return i.ToPipelineOwnerPtrOutputWithContext(context.Background())
}

func (i *pipelineOwnerPtrType) ToPipelineOwnerPtrOutputWithContext(ctx context.Context) PipelineOwnerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOwnerPtrOutput)
}

type PipelineOwnerOutput struct{ *pulumi.OutputState }

func (PipelineOwnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineOwner)(nil)).Elem()
}

func (o PipelineOwnerOutput) ToPipelineOwnerOutput() PipelineOwnerOutput {
	return o
}

func (o PipelineOwnerOutput) ToPipelineOwnerOutputWithContext(ctx context.Context) PipelineOwnerOutput {
	return o
}

func (o PipelineOwnerOutput) ToPipelineOwnerPtrOutput() PipelineOwnerPtrOutput {
	return o.ToPipelineOwnerPtrOutputWithContext(context.Background())
}

func (o PipelineOwnerOutput) ToPipelineOwnerPtrOutputWithContext(ctx context.Context) PipelineOwnerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineOwner) *PipelineOwner {
		return &v
	}).(PipelineOwnerPtrOutput)
}

func (o PipelineOwnerOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PipelineOwner) string { return v.Id }).(pulumi.StringOutput)
}

func (o PipelineOwnerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PipelineOwner) string { return v.Type }).(pulumi.StringOutput)
}

type PipelineOwnerPtrOutput struct{ *pulumi.OutputState }

func (PipelineOwnerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineOwner)(nil)).Elem()
}

func (o PipelineOwnerPtrOutput) ToPipelineOwnerPtrOutput() PipelineOwnerPtrOutput {
	return o
}

func (o PipelineOwnerPtrOutput) ToPipelineOwnerPtrOutputWithContext(ctx context.Context) PipelineOwnerPtrOutput {
	return o
}

func (o PipelineOwnerPtrOutput) Elem() PipelineOwnerOutput {
	return o.ApplyT(func(v *PipelineOwner) PipelineOwner {
		if v != nil {
			return *v
		}
		var ret PipelineOwner
		return ret
	}).(PipelineOwnerOutput)
}

func (o PipelineOwnerPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineOwner) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PipelineOwnerPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineOwner) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineOwnerInput)(nil)).Elem(), PipelineOwnerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineOwnerPtrInput)(nil)).Elem(), PipelineOwnerArgs{})
	pulumi.RegisterOutputType(PipelineOwnerOutput{})
	pulumi.RegisterOutputType(PipelineOwnerPtrOutput{})
}
