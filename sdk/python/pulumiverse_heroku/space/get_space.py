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
        """
        The RFC-1918 CIDR the Private Space will use. It must be a /16 in 10.0.0.0/8, 172.16.0.0/12 or 192.168.0.0/16
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="dataCidr")
    def data_cidr(self) -> str:
        """
        The RFC-1918 CIDR that the Private Space will use for the Heroku-managed peering connection that’s automatically created when using Heroku Data add-ons. It must be between a /16 and a /20
        """
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
        """
        (string) - The name of the Heroku Team.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def organization(self) -> str:
        """
        The Heroku Team that owns this space. The fields for this block are documented below.
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="outboundIps")
    def outbound_ips(self) -> Sequence[str]:
        """
        The space's stable outbound [NAT IPs](https://devcenter.heroku.com/articles/platform-api-reference#space-network-address-translation).
        """
        return pulumi.get(self, "outbound_ips")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region in which the Heroku Private Space is deployed.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def shield(self) -> bool:
        """
        Whether or not the space has [Shield](https://devcenter.heroku.com/articles/private-spaces#shield-private-spaces) turned on. One of `on` or `off`.
        """
        return pulumi.get(self, "shield")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the Heroku Private Space. Either `allocating` or `allocated`.
        """
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
    Use this data source to get information about a [Heroku Private Space](https://www.heroku.com/private-spaces).

    ## Example Usage


    :param str cidr: The RFC-1918 CIDR the Private Space will use. It must be a /16 in 10.0.0.0/8, 172.16.0.0/12 or 192.168.0.0/16
    :param str data_cidr: The RFC-1918 CIDR that the Private Space will use for the Heroku-managed peering connection that’s automatically created when using Heroku Data add-ons. It must be between a /16 and a /20
    :param str name: The name of the Heroku Private Space.
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
def get_space_output(cidr: Optional[pulumi.Input[Optional[str]]] = None,
                     data_cidr: Optional[pulumi.Input[Optional[str]]] = None,
                     name: Optional[pulumi.Input[str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSpaceResult]:
    """
    Use this data source to get information about a [Heroku Private Space](https://www.heroku.com/private-spaces).

    ## Example Usage


    :param str cidr: The RFC-1918 CIDR the Private Space will use. It must be a /16 in 10.0.0.0/8, 172.16.0.0/12 or 192.168.0.0/16
    :param str data_cidr: The RFC-1918 CIDR that the Private Space will use for the Heroku-managed peering connection that’s automatically created when using Heroku Data add-ons. It must be between a /16 and a /20
    :param str name: The name of the Heroku Private Space.
    """
    __args__ = dict()
    __args__['cidr'] = cidr
    __args__['dataCidr'] = data_cidr
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('heroku:space/getSpace:getSpace', __args__, opts=opts, typ=GetSpaceResult)
    return __ret__.apply(lambda __response__: GetSpaceResult(
        cidr=pulumi.get(__response__, 'cidr'),
        data_cidr=pulumi.get(__response__, 'data_cidr'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        organization=pulumi.get(__response__, 'organization'),
        outbound_ips=pulumi.get(__response__, 'outbound_ips'),
        region=pulumi.get(__response__, 'region'),
        shield=pulumi.get(__response__, 'shield'),
        state=pulumi.get(__response__, 'state')))
