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
    'AppConfigDeployTarget',
]

@pulumi.output_type
class AppConfigDeployTarget(dict):
    def __init__(__self__, *,
                 id: str,
                 type: str):
        """
        :param str id: Unique identifier of deploy target.
        :param str type: Type of deploy target. Must be either `space` or `region`.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Unique identifier of deploy target.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of deploy target. Must be either `space` or `region`.
        """
        return pulumi.get(self, "type")


