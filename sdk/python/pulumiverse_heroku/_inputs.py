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
from . import _utilities

__all__ = [
    'ProviderCustomizationArgs',
    'ProviderCustomizationArgsDict',
    'ProviderDelaysArgs',
    'ProviderDelaysArgsDict',
    'ProviderTimeoutsArgs',
    'ProviderTimeoutsArgsDict',
]

MYPY = False

if not MYPY:
    class ProviderCustomizationArgsDict(TypedDict):
        set_addon_config_vars_in_state: NotRequired[pulumi.Input[bool]]
        set_app_all_config_vars_in_state: NotRequired[pulumi.Input[bool]]
elif False:
    ProviderCustomizationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ProviderCustomizationArgs:
    def __init__(__self__, *,
                 set_addon_config_vars_in_state: Optional[pulumi.Input[bool]] = None,
                 set_app_all_config_vars_in_state: Optional[pulumi.Input[bool]] = None):
        if set_addon_config_vars_in_state is not None:
            pulumi.set(__self__, "set_addon_config_vars_in_state", set_addon_config_vars_in_state)
        if set_app_all_config_vars_in_state is not None:
            pulumi.set(__self__, "set_app_all_config_vars_in_state", set_app_all_config_vars_in_state)

    @property
    @pulumi.getter(name="setAddonConfigVarsInState")
    def set_addon_config_vars_in_state(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "set_addon_config_vars_in_state")

    @set_addon_config_vars_in_state.setter
    def set_addon_config_vars_in_state(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "set_addon_config_vars_in_state", value)

    @property
    @pulumi.getter(name="setAppAllConfigVarsInState")
    def set_app_all_config_vars_in_state(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "set_app_all_config_vars_in_state")

    @set_app_all_config_vars_in_state.setter
    def set_app_all_config_vars_in_state(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "set_app_all_config_vars_in_state", value)


if not MYPY:
    class ProviderDelaysArgsDict(TypedDict):
        post_app_create_delay: NotRequired[pulumi.Input[int]]
        post_domain_create_delay: NotRequired[pulumi.Input[int]]
        post_space_create_delay: NotRequired[pulumi.Input[int]]
elif False:
    ProviderDelaysArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ProviderDelaysArgs:
    def __init__(__self__, *,
                 post_app_create_delay: Optional[pulumi.Input[int]] = None,
                 post_domain_create_delay: Optional[pulumi.Input[int]] = None,
                 post_space_create_delay: Optional[pulumi.Input[int]] = None):
        if post_app_create_delay is not None:
            pulumi.set(__self__, "post_app_create_delay", post_app_create_delay)
        if post_domain_create_delay is not None:
            pulumi.set(__self__, "post_domain_create_delay", post_domain_create_delay)
        if post_space_create_delay is not None:
            pulumi.set(__self__, "post_space_create_delay", post_space_create_delay)

    @property
    @pulumi.getter(name="postAppCreateDelay")
    def post_app_create_delay(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "post_app_create_delay")

    @post_app_create_delay.setter
    def post_app_create_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "post_app_create_delay", value)

    @property
    @pulumi.getter(name="postDomainCreateDelay")
    def post_domain_create_delay(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "post_domain_create_delay")

    @post_domain_create_delay.setter
    def post_domain_create_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "post_domain_create_delay", value)

    @property
    @pulumi.getter(name="postSpaceCreateDelay")
    def post_space_create_delay(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "post_space_create_delay")

    @post_space_create_delay.setter
    def post_space_create_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "post_space_create_delay", value)


if not MYPY:
    class ProviderTimeoutsArgsDict(TypedDict):
        addon_create_timeout: NotRequired[pulumi.Input[int]]
elif False:
    ProviderTimeoutsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ProviderTimeoutsArgs:
    def __init__(__self__, *,
                 addon_create_timeout: Optional[pulumi.Input[int]] = None):
        if addon_create_timeout is not None:
            pulumi.set(__self__, "addon_create_timeout", addon_create_timeout)

    @property
    @pulumi.getter(name="addonCreateTimeout")
    def addon_create_timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "addon_create_timeout")

    @addon_create_timeout.setter
    def addon_create_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "addon_create_timeout", value)


