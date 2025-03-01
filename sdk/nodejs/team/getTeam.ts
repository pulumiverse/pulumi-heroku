// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a Heroku Team.
 *
 * ## Example Usage
 */
export function getTeam(args: GetTeamArgs, opts?: pulumi.InvokeOptions): Promise<GetTeamResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("heroku:team/getTeam:getTeam", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTeam.
 */
export interface GetTeamArgs {
    /**
     * The team name
     */
    name: string;
}

/**
 * A collection of values returned by getTeam.
 */
export interface GetTeamResult {
    /**
     * Whether to use this team when none is specified
     */
    readonly default: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Upper limit of members allowed in a team
     */
    readonly membershipLimit: number;
    readonly name: string;
    /**
     * Whether the team is provisioned licenses by Salesforce
     */
    readonly provisionedLicenses: boolean;
    /**
     * type of team Will likely be either "enterprise" or "team"
     */
    readonly type: string;
}
/**
 * Use this data source to get information about a Heroku Team.
 *
 * ## Example Usage
 */
export function getTeamOutput(args: GetTeamOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTeamResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("heroku:team/getTeam:getTeam", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTeam.
 */
export interface GetTeamOutputArgs {
    /**
     * The team name
     */
    name: pulumi.Input<string>;
}
