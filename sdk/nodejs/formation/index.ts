// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { FormationArgs, FormationState } from "./formation";
export type Formation = import("./formation").Formation;
export const Formation: typeof import("./formation").Formation = null as any;
utilities.lazyLoad(exports, ["Formation"], () => require("./formation"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "heroku:formation/formation:Formation":
                return new Formation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("heroku", "formation/formation", _module)
