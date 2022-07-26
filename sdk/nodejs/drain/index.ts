// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DrainArgs, DrainState } from "./drain";
export type Drain = import("./drain").Drain;
export const Drain: typeof import("./drain").Drain = null as any;
utilities.lazyLoad(exports, ["Drain"], () => require("./drain"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "heroku:drain/drain:Drain":
                return new Drain(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("heroku", "drain/drain", _module)
