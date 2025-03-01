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
    'SlugBlobArgs',
    'SlugBlobArgsDict',
]

MYPY = False

if not MYPY:
    class SlugBlobArgsDict(TypedDict):
        method: NotRequired[pulumi.Input[str]]
        """
        HTTP method to upload the archive
        """
        url: NotRequired[pulumi.Input[str]]
        """
        Pre-signed, expiring URL to upload the archive
        """
elif False:
    SlugBlobArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SlugBlobArgs:
    def __init__(__self__, *,
                 method: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] method: HTTP method to upload the archive
        :param pulumi.Input[str] url: Pre-signed, expiring URL to upload the archive
        """
        if method is not None:
            pulumi.set(__self__, "method", method)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def method(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP method to upload the archive
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        Pre-signed, expiring URL to upload the archive
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


