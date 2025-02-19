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

__all__ = ['PeeringConnectionAccepterArgs', 'PeeringConnectionAccepter']

@pulumi.input_type
class PeeringConnectionAccepterArgs:
    def __init__(__self__, *,
                 space: pulumi.Input[str],
                 vpc_peering_connection_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a PeeringConnectionAccepter resource.
        :param pulumi.Input[str] space: The ID of the space.
        :param pulumi.Input[str] vpc_peering_connection_id: The peering connection request ID.
        """
        pulumi.set(__self__, "space", space)
        pulumi.set(__self__, "vpc_peering_connection_id", vpc_peering_connection_id)

    @property
    @pulumi.getter
    def space(self) -> pulumi.Input[str]:
        """
        The ID of the space.
        """
        return pulumi.get(self, "space")

    @space.setter
    def space(self, value: pulumi.Input[str]):
        pulumi.set(self, "space", value)

    @property
    @pulumi.getter(name="vpcPeeringConnectionId")
    def vpc_peering_connection_id(self) -> pulumi.Input[str]:
        """
        The peering connection request ID.
        """
        return pulumi.get(self, "vpc_peering_connection_id")

    @vpc_peering_connection_id.setter
    def vpc_peering_connection_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_peering_connection_id", value)


@pulumi.input_type
class _PeeringConnectionAccepterState:
    def __init__(__self__, *,
                 space: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpc_peering_connection_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PeeringConnectionAccepter resources.
        :param pulumi.Input[str] space: The ID of the space.
        :param pulumi.Input[str] status: The status of the peering connection request.
        :param pulumi.Input[str] type: The type of the peering connection.
        :param pulumi.Input[str] vpc_peering_connection_id: The peering connection request ID.
        """
        if space is not None:
            pulumi.set(__self__, "space", space)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vpc_peering_connection_id is not None:
            pulumi.set(__self__, "vpc_peering_connection_id", vpc_peering_connection_id)

    @property
    @pulumi.getter
    def space(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the space.
        """
        return pulumi.get(self, "space")

    @space.setter
    def space(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the peering connection request.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the peering connection.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vpcPeeringConnectionId")
    def vpc_peering_connection_id(self) -> Optional[pulumi.Input[str]]:
        """
        The peering connection request ID.
        """
        return pulumi.get(self, "vpc_peering_connection_id")

    @vpc_peering_connection_id.setter
    def vpc_peering_connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_peering_connection_id", value)


class PeeringConnectionAccepter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 space: Optional[pulumi.Input[str]] = None,
                 vpc_peering_connection_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource for accepting VPC peering requests to Heroku Private Spaces.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] space: The ID of the space.
        :param pulumi.Input[str] vpc_peering_connection_id: The peering connection request ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PeeringConnectionAccepterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource for accepting VPC peering requests to Heroku Private Spaces.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param PeeringConnectionAccepterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PeeringConnectionAccepterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 space: Optional[pulumi.Input[str]] = None,
                 vpc_peering_connection_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PeeringConnectionAccepterArgs.__new__(PeeringConnectionAccepterArgs)

            if space is None and not opts.urn:
                raise TypeError("Missing required property 'space'")
            __props__.__dict__["space"] = space
            if vpc_peering_connection_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_peering_connection_id'")
            __props__.__dict__["vpc_peering_connection_id"] = vpc_peering_connection_id
            __props__.__dict__["status"] = None
            __props__.__dict__["type"] = None
        super(PeeringConnectionAccepter, __self__).__init__(
            'heroku:space/peeringConnectionAccepter:PeeringConnectionAccepter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            space: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vpc_peering_connection_id: Optional[pulumi.Input[str]] = None) -> 'PeeringConnectionAccepter':
        """
        Get an existing PeeringConnectionAccepter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] space: The ID of the space.
        :param pulumi.Input[str] status: The status of the peering connection request.
        :param pulumi.Input[str] type: The type of the peering connection.
        :param pulumi.Input[str] vpc_peering_connection_id: The peering connection request ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PeeringConnectionAccepterState.__new__(_PeeringConnectionAccepterState)

        __props__.__dict__["space"] = space
        __props__.__dict__["status"] = status
        __props__.__dict__["type"] = type
        __props__.__dict__["vpc_peering_connection_id"] = vpc_peering_connection_id
        return PeeringConnectionAccepter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def space(self) -> pulumi.Output[str]:
        """
        The ID of the space.
        """
        return pulumi.get(self, "space")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the peering connection request.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the peering connection.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcPeeringConnectionId")
    def vpc_peering_connection_id(self) -> pulumi.Output[str]:
        """
        The peering connection request ID.
        """
        return pulumi.get(self, "vpc_peering_connection_id")

