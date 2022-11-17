// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AppArgs, AppState } from "./app";
export type App = import("./app").App;
export const App: typeof import("./app").App = null as any;
utilities.lazyLoad(exports, ["App"], () => require("./app"));

export { ConfigArgs, ConfigState } from "./config";
export type Config = import("./config").Config;
export const Config: typeof import("./config").Config = null as any;
utilities.lazyLoad(exports, ["Config"], () => require("./config"));

export { ConfigAssociationArgs, ConfigAssociationState } from "./configAssociation";
export type ConfigAssociation = import("./configAssociation").ConfigAssociation;
export const ConfigAssociation: typeof import("./configAssociation").ConfigAssociation = null as any;
utilities.lazyLoad(exports, ["ConfigAssociation"], () => require("./configAssociation"));

export { FeatureArgs, FeatureState } from "./feature";
export type Feature = import("./feature").Feature;
export const Feature: typeof import("./feature").Feature = null as any;
utilities.lazyLoad(exports, ["Feature"], () => require("./feature"));

export { GetAppArgs, GetAppResult, GetAppOutputArgs } from "./getApp";
export const getApp: typeof import("./getApp").getApp = null as any;
export const getAppOutput: typeof import("./getApp").getAppOutput = null as any;
utilities.lazyLoad(exports, ["getApp","getAppOutput"], () => require("./getApp"));

export { ReleaseArgs, ReleaseState } from "./release";
export type Release = import("./release").Release;
export const Release: typeof import("./release").Release = null as any;
utilities.lazyLoad(exports, ["Release"], () => require("./release"));

export { WebhookArgs, WebhookState } from "./webhook";
export type Webhook = import("./webhook").Webhook;
export const Webhook: typeof import("./webhook").Webhook = null as any;
utilities.lazyLoad(exports, ["Webhook"], () => require("./webhook"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "heroku:app/app:App":
                return new App(name, <any>undefined, { urn })
            case "heroku:app/config:Config":
                return new Config(name, <any>undefined, { urn })
            case "heroku:app/configAssociation:ConfigAssociation":
                return new ConfigAssociation(name, <any>undefined, { urn })
            case "heroku:app/feature:Feature":
                return new Feature(name, <any>undefined, { urn })
            case "heroku:app/release:Release":
                return new Release(name, <any>undefined, { urn })
            case "heroku:app/webhook:Webhook":
                return new Webhook(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("heroku", "app/app", _module)
pulumi.runtime.registerResourceModule("heroku", "app/config", _module)
pulumi.runtime.registerResourceModule("heroku", "app/configAssociation", _module)
pulumi.runtime.registerResourceModule("heroku", "app/feature", _module)
pulumi.runtime.registerResourceModule("heroku", "app/release", _module)
pulumi.runtime.registerResourceModule("heroku", "app/webhook", _module)
