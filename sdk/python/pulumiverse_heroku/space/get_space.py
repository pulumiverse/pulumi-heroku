# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSpaceResult',
    'AwaitableGetSpaceResult',
    'get_space',
    'get_space_output',
]

@pulumi.output_type
class GetSpaceResult:
    """
    A collection of values returned by getSpace.
    """
    def __init__(__self__, cidr=None, data_cidr=None, id=None, name=None, organization=None, outbound_ips=None, region=None, shield=None, state=None):
        if cidr and not isinstance(cidr, str):
            raise TypeError("Expected argument 'cidr' to be a str")
        pulumi.set(__self__, "cidr", cidr)
        if data_cidr and not isinstance(data_cidr, str):
            raise TypeError("Expected argument 'data_cidr' to be a str")
        pulumi.set(__self__, "data_cidr", data_cidr)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        pulumi.set(__self__, "organization", organization)
        if outbound_ips and not isinstance(outbound_ips, list):
            raise TypeError("Expected argument 'outbound_ips' to be a list")
        pulumi.set(__self__, "outbound_ips", outbound_ips)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if shield and not isinstance(shield, bool):
            raise TypeError("Expected argument 'shield' to be a bool")
        pulumi.set(__self__, "shield", shield)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def cidr(self) -> str:
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="dataCidr")
    def data_cidr(self) -> str:
        return pulumi.get(self, "data_cidr")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def organization(self) -> str:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="outboundIps")
    def outbound_ips(self) -> Sequence[str]:
        return pulumi.get(self, "outbound_ips")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def shield(self) -> bool:
        return pulumi.get(self, "shield")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")


class AwaitableGetSpaceResult(GetSpaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSpaceResult(
            cidr=self.cidr,
            data_cidr=self.data_cidr,
            id=self.id,
            name=self.name,
            organization=self.organization,
            outbound_ips=self.outbound_ips,
            region=self.region,
            shield=self.shield,
            state=self.state)


def get_space(cidr: Optional[str] = None,
              data_cidr: Optional[str] = None,
              name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSpaceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['cidr'] = cidr
    __args__['dataCidr'] = data_cidr
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('heroku:space/getSpace:getSpace', __args__, opts=opts, typ=GetSpaceResult).value

    return AwaitableGetSpaceResult(
        cidr=pulumi.get(__ret__, 'cidr'),
        data_cidr=pulumi.get(__ret__, 'data_cidr'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        organization=pulumi.get(__ret__, 'organization'),
        outbound_ips=pulumi.get(__ret__, 'outbound_ips'),
        region=pulumi.get(__ret__, 'region'),
        shield=pulumi.get(__ret__, 'shield'),
        state=pulumi.get(__ret__, 'state'))


@_utilities.lift_output_func(get_space)
def get_space_output(cidr: Optional[pulumi.Input[Optional[str]]] = None,
                     data_cidr: Optional[pulumi.Input[Optional[str]]] = None,
                     name: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSpaceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
