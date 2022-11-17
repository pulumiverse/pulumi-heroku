// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { SslArgs, SslState } from "./ssl";
export type Ssl = import("./ssl").Ssl;
export const Ssl: typeof import("./ssl").Ssl = null as any;
utilities.lazyLoad(exports, ["Ssl"], () => require("./ssl"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "heroku:ssl/ssl:Ssl":
                return new Ssl(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("heroku", "ssl/ssl", _module)
