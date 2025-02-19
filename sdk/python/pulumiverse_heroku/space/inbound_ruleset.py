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
from . import outputs
from ._inputs import *

__all__ = ['InboundRulesetArgs', 'InboundRuleset']

@pulumi.input_type
class InboundRulesetArgs:
    def __init__(__self__, *,
                 rules: pulumi.Input[Sequence[pulumi.Input['InboundRulesetRuleArgs']]],
                 space: pulumi.Input[str]):
        """
        The set of arguments for constructing a InboundRuleset resource.
        :param pulumi.Input[Sequence[pulumi.Input['InboundRulesetRuleArgs']]] rules: At least one `rule` block. Rules are documented below.
        :param pulumi.Input[str] space: The ID of the space.
        """
        pulumi.set(__self__, "rules", rules)
        pulumi.set(__self__, "space", space)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['InboundRulesetRuleArgs']]]:
        """
        At least one `rule` block. Rules are documented below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['InboundRulesetRuleArgs']]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def space(self) -> pulumi.Input[str]:
        """
        The ID of the space.
        """
        return pulumi.get(self, "space")

    @space.setter
    def space(self, value: pulumi.Input[str]):
        pulumi.set(self, "space", value)


@pulumi.input_type
class _InboundRulesetState:
    def __init__(__self__, *,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['InboundRulesetRuleArgs']]]] = None,
                 space: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InboundRuleset resources.
        :param pulumi.Input[Sequence[pulumi.Input['InboundRulesetRuleArgs']]] rules: At least one `rule` block. Rules are documented below.
        :param pulumi.Input[str] space: The ID of the space.
        """
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if space is not None:
            pulumi.set(__self__, "space", space)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InboundRulesetRuleArgs']]]]:
        """
        At least one `rule` block. Rules are documented below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InboundRulesetRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def space(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the space.
        """
        return pulumi.get(self, "space")

    @space.setter
    def space(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space", value)


class InboundRuleset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InboundRulesetRuleArgs', 'InboundRulesetRuleArgsDict']]]]] = None,
                 space: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource for managing [inbound rulesets](https://devcenter.heroku.com/articles/platform-api-reference#inbound-ruleset) for Heroku Private Spaces.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['InboundRulesetRuleArgs', 'InboundRulesetRuleArgsDict']]]] rules: At least one `rule` block. Rules are documented below.
        :param pulumi.Input[str] space: The ID of the space.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InboundRulesetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource for managing [inbound rulesets](https://devcenter.heroku.com/articles/platform-api-reference#inbound-ruleset) for Heroku Private Spaces.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param InboundRulesetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InboundRulesetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InboundRulesetRuleArgs', 'InboundRulesetRuleArgsDict']]]]] = None,
                 space: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InboundRulesetArgs.__new__(InboundRulesetArgs)

            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
            if space is None and not opts.urn:
                raise TypeError("Missing required property 'space'")
            __props__.__dict__["space"] = space
        super(InboundRuleset, __self__).__init__(
            'heroku:space/inboundRuleset:InboundRuleset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InboundRulesetRuleArgs', 'InboundRulesetRuleArgsDict']]]]] = None,
            space: Optional[pulumi.Input[str]] = None) -> 'InboundRuleset':
        """
        Get an existing InboundRuleset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['InboundRulesetRuleArgs', 'InboundRulesetRuleArgsDict']]]] rules: At least one `rule` block. Rules are documented below.
        :param pulumi.Input[str] space: The ID of the space.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InboundRulesetState.__new__(_InboundRulesetState)

        __props__.__dict__["rules"] = rules
        __props__.__dict__["space"] = space
        return InboundRuleset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.InboundRulesetRule']]:
        """
        At least one `rule` block. Rules are documented below.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def space(self) -> pulumi.Output[str]:
        """
        The ID of the space.
        """
        return pulumi.get(self, "space")

