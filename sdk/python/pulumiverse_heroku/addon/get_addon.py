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
    'GetAddonResult',
    'AwaitableGetAddonResult',
    'get_addon',
    'get_addon_output',
]

@pulumi.output_type
class GetAddonResult:
    """
    A collection of values returned by getAddon.
    """
    def __init__(__self__, app_id=None, config_vars=None, id=None, name=None, plan=None, provider_id=None):
        if app_id and not isinstance(app_id, str):
            raise TypeError("Expected argument 'app_id' to be a str")
        pulumi.set(__self__, "app_id", app_id)
        if config_vars and not isinstance(config_vars, list):
            raise TypeError("Expected argument 'config_vars' to be a list")
        pulumi.set(__self__, "config_vars", config_vars)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if plan and not isinstance(plan, str):
            raise TypeError("Expected argument 'plan' to be a str")
        pulumi.set(__self__, "plan", plan)
        if provider_id and not isinstance(provider_id, str):
            raise TypeError("Expected argument 'provider_id' to be a str")
        pulumi.set(__self__, "provider_id", provider_id)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> str:
        """
        Heroku app ID
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="configVars")
    def config_vars(self) -> Sequence[str]:
        """
        The Configuration variables of the add-on
        """
        return pulumi.get(self, "config_vars")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the add-on
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The add-on name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def plan(self) -> str:
        """
        The plan name
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> str:
        """
        The ID of the plan provider
        """
        return pulumi.get(self, "provider_id")


class AwaitableGetAddonResult(GetAddonResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAddonResult(
            app_id=self.app_id,
            config_vars=self.config_vars,
            id=self.id,
            name=self.name,
            plan=self.plan,
            provider_id=self.provider_id)


def get_addon(name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAddonResult:
    """
    Use this data source to get information about a Heroku Addon.

    ## Example Usage


    :param str name: The add-on name
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('heroku:addon/getAddon:getAddon', __args__, opts=opts, typ=GetAddonResult).value

    return AwaitableGetAddonResult(
        app_id=pulumi.get(__ret__, 'app_id'),
        config_vars=pulumi.get(__ret__, 'config_vars'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        plan=pulumi.get(__ret__, 'plan'),
        provider_id=pulumi.get(__ret__, 'provider_id'))
def get_addon_output(name: Optional[pulumi.Input[str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAddonResult]:
    """
    Use this data source to get information about a Heroku Addon.

    ## Example Usage


    :param str name: The add-on name
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('heroku:addon/getAddon:getAddon', __args__, opts=opts, typ=GetAddonResult)
    return __ret__.apply(lambda __response__: GetAddonResult(
        app_id=pulumi.get(__response__, 'app_id'),
        config_vars=pulumi.get(__response__, 'config_vars'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        plan=pulumi.get(__response__, 'plan'),
        provider_id=pulumi.get(__response__, 'provider_id')))
