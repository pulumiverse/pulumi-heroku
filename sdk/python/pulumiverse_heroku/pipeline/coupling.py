# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['CouplingArgs', 'Coupling']

@pulumi.input_type
class CouplingArgs:
    def __init__(__self__, *,
                 app_id: pulumi.Input[str],
                 pipeline: pulumi.Input[str],
                 stage: pulumi.Input[str]):
        """
        The set of arguments for constructing a Coupling resource.
        :param pulumi.Input[str] app_id: Heroku app ID (do not use app name)
        :param pulumi.Input[str] pipeline: The ID of the pipeline to add this app to.
        :param pulumi.Input[str] stage: The stage to couple this app to. Must be one of
               `review`, `development`, `staging`, or `production`.
        """
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "pipeline", pipeline)
        pulumi.set(__self__, "stage", stage)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[str]:
        """
        Heroku app ID (do not use app name)
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter
    def pipeline(self) -> pulumi.Input[str]:
        """
        The ID of the pipeline to add this app to.
        """
        return pulumi.get(self, "pipeline")

    @pipeline.setter
    def pipeline(self, value: pulumi.Input[str]):
        pulumi.set(self, "pipeline", value)

    @property
    @pulumi.getter
    def stage(self) -> pulumi.Input[str]:
        """
        The stage to couple this app to. Must be one of
        `review`, `development`, `staging`, or `production`.
        """
        return pulumi.get(self, "stage")

    @stage.setter
    def stage(self, value: pulumi.Input[str]):
        pulumi.set(self, "stage", value)


@pulumi.input_type
class _CouplingState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[str]] = None,
                 pipeline: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Coupling resources.
        :param pulumi.Input[str] app_id: Heroku app ID (do not use app name)
        :param pulumi.Input[str] pipeline: The ID of the pipeline to add this app to.
        :param pulumi.Input[str] stage: The stage to couple this app to. Must be one of
               `review`, `development`, `staging`, or `production`.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if pipeline is not None:
            pulumi.set(__self__, "pipeline", pipeline)
        if stage is not None:
            pulumi.set(__self__, "stage", stage)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        Heroku app ID (do not use app name)
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter
    def pipeline(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the pipeline to add this app to.
        """
        return pulumi.get(self, "pipeline")

    @pipeline.setter
    def pipeline(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pipeline", value)

    @property
    @pulumi.getter
    def stage(self) -> Optional[pulumi.Input[str]]:
        """
        The stage to couple this app to. Must be one of
        `review`, `development`, `staging`, or `production`.
        """
        return pulumi.get(self, "stage")

    @stage.setter
    def stage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stage", value)


class Coupling(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 pipeline: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a [Heroku Pipeline Coupling](https://devcenter.heroku.com/articles/pipelines)
        resource.

        A pipeline is a group of Heroku apps that share the same codebase. Once a
        pipeline is created using `pipeline.Pipeline`, and apps are added
        to different stages using `pipeline.Coupling`, you can promote app slugs
        to the downstream stages.

        See `pipeline.Pipeline` for complete usage documentation.

        ## Example Usage

        ## Import

        Pipeline couplings can be imported using the Pipeline coupling `id`, e.g.

        ```sh
        $ pulumi import heroku:pipeline/coupling:Coupling foobar 12345678
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Heroku app ID (do not use app name)
        :param pulumi.Input[str] pipeline: The ID of the pipeline to add this app to.
        :param pulumi.Input[str] stage: The stage to couple this app to. Must be one of
               `review`, `development`, `staging`, or `production`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CouplingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a [Heroku Pipeline Coupling](https://devcenter.heroku.com/articles/pipelines)
        resource.

        A pipeline is a group of Heroku apps that share the same codebase. Once a
        pipeline is created using `pipeline.Pipeline`, and apps are added
        to different stages using `pipeline.Coupling`, you can promote app slugs
        to the downstream stages.

        See `pipeline.Pipeline` for complete usage documentation.

        ## Example Usage

        ## Import

        Pipeline couplings can be imported using the Pipeline coupling `id`, e.g.

        ```sh
        $ pulumi import heroku:pipeline/coupling:Coupling foobar 12345678
        ```

        :param str resource_name: The name of the resource.
        :param CouplingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CouplingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 pipeline: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CouplingArgs.__new__(CouplingArgs)

            if app_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_id'")
            __props__.__dict__["app_id"] = app_id
            if pipeline is None and not opts.urn:
                raise TypeError("Missing required property 'pipeline'")
            __props__.__dict__["pipeline"] = pipeline
            if stage is None and not opts.urn:
                raise TypeError("Missing required property 'stage'")
            __props__.__dict__["stage"] = stage
        super(Coupling, __self__).__init__(
            'heroku:pipeline/coupling:Coupling',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            pipeline: Optional[pulumi.Input[str]] = None,
            stage: Optional[pulumi.Input[str]] = None) -> 'Coupling':
        """
        Get an existing Coupling resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Heroku app ID (do not use app name)
        :param pulumi.Input[str] pipeline: The ID of the pipeline to add this app to.
        :param pulumi.Input[str] stage: The stage to couple this app to. Must be one of
               `review`, `development`, `staging`, or `production`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CouplingState.__new__(_CouplingState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["pipeline"] = pipeline
        __props__.__dict__["stage"] = stage
        return Coupling(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        Heroku app ID (do not use app name)
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter
    def pipeline(self) -> pulumi.Output[str]:
        """
        The ID of the pipeline to add this app to.
        """
        return pulumi.get(self, "pipeline")

    @property
    @pulumi.getter
    def stage(self) -> pulumi.Output[str]:
        """
        The stage to couple this app to. Must be one of
        `review`, `development`, `staging`, or `production`.
        """
        return pulumi.get(self, "stage")

