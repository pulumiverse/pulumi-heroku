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
    'BuildSourceArgs',
    'BuildSourceArgsDict',
    'BuildUserArgs',
    'BuildUserArgsDict',
]

MYPY = False

if not MYPY:
    class BuildSourceArgsDict(TypedDict):
        checksum: NotRequired[pulumi.Input[str]]
        """
        SHA256 hash of the tarball archive to verify its integrity, example:
        `SHA256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
        """
        path: NotRequired[pulumi.Input[str]]
        """
        Local path to the source directory or tarball archive for the app
        """
        url: NotRequired[pulumi.Input[str]]
        """
        `https` location of the source archive for the app
        """
        version: NotRequired[pulumi.Input[str]]
        """
        Use to track what version of your source originated this build. If you are creating builds
        from git-versioned source code, for example, the commit hash, or release tag would be a good value to use for the
        version parameter.
        """
elif False:
    BuildSourceArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class BuildSourceArgs:
    def __init__(__self__, *,
                 checksum: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] checksum: SHA256 hash of the tarball archive to verify its integrity, example:
               `SHA256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
        :param pulumi.Input[str] path: Local path to the source directory or tarball archive for the app
        :param pulumi.Input[str] url: `https` location of the source archive for the app
        :param pulumi.Input[str] version: Use to track what version of your source originated this build. If you are creating builds
               from git-versioned source code, for example, the commit hash, or release tag would be a good value to use for the
               version parameter.
        """
        if checksum is not None:
            pulumi.set(__self__, "checksum", checksum)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def checksum(self) -> Optional[pulumi.Input[str]]:
        """
        SHA256 hash of the tarball archive to verify its integrity, example:
        `SHA256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
        """
        return pulumi.get(self, "checksum")

    @checksum.setter
    def checksum(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "checksum", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        Local path to the source directory or tarball archive for the app
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        `https` location of the source archive for the app
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Use to track what version of your source originated this build. If you are creating builds
        from git-versioned source code, for example, the commit hash, or release tag would be a good value to use for the
        version parameter.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


if not MYPY:
    class BuildUserArgsDict(TypedDict):
        email: NotRequired[pulumi.Input[str]]
        id: NotRequired[pulumi.Input[str]]
elif False:
    BuildUserArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class BuildUserArgs:
    def __init__(__self__, *,
                 email: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None):
        if email is not None:
            pulumi.set(__self__, "email", email)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


