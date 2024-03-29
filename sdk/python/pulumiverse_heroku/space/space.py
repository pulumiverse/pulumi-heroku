# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SpaceArgs', 'Space']

@pulumi.input_type
class SpaceArgs:
    def __init__(__self__, *,
                 organization: pulumi.Input[str],
                 cidr: Optional[pulumi.Input[str]] = None,
                 data_cidr: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 shield: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Space resource.
        """
        pulumi.set(__self__, "organization", organization)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if data_cidr is not None:
            pulumi.set(__self__, "data_cidr", data_cidr)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if shield is not None:
            pulumi.set(__self__, "shield", shield)

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="dataCidr")
    def data_cidr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data_cidr")

    @data_cidr.setter
    def data_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_cidr", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def shield(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "shield")

    @shield.setter
    def shield(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shield", value)


@pulumi.input_type
class _SpaceState:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[str]] = None,
                 data_cidr: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 outbound_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 shield: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering Space resources.
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if data_cidr is not None:
            pulumi.set(__self__, "data_cidr", data_cidr)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if outbound_ips is not None:
            pulumi.set(__self__, "outbound_ips", outbound_ips)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if shield is not None:
            pulumi.set(__self__, "shield", shield)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="dataCidr")
    def data_cidr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data_cidr")

    @data_cidr.setter
    def data_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_cidr", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def organization(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter(name="outboundIps")
    def outbound_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "outbound_ips")

    @outbound_ips.setter
    def outbound_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "outbound_ips", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def shield(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "shield")

    @shield.setter
    def shield(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shield", value)


class Space(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 data_cidr: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 shield: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a Space resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SpaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Space resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SpaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SpaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 data_cidr: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 shield: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SpaceArgs.__new__(SpaceArgs)

            __props__.__dict__["cidr"] = cidr
            __props__.__dict__["data_cidr"] = data_cidr
            __props__.__dict__["name"] = name
            if organization is None and not opts.urn:
                raise TypeError("Missing required property 'organization'")
            __props__.__dict__["organization"] = organization
            __props__.__dict__["region"] = region
            __props__.__dict__["shield"] = shield
            __props__.__dict__["outbound_ips"] = None
        super(Space, __self__).__init__(
            'heroku:space/space:Space',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            data_cidr: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization: Optional[pulumi.Input[str]] = None,
            outbound_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            region: Optional[pulumi.Input[str]] = None,
            shield: Optional[pulumi.Input[bool]] = None) -> 'Space':
        """
        Get an existing Space resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SpaceState.__new__(_SpaceState)

        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["data_cidr"] = data_cidr
        __props__.__dict__["name"] = name
        __props__.__dict__["organization"] = organization
        __props__.__dict__["outbound_ips"] = outbound_ips
        __props__.__dict__["region"] = region
        __props__.__dict__["shield"] = shield
        return Space(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="dataCidr")
    def data_cidr(self) -> pulumi.Output[str]:
        return pulumi.get(self, "data_cidr")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Output[str]:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="outboundIps")
    def outbound_ips(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "outbound_ips")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def shield(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "shield")

