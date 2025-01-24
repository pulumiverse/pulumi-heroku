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

__all__ = ['AddonArgs', 'Addon']

@pulumi.input_type
class AddonArgs:
    def __init__(__self__, *,
                 app_id: pulumi.Input[str],
                 plan: pulumi.Input[str],
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Addon resource.
        :param pulumi.Input[str] app_id: Heroku app ID (do not use app name)
        :param pulumi.Input[str] plan: The addon to add.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] config: Optional plan configuration.
        :param pulumi.Input[str] name: Globally unique name of the add-on.
        """
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "plan", plan)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[str]:
        """
        Heroku app ID (do not use app name)
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Input[str]:
        """
        The addon to add.
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: pulumi.Input[str]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional plan configuration.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Globally unique name of the add-on.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AddonState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 config_var_values: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 config_vars: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Addon resources.
        :param pulumi.Input[str] app_id: Heroku app ID (do not use app name)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] config: Optional plan configuration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] config_vars: The Configuration variables of the add-on
        :param pulumi.Input[str] name: Globally unique name of the add-on.
        :param pulumi.Input[str] plan: The addon to add.
        :param pulumi.Input[str] provider_id: The ID of the plan provider
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if config_var_values is not None:
            pulumi.set(__self__, "config_var_values", config_var_values)
        if config_vars is not None:
            pulumi.set(__self__, "config_vars", config_vars)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if provider_id is not None:
            pulumi.set(__self__, "provider_id", provider_id)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        Heroku app ID (do not use app name)
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional plan configuration.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="configVarValues")
    def config_var_values(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "config_var_values")

    @config_var_values.setter
    def config_var_values(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "config_var_values", value)

    @property
    @pulumi.getter(name="configVars")
    def config_vars(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The Configuration variables of the add-on
        """
        return pulumi.get(self, "config_vars")

    @config_vars.setter
    def config_vars(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "config_vars", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Globally unique name of the add-on.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input[str]]:
        """
        The addon to add.
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the plan provider
        """
        return pulumi.get(self, "provider_id")

    @provider_id.setter
    def provider_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_id", value)


class Addon(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Heroku Add-On resource. These can be attach
        services to a Heroku app.

        ## Example Usage

        ## Import

        Addons can be imported using the Addon `id`, e.g.

        ```sh
        $ pulumi import heroku:addon/addon:Addon foobar 12345678
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Heroku app ID (do not use app name)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] config: Optional plan configuration.
        :param pulumi.Input[str] name: Globally unique name of the add-on.
        :param pulumi.Input[str] plan: The addon to add.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AddonArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Heroku Add-On resource. These can be attach
        services to a Heroku app.

        ## Example Usage

        ## Import

        Addons can be imported using the Addon `id`, e.g.

        ```sh
        $ pulumi import heroku:addon/addon:Addon foobar 12345678
        ```

        :param str resource_name: The name of the resource.
        :param AddonArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AddonArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AddonArgs.__new__(AddonArgs)

            if app_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_id'")
            __props__.__dict__["app_id"] = app_id
            __props__.__dict__["config"] = config
            __props__.__dict__["name"] = name
            if plan is None and not opts.urn:
                raise TypeError("Missing required property 'plan'")
            __props__.__dict__["plan"] = plan
            __props__.__dict__["config_var_values"] = None
            __props__.__dict__["config_vars"] = None
            __props__.__dict__["provider_id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["configVarValues"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Addon, __self__).__init__(
            'heroku:addon/addon:Addon',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            config_var_values: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            config_vars: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            plan: Optional[pulumi.Input[str]] = None,
            provider_id: Optional[pulumi.Input[str]] = None) -> 'Addon':
        """
        Get an existing Addon resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Heroku app ID (do not use app name)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] config: Optional plan configuration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] config_vars: The Configuration variables of the add-on
        :param pulumi.Input[str] name: Globally unique name of the add-on.
        :param pulumi.Input[str] plan: The addon to add.
        :param pulumi.Input[str] provider_id: The ID of the plan provider
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AddonState.__new__(_AddonState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["config"] = config
        __props__.__dict__["config_var_values"] = config_var_values
        __props__.__dict__["config_vars"] = config_vars
        __props__.__dict__["name"] = name
        __props__.__dict__["plan"] = plan
        __props__.__dict__["provider_id"] = provider_id
        return Addon(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        Heroku app ID (do not use app name)
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Optional plan configuration.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="configVarValues")
    def config_var_values(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "config_var_values")

    @property
    @pulumi.getter(name="configVars")
    def config_vars(self) -> pulumi.Output[Sequence[str]]:
        """
        The Configuration variables of the add-on
        """
        return pulumi.get(self, "config_vars")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Globally unique name of the add-on.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output[str]:
        """
        The addon to add.
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> pulumi.Output[str]:
        """
        The ID of the plan provider
        """
        return pulumi.get(self, "provider_id")

