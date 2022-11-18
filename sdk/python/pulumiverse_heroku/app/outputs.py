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
    'AppOrganization',
    'GetAppOrganizationResult',
]

@pulumi.output_type
class AppOrganization(dict):
    def __init__(__self__, *,
                 name: str,
                 locked: Optional[bool] = None,
                 personal: Optional[bool] = None):
        pulumi.set(__self__, "name", name)
        if locked is not None:
            pulumi.set(__self__, "locked", locked)
        if personal is not None:
            pulumi.set(__self__, "personal", personal)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def locked(self) -> Optional[bool]:
        return pulumi.get(self, "locked")

    @property
    @pulumi.getter
    def personal(self) -> Optional[bool]:
        return pulumi.get(self, "personal")


@pulumi.output_type
class GetAppOrganizationResult(dict):
    def __init__(__self__, *,
                 locked: bool,
                 name: str,
                 personal: bool):
        pulumi.set(__self__, "locked", locked)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "personal", personal)

    @property
    @pulumi.getter
    def locked(self) -> bool:
        return pulumi.get(self, "locked")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def personal(self) -> bool:
        return pulumi.get(self, "personal")

