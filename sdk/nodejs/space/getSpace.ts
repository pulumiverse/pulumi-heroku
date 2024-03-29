// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getSpace(args: GetSpaceArgs, opts?: pulumi.InvokeOptions): Promise<GetSpaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("heroku:space/getSpace:getSpace", {
        "cidr": args.cidr,
        "dataCidr": args.dataCidr,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSpace.
 */
export interface GetSpaceArgs {
    cidr?: string;
    dataCidr?: string;
    name: string;
}

/**
 * A collection of values returned by getSpace.
 */
export interface GetSpaceResult {
    readonly cidr: string;
    readonly dataCidr: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly organization: string;
    readonly outboundIps: string[];
    readonly region: string;
    readonly shield: boolean;
    readonly state: string;
}
export function getSpaceOutput(args: GetSpaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSpaceResult> {
    return pulumi.output(args).apply((a: any) => getSpace(a, opts))
}

/**
 * A collection of arguments for invoking getSpace.
 */
export interface GetSpaceOutputArgs {
    cidr?: pulumi.Input<string>;
    dataCidr?: pulumi.Input<string>;
    name: pulumi.Input<string>;
}
