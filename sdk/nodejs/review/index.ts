// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AppConfigArgs, AppConfigState } from "./appConfig";
export type AppConfig = import("./appConfig").AppConfig;
export const AppConfig: typeof import("./appConfig").AppConfig = null as any;
utilities.lazyLoad(exports, ["AppConfig"], () => require("./appConfig"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "heroku:review/appConfig:AppConfig":
                return new AppConfig(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("heroku", "review/appConfig", _module)
