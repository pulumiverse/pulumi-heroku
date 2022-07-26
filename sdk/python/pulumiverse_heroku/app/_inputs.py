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
    'AppOrganizationArgs',
]

@pulumi.input_type
class AppOrganizationArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 locked: Optional[pulumi.Input[bool]] = None,
                 personal: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "name", name)
        if locked is not None:
            pulumi.set(__self__, "locked", locked)
        if personal is not None:
            pulumi.set(__self__, "personal", personal)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def locked(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "locked")

    @locked.setter
    def locked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "locked", value)

    @property
    @pulumi.getter
    def personal(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "personal")

    @personal.setter
    def personal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "personal", value)


