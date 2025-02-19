// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a [Heroku Formation](https://devcenter.heroku.com/articles/platform-api-reference#formation)
 * resource.
 *
 * A formation represents the formation of processes that should be set for an application.
 *
 * Please note the following:
 * * The application must have a dyno in order to update its formation.
 * * If the heroku formation resource is removed and deleted, this will be a no-op action in Heroku.
 *   The Heroku Platform does not have a `DELETE` endpoint for `formation`.
 * * This resource works well with the `heroku.app.Release` resource, which allows you to deploy a slug/release to an application
 *   before the formation can be updated.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Existing formations can be imported using the combination of the application name, a colon, and the formation's type.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import heroku:formation/formation:Formation foobar-web foobar:web
 * ```
 */
export class Formation extends pulumi.CustomResource {
    /**
     * Get an existing Formation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FormationState, opts?: pulumi.CustomResourceOptions): Formation {
        return new Formation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'heroku:formation/formation:Formation';

    /**
     * Returns true if the given object is an instance of Formation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Formation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Formation.__pulumiType;
    }

    /**
     * Heroku app ID (do not use app name)
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * number of processes to maintain
     */
    public readonly quantity!: pulumi.Output<number>;
    /**
     * dyno size (Example: “standard-1X”). Capitalization does not matter.
     */
    public readonly size!: pulumi.Output<string>;
    /**
     * type of process such as "web"
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Formation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FormationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FormationArgs | FormationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FormationState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["quantity"] = state ? state.quantity : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as FormationArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.quantity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'quantity'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["quantity"] = args ? args.quantity : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Formation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Formation resources.
 */
export interface FormationState {
    /**
     * Heroku app ID (do not use app name)
     */
    appId?: pulumi.Input<string>;
    /**
     * number of processes to maintain
     */
    quantity?: pulumi.Input<number>;
    /**
     * dyno size (Example: “standard-1X”). Capitalization does not matter.
     */
    size?: pulumi.Input<string>;
    /**
     * type of process such as "web"
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Formation resource.
 */
export interface FormationArgs {
    /**
     * Heroku app ID (do not use app name)
     */
    appId: pulumi.Input<string>;
    /**
     * number of processes to maintain
     */
    quantity: pulumi.Input<number>;
    /**
     * dyno size (Example: “standard-1X”). Capitalization does not matter.
     */
    size: pulumi.Input<string>;
    /**
     * type of process such as "web"
     */
    type: pulumi.Input<string>;
}
